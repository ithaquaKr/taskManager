package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/note/entities"
	"github.com/ithaquaKr/taskManager/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestNoteRepositoryCreate(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	noteRepo := NewNoteRepo(sqlxDB)

	t.Run("repository_note_create_success", func(t *testing.T) {
		test_listID := uuid.New()
		test_name := "test name"
		test_content := "test content"
		rows := sqlmock.NewRows([]string{"list_id", "name", "content"}).AddRow(test_listID, test_name, test_content)

		notes := &entities.Note{
			ListID:  test_listID,
			Name:    test_name,
			Content: test_content,
		}
		mock.ExpectQuery(createNote).WithArgs(test_listID, test_name, test_content).WillReturnRows(rows)
		createNote, err := noteRepo.CreateNote(context.Background(), notes)
		require.NoError(t, err)
		require.NotNil(t, createNote)
		require.Equal(t, test_listID, createNote.ListID)
		require.Equal(t, test_name, createNote.Name)
		require.Equal(t, test_content, createNote.Content)
	})
}

func TestNoteRepositoryUpdate(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	noteRepo := NewNoteRepo(sqlxDB)

	t.Run("repository_note_update_success", func(t *testing.T) {
		test_listID := uuid.New()
		test_name := "updated name"
		test_content := "updated content"
		rows := sqlmock.NewRows([]string{"list_id", "name", "content"}).AddRow(test_listID, test_name, test_content)
		notes := &entities.Note{
			ListID:  test_listID,
			Name:    test_name,
			Content: test_content,
		}

		mock.ExpectQuery(updateNote).WithArgs(test_listID, test_name, test_content).WillReturnRows(rows)

		updateNote, err := noteRepo.UpdateNote(context.Background(), notes)
		require.NoError(t, err)
		require.NotNil(t, updateNote)
	})
}

func TestNoteRepositoryDelete(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	noteRepo := NewNoteRepo(sqlxDB)

	t.Run("repository_note_delete_success", func(t *testing.T) {
		test_noteID := uuid.New()
		mock.ExpectExec(deleteNote).WithArgs(test_noteID).WillReturnResult(sqlmock.NewResult(1, 1))
		err := noteRepo.DeleteNote(context.Background(), test_noteID)
		require.NoError(t, err)
	})
}

func TestNoteRepositoryGet(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	noteRepo := NewNoteRepo(sqlxDB)

	test_noteID := uuid.New()
	test_listID := uuid.New()
	test_name := "test name"
	test_content := "test content"

	rows := sqlmock.NewRows([]string{"id", "list_id", "name", "content"}).AddRow(test_noteID, test_listID, test_name, test_content)
	t.Run("repository_note_get_success", func(t *testing.T) {
		mock.ExpectQuery(getNote).WithArgs(test_noteID).WillReturnRows(rows)
		getNote, err := noteRepo.GetNote(context.Background(), test_noteID)
		require.NoError(t, err)
		require.NotNil(t, getNote)
		require.Equal(t, test_noteID, getNote.ID)
		require.Equal(t, test_listID, getNote.ListID)
		require.Equal(t, test_name, getNote.Name)
		require.Equal(t, test_content, getNote.Content)
	})

	t.Run("repository_note_get_no_row", func(t *testing.T) {
		mock.ExpectQuery(getNote).WithArgs(test_noteID).WillReturnError(sql.ErrNoRows)
		getNote, err := noteRepo.GetNote(context.Background(), test_noteID)
		require.Error(t, err)
		require.Nil(t, getNote)
	})
}

func TestNoteRepositoryAll(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	noteRepo := NewNoteRepo(sqlxDB)

	notes_data := [][]driver.Value{
		{uuid.New(), uuid.New(), "test name", "test content"},
		{uuid.New(), uuid.New(), "test name 2", "test content 2"},
	}
	rows := sqlmock.NewRows([]string{"id", "list_id", "name", "content"}).AddRows(notes_data...)
	count_row := sqlmock.NewRows([]string{"count"}).AddRow(2)

	// Test get all notes success
	t.Run("repository_note_get_all_success", func(t *testing.T) {
		var pq utils.PaginationQuery
		pq.SetPageNumber("1")
		pq.SetPageSize("10")
		mock.ExpectQuery(getTotal).WillReturnRows(count_row)
		mock.ExpectQuery(allNotes).WithArgs(pq.GetLimit(), pq.GetOffset()).WillReturnRows(rows)
		allNotes, err := noteRepo.AllNotes(context.Background(), &pq)
		require.NoError(t, err)
		require.NotNil(t, allNotes)
		require.Equal(t, 2, allNotes.Paginate.TotalCount)
		require.Equal(t, 1, allNotes.Paginate.TotalPage)
		require.Equal(t, false, allNotes.Paginate.HasMore)
		require.Equal(t, 2, allNotes.Paginate.NextPage)
		require.Equal(t, 0, allNotes.Paginate.PreviousPage)
		require.Equal(t, 2, len(allNotes.Result))
	})
}
