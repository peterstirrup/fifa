package fifa

import (
	"math"
	"math/rand"
)

var (
	// repeatRand is the number of times we roll the dice in our biased random num gen.
	// Higher the number, the greater the bias to 0.
	repeatRand,
	// MinDiff and maxDiff are the minimum and maximum values we can stray from default stat values.
	MinDiff,
	maxDiff int
	// qualityCo is the quality bias between 0 and 1.
	// The higher the number, the lower the general player quality.
	qualityCo float64
	Players   []Player
	// probOppositeFoot is the chance that a left mid has a strong right foot and vice versa
	probOppositeFoot float64
)

const (
	_ = iota
	Goalkeeper
	CentreBack
	LeftBack
	RightBack
	DefensiveMid
	CentreMid
	AttackingMid
	LeftMid
	RightMid
	Striker
)

// GenPlayer generates a player of a random position, with stats, nationality and name populated.
// Check if name generator has been initialised - if not, initialise it.
func GenPlayer() Player {
	p := Player{
		Position: rand.Intn(Striker) + 1,
	}

	p.genStats()
	p.genNationality()
	p.genName()

	return p
}

// GenPlayerBase generates n players, returning a slice of Players and any errors.
func GenPlayerBase(n int) ([]Player, error) {
	var players []Player

	for i := 0; i < n; i++ {
		players = append(players, GenPlayer())
	}

	return players, nil
}

// genStats generates the players full stat profile. Stats are based on the temp player for their
// corresponding position, which are then deviated randomly. Stats are capped to between 0 and 95 for most stats
// with the exception of acceleration and sprint speed, since players in the base game can realistically reach 99 in these stats.
//
// Goalkeeping stats are generated randomly for outfield players - as they are in the base game (probably).
//
// Some stats are generated based off others - by this, we mean they are linked. Ball control and dribbling are linked, for example.
// So, unless it's a fluke, a player with good ball control will be good at dribbling.
func (p *Player) genStats() {
	*p = playerTemplate[p.Position]
	if rand.Float64() > qualityCo {
		p.TalentCo = capNum(rand.Float64()+rand.Float64(), 0, 1)
	} else {
		p.TalentCo = capNum((rand.Float64())-rand.Float64(), 0, 1)
	}

	// Get random negative coefficient and apply to stats
	randNeg := randNegative()

	// Attack
	p.Crossing = capNum(math.Round(p.Crossing+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Finishing = capNum(math.Round(p.Finishing+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Heading = capNum(math.Round(p.Heading+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.ShortPass = capNum(math.Round(p.ShortPass+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Volleys = capNum(math.Round(p.Volleys+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	// Defending
	p.Marking = capNum(math.Round(p.Marking+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.StandTackle = capNum(math.Round(p.StandTackle+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.SlideTackle = capNum(math.Round(p.SlideTackle+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	// Skill
	p.Dribbling = capNum(math.Round(p.Dribbling+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Curve = capNum(math.Round(p.Curve+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.FKAcc = capNum(math.Round(p.FKAcc+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.LongPass = capNum(math.Round(p.LongPass+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.BallControl = capNum(math.Round(p.Dribbling+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	// Power
	p.ShotPower = capNum(math.Round(p.ShotPower+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Jumping = capNum(math.Round(p.Jumping+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Stamina = capNum(math.Round(p.Stamina+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Strength = capNum(math.Round(p.Strength+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.LongShots = capNum(math.Round(p.LongShots+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	// Movement
	p.Acceleration = capNum(math.Round(p.Acceleration+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(3)+97))
	p.SprintSpeed = capNum(math.Round(p.SprintSpeed+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(3)+97))
	p.Agility = capNum(math.Round(p.Agility+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Reactions = capNum(math.Round(p.Reactions+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Balance = capNum(math.Round(p.Balance+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	// Mentality
	p.Aggression = capNum(math.Round(p.Aggression+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Interceptions = capNum(math.Round(p.Interceptions+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.AttPosition = capNum(math.Round(p.AttPosition+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Vision = capNum(math.Round(p.Vision+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	p.Penalties = capNum(math.Round(p.Penalties+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))

	// Goalkeeper
	// Randomly generate if outfield
	switch p.Position {
	case Goalkeeper:
		p.Diving = capNum(math.Round(p.Diving+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
		p.Handling = capNum(math.Round(p.Handling+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
		p.Kicking = capNum(math.Round(p.Kicking+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
		p.Positioning = capNum(math.Round(p.Positioning+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
		p.Reflexes = capNum(math.Round(p.Reflexes+weightedRand(maxDiff, p.TalentCo)+randNeg), float64(rand.Intn(10)+5), float64(rand.Intn(5)+95))
	default:
		p.Diving = float64(rand.Intn(10) + 5)
		p.Handling = float64(rand.Intn(10) + 5)
		p.Kicking = float64(rand.Intn(10) + 5)
		p.Positioning = float64(rand.Intn(10) + 5)
		p.Reflexes = float64(rand.Intn(10) + 5)
	}

	p.genSkillAbility()
	p.genFootTraits()
	p.genHeight()
	p.genWeight()
	p.GetOverall()
	p.GetValue()

	return
}

// SetRepeatRand sets the repeatRand var to the passed in value.
func SetRepeatRand(rr int) {
	repeatRand = rr
}

// SetMaxDiff sets the maxDiff var to the passed in value.
func SetMaxDiff(md int) {
	maxDiff = md
}

// SetQualityCo sets the qualityCo var to the passed in value.
func SetQualityCo(qco float64) {
	qualityCo = qco
}

// SetProbOppositeFoot sets the probOppositeFoot var to the passed in value.
func SetProbOppositeFoot(p float64) {
	probOppositeFoot = p
}
