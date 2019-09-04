package fifa

import (
	"fmt"
	"math"
	"math/rand"
)

// freakishlyTallProb is the probability that a player is freakishly tall.
var freakishlyTallProb = 0.2

// genHeight generates a random height for a player dependent on their position. If the player has a heading ability
// over 85, the player can be freakishly tall, meaning the player's height is over 201cm. This is also dependent on
// the freakishly tall probability variable.
func (p *Player) genHeight() {
	switch p.Position {
	case Goalkeeper:
		p.Height = randRange(183, 201)
	case CentreBack:
		p.Height = randRange(173, 195)
	default:
		p.Height = randRange(165, 193)
	}

	// Could they be freakishly tall?
	if p.Heading > 85 {
		if rand.Float64() < freakishlyTallProb {
			p.Height = randRange(201, 211)
		}
	}
}

// GetHeightInFeet returns the players height in feet as a string.
func (p *Player) GetHeightInFeet() string {
	numInches := math.Round(float64(p.Height) / 2.54)
	feet := math.Floor(numInches / 12)
	inches := math.Mod(numInches, 12)
	return fmt.Sprintf(`%v'%v"`, feet, inches)
}

// SetFreakishlyTallProb sets the tall probability variable to p.
func SetFreakishlyTallProb(p float64) {
	freakishlyTallProb = capNum(p, 0, 100)
}
