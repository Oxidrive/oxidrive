package file

import (
	"context"
	"errors"
	"fmt"

	"github.com/oxidrive/oxidrive/server/internal/core/list"
	"github.com/oxidrive/oxidrive/server/internal/core/user"
)

var (
	ErrFileNotFound error = errors.New("file does not exist")
)

type FileUpload struct {
	Content     Content
	ContentType ContentType
	Path        Path
	Size        Size
}

type Service struct {
	contents Contents
	files    Files
}

func NewService(filesContent Contents, filesMetadata Files) Service {
	return Service{
		contents: filesContent,
		files:    filesMetadata,
	}
}

func (s *Service) List(ctx context.Context, prefix *Path, params ...list.Param) (list.Of[File], error) {
	return s.files.List(ctx, prefix, list.NewParams(params...))
}

func (s *Service) ByID(ctx context.Context, id ID) (*File, error) {
	return s.files.ByID(ctx, id)
}

func (s *Service) ByOwnerByPath(ctx context.Context, owner user.ID, path Path) (*File, error) {
	return s.files.ByOwnerByPath(ctx, owner, path)
}

func (s *Service) Upload(ctx context.Context, upload FileUpload, owner user.ID) (ID, error) {
	f, err := s.files.ByOwnerByPath(ctx, owner, upload.Path)
	if err != nil {
		return EmptyID(), err
	}

	if f == nil {
		f, err = Create(upload.Content, upload.ContentType, upload.Path, upload.Size, owner)
	} else {
		err = f.UpdateContent(upload.Content, upload.ContentType, upload.Size)
	}

	if err != nil {
		return EmptyID(), err
	}

	if err := s.contents.Store(ctx, *f); err != nil {
		return EmptyID(), fmt.Errorf("failed to store the file content: %w", err)
	}

	if f, err = s.files.Save(ctx, *f); err != nil {
		return EmptyID(), fmt.Errorf("failed to save the file metadata: %w", err)
	}

	return f.ID, nil
}

func (s *Service) Download(ctx context.Context, f File) (Content, error) {
	return s.contents.Load(ctx, f)
}

func (s *Service) Move(ctx context.Context, f File, newPath Path) (*File, error) {
	if err := f.ChangePath(newPath); err != nil {
		return nil, fmt.Errorf("failed to change path of file %s: %w", f.ID, err)
	}

	updated, err := s.files.Save(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("failed to store file %s: %w", f.ID, err)
	}

	return updated, nil
}

func (s *Service) Delete(ctx context.Context, id ID) error {
	f, err := s.files.ByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to load file %s: %w", id, err)
	}

	if f == nil {
		return ErrFileNotFound
	}

	if err := s.files.Delete(ctx, *f); err != nil {
		// if the metadata are missing then we don't really care. Our goal was to remove it after all
		if errors.Is(err, ErrFileNotFound) {
			return fmt.Errorf("failed to delete metadata for file %s: %w", id, err)
		}
	}

	if err := s.contents.Delete(ctx, *f); err != nil {
		// if the content is missing then we don't really care. Our goal was to remove it after all
		if !errors.Is(err, ErrFileNotFound) {
			return fmt.Errorf("failed to delete content of file %s: %w", id, err)
		}
	}

	return nil
}
