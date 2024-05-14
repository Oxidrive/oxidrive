package handler

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/oxidrive/oxidrive/server/internal/app"
	"github.com/oxidrive/oxidrive/server/internal/auth"
	"github.com/oxidrive/oxidrive/server/internal/core/file"
	"github.com/oxidrive/oxidrive/server/internal/web/api"
)

type Files struct {
	Logger             zerolog.Logger
	App                *app.Application
	MultipartMaxMemory int64
}

func (f *Files) Upload(ctx context.Context, request api.FilesUploadRequestObject) (api.FilesUploadResponseObject, error) {
	owner := auth.GetCurrentUser(ctx)

	form, err := request.Body.ReadForm(f.MultipartMaxMemory)
	if err != nil {
		return nil, err
	}

	paths := form.Value["path"]
	if len(paths) == 0 {
		return api.FilesUpload400JSONResponse{ErrorJSONResponse: api.ErrorJSONResponse(api.Error{
			Error:   "missing_path",
			Message: "form is missing required parameter 'path'",
		})}, nil
	}

	p := paths[0]

	files := form.File["file"]
	if len(files) == 0 {
		return api.FilesUpload400JSONResponse{ErrorJSONResponse: api.ErrorJSONResponse(api.Error{
			Error:   "missing_file",
			Message: "form is missing required parameter 'file'",
		})}, nil
	}

	fh := files[0]
	ff, err := fh.Open()
	if err != nil {
		return nil, err
	}

	upload := file.FileUpload{
		Content: file.Content(ff),
		Path:    file.Path(p),
		Size:    file.Size(fh.Size),
	}

	id, err := f.App.Files().Upload(ctx, upload, owner.ID)
	if err != nil {
		return nil, err
	}

	return api.FilesUpload200JSONResponse(api.FileUploadResponse{
		Ok: true,
		Id: id.String(),
	}), nil
}
