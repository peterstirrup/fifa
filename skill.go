package player

var (
	// probSkillAbility is a map of positions to a map of skill ability (stars) to the probability that a player
	// in that position has that ability.
	probSkillAbility = map[int]map[int]float64{
		Goalkeeper: {
			1: 0.95,
			2: 0.05,
			3: 0,
			4: 0,
			5: 0,
		},
		CentreBack: {
			1: 0.05,
			2: 0.8,
			3: 0.1,
			4: 0.05,
			5: 0,
		},
		LeftBack:  probFullBackSkillAbility,
		RightBack: probFullBackSkillAbility,
		DefensiveMid: {
			1: 0,
			2: 0.3,
			3: 0.5,
			4: 0.2,
			5: 0,
		},
		CentreMid: {
			1: 0,
			2: 0.1,
			3: 0.6,
			4: 0.2,
			5: 0.1,
		},
		AttackingMid: {
			1: 0,
			2: 0,
			3: 0.4,
			4: 0.4,
			5: 0.2,
		},
		LeftMid: probAttackerSkillAbility,
		RightMid: probAttackerSkillAbility,
		Striker: probAttackerSkillAbility,
	}

	// probFullBackSkillAbility is the skill ability probability of a fullback.
	probFullBackSkillAbility = map[int]float64{
		1: 0,
		2: 0.3,
		3: 0.5,
		4: 0.15,
		5: 0.05,
	}

	// probFullBackSkillAbility is the skill ability probability of an attacker (winger/striker)
	probAttackerSkillAbility = map[int]float64{
		1: 0,
		2: 0,
		3: 0.2,
		4: 0.6,
		5: 0.2,
	}
)

// genSkillAbility genWeakFootAbility generates a players skill ability based on the probSkillAbility map.
// Dependent on position (i.e. an attacker is more likely to have 4 star skills than a defender).
func (p *Player) genSkillAbility() {
	p.SkillMoves = randItemFromSlice(probSkillAbility[p.Position])
}
