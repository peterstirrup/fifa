package fifa

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountry_GetNationalityCode(t *testing.T) {
	table := []struct {
		country string
		tc      Player
		exp     string
	}{
		{"Argentina", Player{Nationality: Argentina}, "AR"},
		{"England", Player{Nationality: England}, "GB"},
		{"Italy", Player{Nationality: Italy}, "IT"},
		{"Netherlands", Player{Nationality: Netherlands}, "NL"},
		{"Spain", Player{Nationality: Spain}, "ES"},
		{"Undefined", Player{Nationality: 100}, "NaN"},
	}

	for _, tt := range table {
		t.Run(tt.country, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, tt.tc.GetNationalityCode())
		})
	}
}

func TestCountry_GetNationalityStr(t *testing.T) {
	table := []struct {
		tc  Player
		exp string
	}{
		{Player{Nationality: Argentina}, "Argentina"},
		{Player{Nationality: England}, "England"},
		{Player{Nationality: Italy}, "Italy"},
		{Player{Nationality: Netherlands}, "Netherlands"},
		{Player{Nationality: Spain}, "Spain"},
		{Player{Nationality: 100}, "NaN"},
	}

	for _, tt := range table {
		t.Run(tt.exp, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, tt.tc.GetNationalityStr())
		})
	}
}
