package fifa

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayer_genStrongFoot(t *testing.T) {
	table := []struct {
		foot string
		tc   Player
		exp  int
	}{
		{"Left", Player{Position: LeftBack}, LeftFoot},
		{"Right", Player{Position: RightBack}, RightFoot},
	}

	for _, tt := range table {
		t.Run(tt.foot, func(st *testing.T) {
			tt.tc.genStrongFoot()
			r := require.New(st)
			r.Equal(tt.exp, tt.tc.Foot)
		})
	}
}

func TestPlayer_GetFootStr(t *testing.T) {
	table := []struct {
		tc  Player
		exp string
	}{
		{Player{Foot: LeftFoot}, "Left"},
		{Player{Foot: RightFoot}, "Right"},
	}

	for _, tt := range table {
		t.Run(tt.exp, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, tt.tc.GetFootStr())
		})
	}
}
