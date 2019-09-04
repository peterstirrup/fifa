package player

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_GetOverall(t *testing.T) {
	table := []struct {
		p   Player
	}{
		{playerTemplate[Goalkeeper]},
		{playerTemplate[CentreBack]},
		{playerTemplate[LeftBack]},
		{playerTemplate[RightBack]},
		{playerTemplate[DefensiveMid]},
		{playerTemplate[CentreMid]},
		{playerTemplate[AttackingMid]},
		{playerTemplate[LeftMid]},
		{playerTemplate[RightMid]},
		{playerTemplate[Striker]},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s: %v", tt.p.GetPositionStr(), tt.p.Overall), func(st *testing.T) {
			exp := tt.p.Overall
			tt.p.GetOverall()

			r := require.New(st)
			r.Equal(exp, tt.p.Overall)
		})
	}
}
