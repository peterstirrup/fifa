package player

import "math/rand"


const (
	_ = iota
	LeftFoot
	RightFoot
)

var (
	// probFoot is a map of strong foot to how probable it is that a player has.
	// Source: http://www.footballtechnicallab.com/the-myth-of-the-two-footed-footballer-part-1/
	probFoot = map[int]float64{
		LeftFoot:  0.26,
		RightFoot: 0.74,
	}

	// probWeakFootAbility is a map of weak foot ability to a probability that the player has that.
	probWeakFootAbility = map[int]float64{
		1: 0.05,
		2: 0.2,
		3: 0.4,
		4: 0.3,
		5: 0.05,
	}

	// probTwoFoot is the probability that a player is two footed (i.e. both feet are equally strong).
	probTwoFoot = 0.1
)

// genFootTraits generates a players strong foot and weak foot ability.
func (p *Player) genFootTraits() {
	p.genStrongFoot()
	p.genWeakFootAbility()
}

// genStrongFoot assigns a strong foot to the player and a weak foot ability. If the player's position is a full back
// the player's strong foot is correspondent to that side (i.e. left back's strong foot is left).
// If the player is a sided midfielder, the players foot is dependent on probOppositeFoot (the probability the
// player has a strong foot which is the opposite of their side).
//
// For any other position, the strong foot is randomly weighted dependent on probFoot.
func (p *Player) genStrongFoot() {
	switch p.Position {
	case LeftBack:
		p.Foot = LeftFoot
	case RightBack:
		p.Foot = RightFoot
	case LeftMid:
		if rand.Float64() < probOppositeFoot {
			p.Foot = RightFoot
		} else {
			p.Foot = LeftFoot
		}
	case RightMid:
		if rand.Float64() < probOppositeFoot {
			p.Foot = LeftFoot
		} else {
			p.Foot = RightFoot
		}
	default:
		p.Foot = randItemFromSlice(probFoot)
	}
}

// genWeakFootAbility generates a players weak foot ability based on the probWeakFootAbility map.
// Not dependent on position - a possible todo.
func (p *Player) genWeakFootAbility() {
	p.WeakFoot = randItemFromSlice(probWeakFootAbility)
}

// GetFootStr returns the players pref. foot as a string.
func (p *Player) GetFootStr() string {
	switch p.Foot {
	case LeftFoot:
		return "Left"
	case RightFoot:
		return "Right"
	}

	return "NaN"
}
