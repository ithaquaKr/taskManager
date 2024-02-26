package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestNoteRepository_Create(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	noteRepo := NewNoteRepo(sqlxDB)

	t.Run("success", func(t *testing.T) {
		test_listID := uuid.New()
		test_name := "test name"
		test_content := "test content"
		rows := sqlmock.NewRows([]string{"list_id", "name", "content"}).AddRow(test_listID, test_name, test_content)

		notes := &models.Note{
			ListID:  test_listID,
			Name:    test_name,
			Content: test_content,
		}
		mock.ExpectQuery(createNote).WithArgs(test_listID, test_name, test_content).WillReturnRows(rows)
		createNote, err := noteRepo.CreateNote(context.Background(), notes)
		require.NoError(t, err)
		require.NotNil(t, createNote)
		require.Equal(t, test_listID, createNote.ID)
		require.Equal(t, test_name, createNote.Content)
		require.Equal(t, test_content, createNote.Content)
	})
}

func TestNoteRepository_Update(t *testing.T) {
	t.Parallel()
}

func TestNoteRepository_Delete(t *testing.T) {
	t.Parallel()
}

func TestNoteRepository_Get(t *testing.T) {
	// Implement this
}

func TestNoteRepository_All(t *testing.T) {
	// Implement this
}
