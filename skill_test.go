package fifa

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_genSkillAbility(t *testing.T) {
	table := []struct {
		position int
		exp []int
	}{
		{Goalkeeper, []int{1, 2}},
		{CentreBack, []int{1, 2, 3, 4}},
		{LeftBack, []int{2, 3, 4, 5}},
		{RightBack, []int{2, 3, 4, 5}},
		{DefensiveMid, []int{2, 3, 4}},
		{CentreMid, []int{2, 3, 4, 5}},
		{AttackingMid, []int{3, 4, 5}},
		{LeftMid, []int{3, 4, 5}},
		{RightMid, []int{3, 4, 5}},
		{Striker, []int{3, 4, 5}},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%v", tt.position), func(st *testing.T) {
			p := Player {
				Position: tt.position,
			}
			p.genSkillAbility()

			r := require.New(st)
			r.Contains(tt.exp, p.SkillMoves)
		})
	}
}
