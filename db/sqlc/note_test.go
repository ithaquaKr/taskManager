package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomNote(t *testing.T, list List) Note {
	arg := CreateNoteParams{
		ListID:   list.ID,
		Title:    gofakeit.Sentence(10),
		Content:  sql.NullString{String: gofakeit.Sentence(50), Valid: true},
		Reminder: sql.NullTime{Time: gofakeit.Date()},
	}
	note, err := testQueries.CreateNote(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, note)

	require.Equal(t, arg.ListID, note.ListID)
	require.Equal(t, arg.Title, note.Title)
	require.Equal(t, arg.Content, note.Content)
	// require.WithinDuration(t, arg.Reminder.Time, note.Reminder.Time, time.Second)

	return note
}

func TestCreateNote(t *testing.T) {
	list := createRandomList(t)
	createRandomNote(t, list)
}

func TestGetNote(t *testing.T) {
	list := createRandomList(t)
	note1 := createRandomNote(t, list)
	note2, err := testQueries.GetNote(context.Background(), note1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, note2)
}
