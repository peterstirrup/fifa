package player

import (
	"testing"
)

func TestPlayer_genWeight(t *testing.T) {
	t.Run("weight between 150 and 220", func(st *testing.T) {
		for i := 0; i < 100; i++ {
			var p Player
			p.genWeight()
			if p.Weight < 150 || p.Weight > 220 {
				st.Fail()
			}
		}
	})
}
