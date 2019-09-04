package fifa

// positions are the positions a player can take in a human readable string.
var positions = map[int]string{
	Goalkeeper:   "GK",
	CentreBack:   "CB",
	LeftBack:     "LB",
	RightBack:    "RB",
	DefensiveMid: "CDM",
	CentreMid:    "CM",
	AttackingMid: "CAM",
	LeftMid:      "LM",
	RightMid:     "RM",
	Striker:      "ST",
}

// GetPositionStr converts the players position ID to a string and returns it. If position ID isn't found (i.e. larger
// than the positions map or negative) an empty string is returned. This is to avoid out of bounds panics if the user
// erroneously assigns the position to something unacceptable.
func (p *Player) GetPositionStr() string {
	if p.Position < 0 || p.Position > len(positions) {
		return ""
	}
	return positions[p.Position]
}
