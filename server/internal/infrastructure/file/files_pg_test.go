package file

import (
	"context"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/oxidrive/oxidrive/server/internal/core/file"
	"github.com/oxidrive/oxidrive/server/internal/core/user"
	userinfra "github.com/oxidrive/oxidrive/server/internal/infrastructure/user"
	"github.com/oxidrive/oxidrive/server/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPgFiles_Save(t *testing.T) {
	t.Run("saves a new file", func(t *testing.T) {
		ctx, done := testutil.IntegrationTest(context.Background(), t, testutil.WithPgDB())
		defer done()

		db := testutil.PgDBFromContext(ctx, t)
		u := insertPgUser(t, db, "username", "pwd")

		files := NewPgFiles(db)
		readerMock := strings.NewReader("")
		fileToSave, err := file.Create(readerMock, "filepath", 10, u.ID)
		require.NoError(t, err)

		saved, err := files.Save(ctx, *fileToSave)

		assert.NoError(t, err)
		assert.Equal(t, fileToSave.Name, saved.Name)
		assert.Equal(t, fileToSave.Path, saved.Path)
		assert.Equal(t, fileToSave.Size, saved.Size)
	})

	t.Run("saves an existing file", func(t *testing.T) {
		ctx, done := testutil.IntegrationTest(context.Background(), t, testutil.WithPgDB())
		defer done()

		db := testutil.PgDBFromContext(ctx, t)
		u := insertPgUser(t, db, "username", "pwd")

		files := NewPgFiles(db)
		readerMock := strings.NewReader("")
		fileToSave, err := file.Create(readerMock, "filepath", 10, u.ID)
		require.NoError(t, err)

		saved, err := files.Save(ctx, *fileToSave)

		assert.NoError(t, err)
		assert.Equal(t, fileToSave.Name, saved.Name)
		assert.Equal(t, fileToSave.Path, saved.Path)
		assert.Equal(t, fileToSave.Size, saved.Size)

		fileToSave.Name = "changed"
		fileToSave.Path = "changed"
		fileToSave.Size = 20

		saved, err = files.Save(ctx, *fileToSave)

		assert.NoError(t, err)
		assert.Equal(t, fileToSave.Name, saved.Name)
		assert.Equal(t, fileToSave.Path, saved.Path)
		assert.Equal(t, fileToSave.Size, saved.Size)
	})

	t.Run("do not saves with a not existing user", func(t *testing.T) {
		ctx, done := testutil.IntegrationTest(context.Background(), t, testutil.WithPgDB())
		defer done()

		db := testutil.PgDBFromContext(ctx, t)
		u, err := user.Create("username", "password")
		require.NoError(t, err)

		files := NewPgFiles(db)
		readerMock := strings.NewReader("")
		fileToSave, err := file.Create(readerMock, "filepath", 10, u.ID)
		require.NoError(t, err)

		saved, err := files.Save(ctx, *fileToSave)

		assert.Error(t, err)
		assert.Nil(t, saved)
	})
}

func TestPgFiles_ByOwnerByPath(t *testing.T) {
	t.Run("returns an existing file", func(t *testing.T) {
		ctx, done := testutil.IntegrationTest(context.Background(), t, testutil.WithPgDB())
		defer done()

		db := testutil.PgDBFromContext(ctx, t)
		u := insertPgUser(t, db, "username", "pwd")

		files := NewPgFiles(db)

		readerMock := strings.NewReader("")
		file, err := file.Create(readerMock, "filepath", 10, u.ID)
		require.NoError(t, err)

		file, err = files.Save(ctx, *file)
		require.NoError(t, err)

		found, err := files.ByOwnerByPath(ctx, u.ID, file.Path)
		assert.NoError(t, err)
		assert.Equal(t, file.ID, found.ID)
		assert.Equal(t, file.Name, found.Name)
		assert.Equal(t, file.Path, found.Path)
		assert.Equal(t, file.Size, found.Size)
		assert.Equal(t, file.OwnerID, found.OwnerID)
	})

	t.Run("returns nil if the file doesn't exist", func(t *testing.T) {
		ctx, done := testutil.IntegrationTest(context.Background(), t, testutil.WithPgDB())
		defer done()

		db := testutil.PgDBFromContext(ctx, t)
		u := insertPgUser(t, db, "username", "pwd")

		files := NewPgFiles(db)

		found, err := files.ByOwnerByPath(ctx, u.ID, "some/path")
		assert.NoError(t, err)
		assert.Nil(t, found)
	})
}

func insertPgUser(t *testing.T, db *sqlx.DB, username string, password string) user.User {
	t.Helper()

	users := userinfra.NewPgUsers(db)
	u := testutil.Must(user.Create(username, password))

	if _, err := users.Save(context.Background(), *u); err != nil {
		t.Fatal(err)
	}

	return *u
}