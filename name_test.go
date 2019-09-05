package fifa

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_genName(t *testing.T) {
	table := []struct {
		countryID int
		c         country
		expFirst  string
		expSecond string
	}{
		{100, country{normalName: 1, firstNames: []string{"John"}, secondNames: []string{"Doe"}}, "John", "Doe"},
		{101, country{doubleBarrel: 1, firstNames: []string{"John"}, secondNames: []string{"Doe"}}, "John", "Doe-Doe"},
		{102, country{middleName: 1, firstNames: []string{"John"}, secondNames: []string{"Doe"}}, "John John", "Doe"},
		{103, country{oneName: 1, firstNames: []string{"John"}, secondNames: []string{"Doe"}}, "", "John"},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%+v", tt.c), func(st *testing.T) {
			countries[tt.countryID] = tt.c
			p := Player{
				Nationality: tt.countryID,
			}
			p.genName()

			r := require.New(st)
			r.Equal(tt.expFirst, p.FirstName)
			r.Equal(tt.expSecond, p.SecondName)

			// clean up
			countries[tt.countryID] = country{}
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
