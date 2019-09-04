package fifa

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_randItemFromSlice(t *testing.T) {
	table := []struct {
		m   map[int]float64
		exp int
	}{
		{map[int]float64{1: 1, 2: 0, 3: 0}, 1},
		{map[int]float64{1: 0, 2: 1, 3: 0}, 2},
		{map[int]float64{1: 0, 2: 0, 3: 1}, 3},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%d", tt.exp), func(st *testing.T) {
			r := require.New(st)
			r.Equal(randItemFromSlice(tt.m), tt.exp)
		})
	}
}

func TestPlayer_randRange(t *testing.T) {
	table := []struct {
		min int
		max int
	}{
		{0, 1},
		{0, 2},
		{1, 3},
		{-1, 5},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("min: %d, max: %d", tt.min, tt.max), func(st *testing.T) {
			if n := randRange(tt.min, tt.max); n < tt.min || n > tt.max {
				t.Fail()
			}
		})
	}
}

func TestPlayer_capNum(t *testing.T) {
	table := []struct {
		n   float64
		min float64
		max float64
	}{
		{0.5, 0, 1},
		{0, 0, 2},
		{3, 1, 3},
		{-1, -1, 5},
		{100, -1, 5},
		{-100, -1, 5},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("min: %v, max: %v", tt.min, tt.max), func(st *testing.T) {
			if n := capNum(tt.n, tt.min, tt.max); n < tt.min || n > tt.max {
				t.Fail()
			}
		})
	}
}

func TestPlayer_weightedRand(t *testing.T) {
	repeatRand = 3.0

	t.Run(fmt.Sprintf("talentCo = 0"), func(st *testing.T) {
		if n := weightedRand(1, 0); n > 0 {
			t.Fail()
		}
	})
	t.Run(fmt.Sprintf("talentCo = 1"), func(st *testing.T) {
		if n := weightedRand(1, 1); n < 0 {
			t.Fail()
		}
	})
}

func TestPlayer_posWeightedRand(t *testing.T) {
	table := []struct {
		max int
	}{
		{0},
		{1},
		{2},
		{3},
		{40},
		{500},
	}

	repeatRand = 3.0

	for _, tt := range table {
		t.Run(fmt.Sprintf("max: %v", tt.max), func(st *testing.T) {
			if n := posWeightedRand(tt.max); n > float64(tt.max) {
				t.Fail()
			}
		})
	}
}

func TestPlayer_randNegative(t *testing.T) {
	for i := 0; i < 100; i++ {
		if randNegative() > 0 {
			t.Fail()
		}
	}
}
