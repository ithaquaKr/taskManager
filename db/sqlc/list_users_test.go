package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomListUsers(t *testing.T, list List, user User) ListUser {
	arg := CreateListUserParams{
		ListID: uuid.NullUUID{UUID: list.ID, Valid: true},
		UserID: uuid.NullUUID{UUID: user.ID, Valid: true},
	}
	list_user, err := testQueries.CreateListUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, list_user)

	return list_user
}

func TestCreateListUser(t *testing.T) {
	list := createRandomList(t)
	user := createRandomUser(t)
	createRandomListUsers(t, list, user)
}

func TestGetListUser(t *testing.T) {
	list := createRandomList(t)
	user := createRandomUser(t)
	list_user1 := createRandomListUsers(t, list, user)
	list_user2, err := testQueries.GetListUser(context.Background(), list_user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, list_user2)
}
