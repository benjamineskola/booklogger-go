package models_test

import (
	"booklogger/models"
	"testing"
)

func TestCreateAuthor(t *testing.T) { //nolint:cyclop
	t.Parallel()
	t.Run("Creating an author parses the name out", func(t *testing.T) {
		t.Parallel()
		author := models.NewAuthor("Karl Marx")
		if author.Surname != "Marx" {
			t.Fatal("surname not set")
		}
		if author.Forenames != "Karl" {
			t.Fatal("forenames not set")
		}
	})

	t.Run("Creating an author parses the name out including middle names", func(t *testing.T) {
		t.Parallel()
		author := models.NewAuthor("Karl Heinrich Marx")
		if author.Surname != "Marx" {
			t.Fatal("surname not set")
		}
		if author.Forenames != "Karl Heinrich" {
			t.Fatal("forenames not set")
		}
	})

	t.Run("Surname prefixes are included in the surname", func(t *testing.T) {
		t.Parallel()
		author := models.NewAuthor("Ursula K. le Guin")
		if author.Surname != "le Guin" {
			t.Fatal("surname not set")
		}
		if author.Forenames != "Ursula K." {
			t.Fatal("forenames not set")
		}
	})

	t.Run("A surname can have multiple prefixes", func(t *testing.T) {
		t.Parallel()
		author := models.NewAuthor("Miguel A. de la Torre")
		if author.Surname != "de la Torre" {
			t.Fatal("surname not set correctly")
		}
		if author.Forenames != "Miguel A." {
			t.Fatal("forenames not set correctly")
		}
	})

	t.Run("Names presented as initials are normalised", func(t *testing.T) {
		t.Parallel()
		author := models.NewAuthor("J. R. R. Tolkien")
		if author.Forenames != "J.R.R." {
			t.Fatal("forenames divided incorrectly")
		}
		author = models.NewAuthor("George R. R. Martin")
		if author.Forenames != "George R.R." {
			t.Fatal("forenames divided incorrectly")
		}
	})
}
