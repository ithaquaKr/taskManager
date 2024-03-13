package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/task/entities"
	"github.com/ithaquaKr/taskManager/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestTaskRepositoryCreate(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	taskRepo := NewTaskRepo(sqlxDB)

	t.Run("repository_task_create_success", func(t *testing.T) {
		id_test := uuid.New()
		listID_test := uuid.New()
		name_test := "test name"
		description_test := "test description"
		status_test := "test status"
		tag_test := "test tag"
		priority_test := "no_priority"
		dueDate_test := time.Now()
		rows := sqlmock.NewRows([]string{
			"id", "list_id", "name", "description", "status", "tag", "priority", "due_date",
		}).AddRow(id_test, listID_test, name_test, description_test, status_test, tag_test, priority_test, dueDate_test)

		tasks := &entities.Task{
			ListID:      listID_test,
			Name:        name_test,
			Description: &description_test,
			Status:      status_test,
			Tag:         &tag_test,
			Priority:    priority_test,
			DueDate:     &dueDate_test,
		}
		mock.ExpectQuery(createTask).WithArgs(
			listID_test, name_test, &tasks.Description, status_test, &tasks.Tag, priority_test, &tasks.DueDate,
		).WillReturnRows(rows)
		createTask, err := taskRepo.CreateTask(context.Background(), tasks)
		require.NoError(t, err)
		require.NotNil(t, createTask)
		require.Equal(t, id_test, createTask.ID)
		require.Equal(t, listID_test, createTask.ListID)
		require.Equal(t, name_test, createTask.Name)
		require.Equal(t, &description_test, createTask.Description)
		require.Equal(t, status_test, createTask.Status)
		require.Equal(t, &tag_test, createTask.Tag)
		require.Equal(t, priority_test, createTask.Priority)
		require.Equal(t, &dueDate_test, createTask.DueDate)
	})
}

// func TestTaskRepositoryUpdate(t *testing.T) {
// 	t.Parallel()
// 	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 	require.NoError(t, err)
// 	defer db.Close()

// 	sqlxDB := sqlx.NewDb(db, "sqlmock")
// 	defer sqlxDB.Close()

// 	taskRepo := NewTaskRepo(sqlxDB)

// 	// Create a task
// 	old_task := entities.Task{
// 		ListID:   uuid.New(),
// 		Name:     "test_update_name",
// 		Status:   "doing",
// 		Priority: "no_priority",
// 	}

// 	t.Run("update_task_repository_success", func(t *testing.T) {
// 		listID_test := uuid.New()
// 		name_test := "test name"
// 		description_test := "test description"
// 		status_test := "test status"
// 		tag_test := "test tag"
// 		priority_test := "no_priority"
// 		dueDate_test := time.Now()

// 	})

// }

func TestTaskRepositoryDelete(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	taskRepo := NewTaskRepo(sqlxDB)

	t.Run("repository_task_delete_success", func(t *testing.T) {
		test_taskID := uuid.New()
		mock.ExpectExec(deleteTask).WithArgs(test_taskID).WillReturnResult(sqlmock.NewResult(1, 1))
		err := taskRepo.DeleteTask(context.Background(), test_taskID)
		require.NoError(t, err)
	})

	t.Run("repository_task_delete_no_row", func(t *testing.T) {
		test_taskID := uuid.New()
		mock.ExpectExec(deleteTask).WithArgs(test_taskID).WillReturnResult(sqlmock.NewResult(1, 0))
		err := taskRepo.DeleteTask(context.Background(), test_taskID)
		require.Error(t, err)
	})
}

func TestTaskRepositoryGet(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	test_taskID := uuid.New()
	test_listID := uuid.New()
	test_name := "test_task"
	test_description := "test_description"
	test_status := "doing"
	test_tag := "test_tag"
	test_priority := "high"
	test_due_date := time.Now()

	taskRepo := NewTaskRepo(sqlxDB)

	rows := sqlmock.NewRows([]string{
		"id", "list_id", "name", "description", "status", "tag", "priority", "due_date",
	}).AddRow(
		test_taskID,
		test_listID,
		test_name,
		test_description,
		test_status,
		test_tag,
		test_priority,
		test_due_date,
	)

	t.Run("repository_task_get_success", func(t *testing.T) {
		mock.ExpectQuery(getTask).WithArgs(test_taskID).WillReturnRows(rows)
		getTask, err := taskRepo.GetTask(context.Background(), test_taskID)
		require.NoError(t, err)
		require.NotNil(t, getTask)
		require.Equal(t, test_taskID, getTask.ID)
		require.Equal(t, test_listID, getTask.ListID)
		require.Equal(t, test_name, getTask.Name)
		require.Equal(t, &test_description, getTask.Description)
		require.Equal(t, test_status, getTask.Status)
		require.Equal(t, &test_tag, getTask.Tag)
		require.Equal(t, test_priority, getTask.Priority)
		require.Equal(t, &test_due_date, getTask.DueDate)
	})

	t.Run("repository_task_get_no_row", func(t *testing.T) {
		mock.ExpectQuery(getTask).WithArgs(test_taskID).WillReturnError(sql.ErrNoRows)
		getTask, err := taskRepo.GetTask(context.Background(), test_taskID)
		require.Error(t, err)
		require.Nil(t, getTask)
	})
}

func TestTaskRepositoryAll(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	taskRepo := NewTaskRepo(sqlxDB)

	tasks_data := [][]driver.Value{
		{
			uuid.New(),
			uuid.New(),
			"test_task_1",
			"test_description_1",
			"doing",
			"test_tag_1",
			"high",
			time.Now(),
		},
		{
			uuid.New(),
			uuid.New(),
			"test_task_2",
			"test_description_2",
			"doing",
			"test_tag_2",
			"no_priority",
			time.Now(),
		},
	}
	rows := sqlmock.NewRows([]string{
		"id", "list_id", "name", "description", "status", "tag", "priority", "due_date",
	}).AddRows(tasks_data...)
	count_row := sqlmock.NewRows([]string{"count"}).AddRow(2)

	t.Run("repository_task_get_all_success", func(t *testing.T) {
		var pq utils.PaginationQuery
		pq.SetPageNumber("1")
		pq.SetPageSize("10")
		mock.ExpectQuery(getTotal).WillReturnRows(count_row)
		mock.ExpectQuery(allTasks).WithArgs(pq.GetLimit(), pq.GetOffset()).WillReturnRows(rows)
		allNotes, err := taskRepo.AllTasks(context.Background(), &pq)
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
