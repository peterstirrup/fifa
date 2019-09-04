package player

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_GenPlayer(t *testing.T) {
	p := GenPlayer()

	table := []struct {
		attr string
		tc   interface{}
	}{
		{"Nationality", p.Nationality},
		{"Name", p.SecondName},
	}

	for _, tt := range table {
		t.Run(tt.attr, func(st *testing.T) {
			r := require.New(st)
			r.NotZero(tt.tc)
		})
	}
}

func TestPlayer_genStats(t *testing.T) {
	p := Player{
		Position: Striker,
	}
	p.genStats()

	table := []struct {
		attr string
		tc   interface{}
	}{
		{"Crossing", p.Crossing},
		{"Shot Power", p.ShotPower},
		{"Diving", p.Diving},
	}

	for _, tt := range table {
		t.Run(tt.attr, func(st *testing.T) {
			r := require.New(st)
			r.NotZero(tt.tc)
		})
	}
}

func TestPlayer_GenPlayerBase(t *testing.T) {
	table := []struct {
		n   int
		exp int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{50, 50},
		{-1, 0},
		{-100, 0},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("generate %d players", tt.n), func(st *testing.T) {
			players, err := GenPlayerBase(tt.n)
			if err != nil {
				st.Fatal(err)
			}

			r := require.New(st)
			r.Equal(len(players), tt.exp)
		})
	}
}
