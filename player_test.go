package fifa

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_Format(t *testing.T) {
	table := []struct {
		p         Player
		expPlayer Player
	}{
		{Player{FirstName: "John", Foot: RightFoot, Position: Striker, Height: 190, Nationality: France},
			Player{FootStr: "Right", PositionStr: "ST", HeightFeet: `6'3"`, FirstNameShort: "J", Nation: "France", NationCode: "FR"}},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%+v", tt.p), func(st *testing.T) {
			tt.p.Format()
			r := require.New(st)
			r.Equal(tt.expPlayer.FootStr, tt.p.FootStr)
			r.Equal(tt.expPlayer.PositionStr, tt.p.PositionStr)
			r.Equal(tt.expPlayer.HeightFeet, tt.p.HeightFeet)
			r.Equal(tt.expPlayer.FirstNameShort, tt.p.FirstNameShort)
			r.Equal(tt.expPlayer.Nation, tt.p.Nation)
			r.Equal(tt.expPlayer.NationCode, tt.p.NationCode)
		})
	}
}
