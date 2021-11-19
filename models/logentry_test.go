package models_test

import (
	"booklogger/models"
	"testing"
	"time"
)

func TestLogEntry(t *testing.T) {
	t.Parallel()

	book := models.NewBook("Das Kapital")
	book.FirstAuthor = *models.NewAuthor("Karl Marx")

	t.Run("Entry with no dates", func(t *testing.T) {
		t.Parallel()
		entry := models.LogEntry{Book: *book}
		if entry.String() != "Karl Marx, Das Kapital, unfinished" {
			t.Fatal("Entry does not match", entry.String())
		}
	})

	t.Run("Entry with a start date and no end date", func(t *testing.T) {
		t.Parallel()
		entry := models.LogEntry{Book: *book}
		entry.StartNow()
		if entry.String() != "Karl Marx, Das Kapital, from "+time.Now().
			Format("2006-01-02")+
			", unfinished" {
			t.Fatal("Entry does not match:", entry.String())
		}
	})

	t.Run("Entry with a start date and end date", func(t *testing.T) {
		t.Parallel()
		entry := models.LogEntry{Book: *book}
		start, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
		entry.StartAt(&start)
		entry.EndNow()
		if entry.String() != "Karl Marx, Das Kapital, from 2012-11-01, to "+time.Now().
			Format("2006-01-02") {
			t.Fatal("Entry does not match:", entry.String())
		}
	})
}
