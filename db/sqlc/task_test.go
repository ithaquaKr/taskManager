package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ithaquaKr/taskManager/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTask(t *testing.T, list List) Task {
	arg := CreateTaskParams{
		ListID:      list.ID,
		Title:       gofakeit.Sentence(5),
		Description: sql.NullString{String: gofakeit.Sentence(10), Valid: true},
		Status:      utils.RandomStatus(),
		Tag:         sql.NullString{String: gofakeit.Noun(), Valid: true},
		Priority:    utils.RandomPriority(),
		DueDate:     sql.NullTime{Time: gofakeit.Date()},
	}
	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.ListID, task.ListID)
	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Description, task.Description)
	require.Equal(t, arg.Status, task.Status)
	require.Equal(t, arg.Tag, task.Tag)
	require.Equal(t, arg.Priority, task.Priority)
	// require.WithinDuration(t, arg.DueDate.Time, task.DueDate.Time, time.Second)

	return task
}

func TestCreateTask(t *testing.T) {
	list := createRandomList(t)
	createRandomTask(t, list)
}

func TestGetTask(t *testing.T) {
	list := createRandomList(t)
	task1 := createRandomTask(t, list)
	task2, err := testQueries.GetTask(context.Background(), task1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)
}
