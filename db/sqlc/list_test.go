package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ithaquaKr/taskManager/utils"
	"github.com/stretchr/testify/require"
)

func createRandomList(t *testing.T) List {
	arg := CreateListParams{
		Name: gofakeit.Word(),
		Type: utils.RandomType(),
	}
	list, err := testQueries.CreateList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, list)

	return list
}

func TestCreateList(t *testing.T) {
	createRandomList(t)
}

func TestGetList(t *testing.T) {
	list1 := createRandomList(t)
	list2, err := testQueries.GetList(context.Background(), list1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, list2)

	require.Equal(t, list1.ID, list2.ID)
	require.Equal(t, list1.Name, list2.Name)
	require.Equal(t, list1.Type, list2.Type)
	require.WithinDuration(t, list1.CreatedAt.Time, list2.CreatedAt.Time, time.Second)
}

func TestListList(t *testing.T) {
}
