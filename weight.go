package player

import "math/rand"

// genWeight generates a random weight between 150 and 220 for a player.
func (p *Player) genWeight() {
	p.Weight = rand.Intn(70) + 150
}