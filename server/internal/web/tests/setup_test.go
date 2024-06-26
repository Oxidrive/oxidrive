package web

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/oxidrive/oxidrive/server/internal/app"
	"github.com/oxidrive/oxidrive/server/internal/config"
	"github.com/oxidrive/oxidrive/server/internal/infrastructure"
	"github.com/oxidrive/oxidrive/server/internal/testutil"
	"github.com/oxidrive/oxidrive/server/internal/web"
)

const timeout time.Duration = 5 * time.Second

func setup(ctx context.Context, t *testing.T) (*app.Application, http.Handler) {
	t.Helper()

	dir := testutil.TempDirFromContext(ctx, t)
	dburl := testutil.Must(url.Parse(testutil.SqliteUrlFromContext(ctx, t)))
	db := testutil.SqliteDBFromContext(ctx, t)

	require.NoError(t, os.MkdirAll(filepath.Join(dir, "files"), 0700))

	logger := zerolog.New(zerolog.NewTestWriter(t))

	cfg := config.Config{
		PublicURL: testutil.Must(url.Parse("https://example.com")),
		AuthConfig: config.AuthConfig{
			SessionDuration: 5 * time.Minute,
		},
		DatabaseConfig: config.DatabaseConfig{
			Url: dburl,
		},
		StorageConfig: config.StorageConfig{
			StoragePrefix:      "/files",
			ThroughputInByte:   32,
			MultipartMaxMemory: 32,
			StorageFSConfig: config.StorageFSConfig{
				StorageFSDataDir: dir,
			},
		},
	}

	app := app.NewApplication(cfg, infrastructure.Setup(cfg, db, logger))

	wcfg := &web.Config{
		Address:            "",
		Application:        app,
		Logger:             logger,
		FrontendFolder:     "",
		MultipartMaxMemory: 0,
	}

	return app, testutil.Must(web.Router(wcfg))

}
