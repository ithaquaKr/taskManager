package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/list/entities"
	"github.com/ithaquaKr/taskManager/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestListRepositoryCreate(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	listRepo := NewListRepo(sqlxDB)

	// Test create list repository success
	t.Run("repository_list_create_success", func(t *testing.T) {
		title_test := "test title"
		type_test := "task"

		rows := sqlmock.NewRows([]string{"title", "type"}).AddRow(title_test, type_test)

		lists := &entities.List{
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

func TestListRepositoryUpdate(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	defer sqlxDB.Close()

	listRepo := NewListRepo(sqlxDB)
	// Test update list repository success
	t.Run("repository_list_update_success", func(t *testing.T) {
		test_listID := uuid.New()
		test_title := "update title"
		test_type := "note"

		rows := sqlmock.NewRows([]string{"id", "title", "type"}).AddRow(test_listID, test_title, test_type)
		list := &entities.List{
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

func TestListRepositoryDelete(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	listRepo := NewListRepo(sqlxDB)
	// Test delete list repository success
	t.Run("repository_list_delete_success", func(t *testing.T) {
		test_listID := uuid.New()
		mock.ExpectExec(deleteList).WithArgs(test_listID).WillReturnResult(sqlmock.NewResult(1, 1))
		err := listRepo.DeleteList(context.Background(), test_listID)
		require.NoError(t, err)
	})
}

func TestListRepositoryGet(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	listRepo := NewListRepo(sqlxDB)

	test_listID := uuid.New()
	test_title := "test title"
	test_type := "task"

	rows := sqlmock.NewRows([]string{"id", "title", "type"}).AddRow(test_listID, test_title, test_type)
	// Test get list repository success
	t.Run("repository_list_get_success", func(t *testing.T) {
		mock.ExpectQuery(getList).WithArgs(test_listID).WillReturnRows(rows)
		getList, err := listRepo.GetList(context.Background(), test_listID)
		require.NoError(t, err)
		require.NotNil(t, getList)
		require.Equal(t, test_listID, getList.ID)
		require.Equal(t, test_title, getList.Title)
		require.Equal(t, test_type, getList.Type)
	})
	// Test get list repository failed
	t.Run("repository_list_get_no_row", func(t *testing.T) {
		mock.ExpectQuery(getList).WithArgs(uuid.New()).WillReturnError(sql.ErrNoRows)
		getList, err := listRepo.GetList(context.Background(), test_listID)
		require.Error(t, err)
		require.Nil(t, getList)
	})
}
func TestListRepositoryAll(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	listRepo := NewListRepo(sqlxDB)

	lists_data := [][]driver.Value{
		{
			uuid.New(),
			"test title 1",
			"task",
		},
		{
			uuid.New(),
			"test title 2",
			"task",
		},
		{
			uuid.New(),
			"test title 3",
			"note",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "type"}).AddRows(lists_data...)
	count_row := sqlmock.NewRows([]string{"count"}).AddRow(len(lists_data))
	// Test get all list repository success
	t.Run("repository_list_get_all_success", func(t *testing.T) {
		var pq utils.PaginationQuery
		pq.SetPageNumber("1")
		pq.SetPageSize("10")
		mock.ExpectQuery(getTotal).WillReturnRows(count_row)
		mock.ExpectQuery(allLists).WithArgs(pq.GetLimit(), pq.GetOffset()).WillReturnRows(rows)
		allLists, err := listRepo.AllLists(context.Background(), &pq)
		require.NoError(t, err)
		require.NotNil(t, allLists)
		require.Equal(t, 3, allLists.Paginate.TotalCount)
		require.Equal(t, 1, allLists.Paginate.TotalPage)
		require.Equal(t, false, allLists.Paginate.HasMore)
		require.Equal(t, 2, allLists.Paginate.NextPage)
		require.Equal(t, 0, allLists.Paginate.PreviousPage)
		require.Equal(t, 3, len(allLists.Result))
	})
}
