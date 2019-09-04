package fifa

import (
	"fmt"
	"strconv"
)

// valueMod is the modifier for a players value based on position, because, for example,
// in the real transfer market goalkeepers are not as expensive as strikers.
var valueMod = map[int]float64{
	Goalkeeper:   0.4,
	LeftBack:     0.4,
	CentreBack:   0.5,
	RightBack:    0.4,
	DefensiveMid: 0.5,
	CentreMid:    0.5,
	AttackingMid: 0.7,
	LeftMid:      0.8,
	RightMid:     0.8,
	Striker:      1,
}

// GetValue calculates the players value based on player overall and position, and truncates the number
// to 2 dp. Sets the Value attribute. This is a deterministic process, with no chaos or randomness.
func (p *Player) GetValue() {
	if p.Overall < 70 {
		p.Value = floatToTwoDecimalPlaces(valueMod[p.Position] * p.Overall * 0.02)
	} else if p.Overall >= 70 && p.Overall < 75 {
		p.Value = floatToTwoDecimalPlaces(valueMod[p.Position] * p.Overall * 0.08)
	} else if p.Overall >= 75 && p.Overall < 80 {
		p.Value = floatToTwoDecimalPlaces(valueMod[p.Position] * p.Overall * 0.15)
	} else if p.Overall >= 80 && p.Overall < 85 {
		p.Value = floatToTwoDecimalPlaces(valueMod[p.Position] * p.Overall * 0.3022)
	} else if p.Overall >= 85 && p.Overall < 90 {
		p.Value = floatToTwoDecimalPlaces(valueMod[p.Position] * p.Overall * 0.6)
	} else if p.Overall >= 90 {
		p.Value = floatToTwoDecimalPlaces(valueMod[p.Position] * p.Overall * 0.8)
	}
}

// floatToTwoDecimalPlaces converts float64 to two decimal places.
func floatToTwoDecimalPlaces(n float64) float64 {
	i, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", n), 64)
	return i
}
