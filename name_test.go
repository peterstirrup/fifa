package fifa

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_genName(t *testing.T) {
	table := []struct {
		countryID int
		c         Country
		expFirst  string
		expSecond string
	}{
		{100, Country{NormalName: 1, FirstNames: []string{"John"}, SecondNames: []string{"Doe"}}, "John", "Doe"},
		{101, Country{DoubleBarrel: 1, FirstNames: []string{"John"}, SecondNames: []string{"Doe"}}, "John", "Doe-Doe"},
		{102, Country{MiddleName: 1, FirstNames: []string{"John"}, SecondNames: []string{"Doe"}}, "John John", "Doe"},
		{103, Country{OneName: 1, FirstNames: []string{"John"}, SecondNames: []string{"Doe"}}, "", "John"},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%+v", tt.c), func(st *testing.T) {
			Countries[tt.countryID] = tt.c
			p := Player{
				Nationality: tt.countryID,
			}
			p.genName()

			r := require.New(st)
			r.Equal(tt.expFirst, p.FirstName)
			r.Equal(tt.expSecond, p.SecondName)

			// clean up
			Countries[tt.countryID] = Country{}
		})
	}
}

func TestPlayer_GetShortenedFirstName(t *testing.T) {
	table := []struct {
		act string
		exp string
	}{
		{"John", "J"},
		{"Peter", "P"},
		{"Doelinho", "D"},
		{"", ""},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf(tt.act), func(st *testing.T) {
			p := Player{
				FirstName: tt.act,
			}

			r := require.New(st)
			r.Equal(tt.exp, p.GetShortenedFirstName())
		})
	}
}
