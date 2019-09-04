package player

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_genHeight(t *testing.T) {
	SetFreakishlyTallProb(1)

	table := []struct {
		p   Player
		min int
		max int
	}{
		{Player{Position: Goalkeeper}, 183, 201},
		{Player{Position: CentreBack}, 173, 195},
		{Player{Position: RightBack}, 165, 193},
		{Player{Position: CentreMid}, 165, 193},
		{Player{Position: LeftMid}, 165, 193},
		{Player{Position: Striker}, 165, 193},
		{Player{Position: Striker, Heading: 90}, 201, 211},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("player position: %d", tt.p.Position), func(st *testing.T) {
			for i := 0; i < 100; i++ {
				tt.p.genHeight()
			}

			if tt.p.Height < tt.min || tt.p.Height > tt.max {
				st.Fail()
			}
		})
	}
}

func TestPlayer_GetHeightInFeet(t *testing.T) {
	table := []struct {
		h    int
		feet string
	}{
		{183, `6'0"`},
		{147, `4'10"`},
		{164, `5'5"`},
		{207, `6'9"`},
		{198, `6'6"`},
	}

	for _, tt := range table {
		t.Run(tt.feet, func(st *testing.T) {
			p := Player{Height:tt.h}
			r := require.New(st)
			r.Equal(p.GetHeightInFeet(), tt.feet)
		})
	}
}
