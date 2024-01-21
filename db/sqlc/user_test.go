package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:     gofakeit.Username(),
		HashPassword: gofakeit.Password(true, true, true, true, false, 16),
		Email:        gofakeit.Email(),
		FirstName:    gofakeit.FirstName(),
		LastName:     gofakeit.LastName(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetAccount(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}
