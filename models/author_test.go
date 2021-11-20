package models_test

import (
	"booklogger/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAuthor(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []struct {
		input     string
		surname   string
		forenames string
	}{
		{"Karl Marx", "Marx", "Karl"},
		{"Karl Heinrich Marx", "Marx", "Karl Heinrich"},
		{"Ursula K. le Guin", "le Guin", "Ursula K."},
		{"Ursula K. Le Guin", "Le Guin", "Ursula K."},
		{"Miguel A. de la Torre", "de la Torre", "Miguel A."},
		{"J.R.R. Tolkien", "Tolkien", "J.R.R."},
		{"J. R. R. Tolkien", "Tolkien", "J.R.R."},
		{"George R.R. Martin", "Martin", "George R.R."},
		{"George R. R. Martin", "Martin", "George R.R."},
	}

	for _, test := range tests {
		test := test
		t.Run(test.input, func(t *testing.T) {
			t.Parallel()
			author := models.NewAuthor(test.input)

			assert.Equal(test.surname, author.Surname)
			assert.Equal(test.forenames, author.Forenames)
		})
	}
}

func TestDisplayAuthor(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []struct {
		name     string
		author   models.Author
		expected string
	}{
		{"with initials", models.Author{Surname: "Tolkien", Forenames: "J.R.R."}, "J.R.R. Tolkien"},
		{
			"with preferred name",
			models.Author{
				Surname:            "Tolkien",
				PreferredForenames: "J.R.R.",
				Forenames:          "John Ronald Reuel",
			},
			"J.R.R. Tolkien",
		},
		{
			"surname first",
			models.Author{Surname: "Mao", Forenames: "Zedong", SurnameFirst: true},
			"Mao Zedong",
		},
		{"single name", models.Author{Surname: "Apple"}, "Apple"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(test.expected, test.author.String(), test.name)
		})
	}
}
