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

func TestListRepository_Create(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	listRepo := NewListRepo(sqlxDB)

	t.Run("success", func(t *testing.T) {
		title_test := "test title"
		type_test := "task"

		rows := sqlmock.NewRows([]string{"title", "type"}).AddRow(title_test, type_test)

		lists := &models.List{
			Title: title_test,
			Type:  type_test,
		}
		mock.ExpectQuery(createList).WithArgs(title_test, type_test).WillReturnRows(rows)
		createList, err := listRepo.CreateList(context.Background(), lists)
		require.NoError(t, err)
		require.NotNil(t, createList)
		require.Equal(t, title_test, createList.Title)
		require.Equal(t, type_test, createList.Type)
	})
}

func TestListRepository_Update(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	listRepo := NewListRepo(sqlxDB)

	t.Run("success", func(t *testing.T) {
		test_listID := uuid.New()
		test_title := "update title"
		test_type := "note"

		rows := sqlmock.NewRows([]string{"id", "title", "type"}).AddRow(test_listID, test_title, test_type)
		list := &models.List{
			ID:    test_listID,
			Title: test_title,
			Type:  test_type,
		}
		mock.ExpectQuery(updateList).WithArgs(test_title, test_type, test_listID).WillReturnRows(rows)
		updateList, err := listRepo.UpdateList(context.Background(), list)
		require.NoError(t, err)
		require.NotNil(t, updateList)
		require.Equal(t, test_listID, updateList.ID)
	})
}

func TestListRepository_Get(t *testing.T) {
	// Implement me!
}

func TestListRepository_Delete(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	listRepo := NewListRepo(sqlxDB)

	t.Run("success", func(t *testing.T) {
		test_listID := uuid.New()
		mock.ExpectExec(deleteList).WithArgs(test_listID).WillReturnResult(sqlmock.NewResult(1, 1))
		err := listRepo.DeleteList(context.Background(), test_listID)
		require.NoError(t, err)
	})
}

func TestListRepository_All(t *testing.T) {
	// Implement me!
}
