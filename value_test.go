package fifa

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_GetValue(t *testing.T) {
	table := []struct {
		position int
		overall  float64
		exp      float64
	}{
		{Goalkeeper, 86, 20.64},
		{CentreBack, 89, 26.7},
		{RightBack, 76, 4.56},
		{CentreMid, 84, 12.69},
		{AttackingMid, 87, 36.54},
		{LeftMid, 92, 58.88},
		{Striker, 81, 24.48},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("position: %v, overall: %v", tt.position, tt.overall), func(st *testing.T) {
			p := Player{
				Position: tt.position,
				Overall:  tt.overall,
			}
			p.GetValue()

			r := require.New(st)
			r.Equal(tt.exp, p.Value)
		})
	}
}

func TestPlayer_floatToTwoDecimalPlaces(t *testing.T) {
	table := []struct {
		tc  float64
		exp float64
	}{
		{2.699999, 2.7},
		{-3.28734, -3.29},
		{13281.23823901, 13281.24},
		{-39132.3281239, -39132.33},
		{0.542, 0.54},
		{0, 0},
		{0.004, 0},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%v", tt.tc), func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, floatToTwoDecimalPlaces(tt.tc))
		})
	}
}
