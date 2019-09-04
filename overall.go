package player

import "math"

// GetOverall calculates the players overall, based on the players stats and their position.
// Overalls are generated from coefficients of various stats which are dependent on player position -
// for example, a strikers overall relies more heavily on finishing than a goalkeeper.
func (p *Player) GetOverall() {
	switch p.Position {
	case Goalkeeper:
		p.Overall = math.Round(
			(0.24 * p.Diving) +
				(0.22 * p.Handling) +
				(0.22 * p.Positioning) +
				(0.22 * p.Reflexes) +
				(0.06 * p.Reactions) +
				(0.04 * p.Kicking))
	case CentreBack:
		p.Overall = math.Round(
			(0.15 * p.Marking) +
				(0.15 * p.StandTackle) +
				(0.15 * p.SlideTackle) +
				(0.1 * p.Heading) +
				(0.1 * p.Strength) +
				(0.08 * p.Aggression) +
				(0.08 * p.Interceptions) +
				(0.05 * p.ShortPass) +
				(0.05 * p.BallControl) +
				(0.05 * p.Reactions) +
				(0.04 * p.Jumping))
	case LeftBack, RightBack:
		p.Overall = math.Round(
			(0.13 * p.SlideTackle) +
				(0.12 * p.StandTackle) +
				(0.12 * p.Interceptions) +
				(0.1 * p.Marking) +
				(0.08 * p.Stamina) +
				(0.08 * p.Reactions) +
				(0.07 * p.Crossing) +
				(0.07 * p.Heading) +
				(0.07 * p.BallControl) +
				(0.06 * p.ShortPass) +
				(0.05 * p.SprintSpeed) +
				(0.05 * p.Aggression))
	case DefensiveMid:
		p.Overall = math.Round(
			(0.13 * p.ShortPass) +
				(0.13 * p.Interceptions) +
				(0.11 * p.LongPass) +
				(0.1 * p.Marking) +
				(0.1 * p.StandTackle) +
				(0.09 * p.BallControl) +
				(0.09 * p.Reactions) +
				(0.08 * p.Vision) +
				(0.06 * p.Stamina) +
				(0.06 * p.Strength) +
				(0.05 * p.Aggression))
	case CentreMid:
		p.Overall = math.Round(
			(0.15 * p.ShortPass) +
				(0.15 * p.LongPass) +
				(0.15 * p.Vision) +
				(0.1 * p.BallControl) +
				(0.1 * p.Dribbling) +
				(0.08 * p.Reactions) +
				(0.08 * p.Interceptions) +
				(0.05 * p.AttPosition) +
				(0.05 * p.StandTackle) +
				(0.05 * p.Stamina) +
				(0.04 * p.LongShots))
	case AttackingMid:
		p.Overall = math.Round(
			(0.16 * p.ShortPass) +
				(0.16 * p.Vision) +
				(0.13 * p.BallControl) +
				(0.12 * p.AttPosition) +
				(0.11 * p.Dribbling) +
				(0.08 * p.Reactions) +
				(0.06 * p.LongShots) +
				(0.05 * p.Finishing) +
				(0.05 * p.ShotPower) +
				(0.04 * p.Acceleration) +
				(0.04 * p.Agility))
	case LeftMid, RightMid:
		p.Overall = math.Round(
			(0.16 * p.Crossing) +
				(0.15 * p.Dribbling) +
				(0.13 * p.BallControl) +
				(0.1 * p.ShortPass) +
				(0.09 * p.AttPosition) +
				(0.06 * p.Acceleration) +
				(0.06 * p.SprintSpeed) +
				(0.06 * p.Reactions) +
				(0.05 * p.Agility) +
				(0.05 * p.Vision) +
				(0.04 * p.Finishing) +
				(0.03 * p.LongShots))
	case Striker:
		p.Overall = math.Round(
			(0.2 * p.Finishing) +
				(0.12 * p.AttPosition) +
				(0.1 * p.Heading) +
				(0.1 * p.ShotPower) +
				(0.1 * p.Reactions) +
				(0.08 * p.Dribbling) +
				(0.08 * p.BallControl) +
				(0.05 * p.Volleys) +
				(0.05 * p.LongShots) +
				(0.05 * p.Acceleration) +
				(0.04 * p.SprintSpeed) +
				(0.03 * p.SprintSpeed))
	}
}
