package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/task/entities"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestTaskRepository_Create(t *testing.T) {
	// TODO: add due date to test
	t.Parallel()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	taskRepo := NewTaskRepo(sqlxDB)

	t.Run("success", func(t *testing.T) {
		listID_test := uuid.New()
		name_test := "test name"
		description_test := "test description"
		status_test := "test status"
		tag_test := "test tag"
		priority_test := "no_priority"
		// dueDate_test := time.Now()
		rows := sqlmock.NewRows([]string{
			"list_id", "name", "description", "status", "tag", "priority",
		}).AddRow(listID_test, name_test, description_test, status_test, tag_test, priority_test)

		tasks := &entities.Task{
			ListID: listID_test,
			Name:   name_test,
			// Description: "Dest",
			Status: status_test,
			// Tag:         "tag",
			Priority: priority_test,
		}
		mock.ExpectQuery(createTask).WithArgs(
			listID_test, name_test, &tasks.Description, status_test, &tasks.Tag, priority_test, &tasks.DueDate,
		).WillReturnRows(rows)
		createTask, err := taskRepo.CreateTask(context.Background(), tasks)
		require.NoError(t, err)
		require.NotNil(t, createTask)
		require.Equal(t, listID_test, createTask.ListID)
		require.Equal(t, name_test, createTask.Name)
		// require.Equal(t, description_test, createTask.Description)
		require.Equal(t, status_test, createTask.Status)
		// require.Equal(t, tag_test, createTask.Tag)
		require.Equal(t, priority_test, createTask.Priority)
	})
}

func TestTaskRepository_Update(t *testing.T) {
	// Implement this test
}

func TestTaskRepository_Delete(t *testing.T) {
	// Implement this test
}

func TestTaskRepository_Get(t *testing.T) {
	// Implement this test
}

func TestTaskRepository_All(t *testing.T) {
	// Implement this test
}
