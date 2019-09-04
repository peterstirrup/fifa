package player

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_GetPositionStr(t *testing.T) {
	table := []struct {
		p   Player
		exp string
	}{
		{Player{Position:Goalkeeper}, "GK"},
		{Player{Position:CentreBack}, "CB"},
		{Player{Position:LeftBack}, "LB"},
		{Player{Position:RightBack}, "RB"},
		{Player{Position:DefensiveMid}, "CDM"},
		{Player{Position:CentreMid}, "CM"},
		{Player{Position:AttackingMid}, "CAM"},
		{Player{Position:LeftMid}, "LM"},
		{Player{Position:RightMid}, "RM"},
		{Player{Position:Striker}, "ST"},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf(tt.exp, tt.p.GetPositionStr(), tt.p), func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, tt.p.GetPositionStr())
		})
	}
}
