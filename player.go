package player

import (
	"database/sql"
)

// Player is contains all of the stats associated with a single player,
// including technical stats, names and meta stats (injury probability etc).
type Player struct {
	ID          int           `json:"id"`
	Value       float64       `json:"value"`
	UserID      sql.NullInt64 `json:"user_id"`
	Position    int           `json:"position_id"`
	Foot        int           `json:"foot_id"`
	Nationality int           `json:"nationality_id"`
	FirstName   string        `json:"first_name"`
	SecondName  string        `json:"second_name"`
	// UI formatted attributes
	PositionStr    string `json:"position"`
	FootStr        string `json:"foot"`
	HeightFeet     string `json:"height_feet"`
	FirstNameShort string `json:"first_name_short"`
	Nation         string `json:"nation"`
	// NationCode is the nationality in code format for CSS
	NationCode string  `json:"nation_code"`
	Height     int     `json:"height_cm"`
	Weight     int     `json:"weight"`
	Overall    float64 `json:"overall"`
	SkillMoves int     `json:"skill_moves"`
	WeakFoot   int     `json:"weak_foot"`
	// Attack
	Crossing  float64 `json:"crossing"`
	Finishing float64 `json:"finishing"`
	Heading   float64 `json:"heading"`
	ShortPass float64 `json:"short_pass"`
	Volleys   float64 `json:"volleys"`
	// Defending
	Marking     float64 `json:"marking"`
	StandTackle float64 `json:"standing_tackle"`
	SlideTackle float64 `json:"slide_tackle"`
	// Skill
	Dribbling   float64 `json:"dribbling"`
	Curve       float64 `json:"curve"`
	FKAcc       float64 `json:"fk_acc"`
	LongPass    float64 `json:"long_pass"`
	BallControl float64 `json:"ball_control"`
	// Power
	ShotPower float64 `json:"shot_power"`
	Jumping   float64 `json:"jumping"`
	Stamina   float64 `json:"stamina"`
	Strength  float64 `json:"strength"`
	LongShots float64 `json:"long_shots"`
	// Movement
	Acceleration float64 `json:"acceleration"`
	SprintSpeed  float64 `json:"sprint_speed"`
	Agility      float64 `json:"agility"`
	Reactions    float64 `json:"reactions"`
	Balance      float64 `json:"balance"`
	// Mentality
	Aggression    float64 `json:"aggresion"`
	Interceptions float64 `json:"interceptions"`
	AttPosition   float64 `json:"att_position"`
	Vision        float64 `json:"vision"`
	Penalties     float64 `json:"penalties"`
	// Goalkeeper
	Diving      float64 `json:"diving"`
	Handling    float64 `json:"handling"`
	Kicking     float64 `json:"kicking"`
	Positioning float64 `json:"positioning"`
	Reflexes    float64 `json:"reflexes"`
	// Meta
	// TalentCo is the coefficient for how talented this player is.
	TalentCo float64 `json:"-"`
	// InjuryCo is the coefficient for how injury prone this player is.
	InjuryCo float64 `json:"-"`
}

// FormatPlayers takes a slice of players and formats each one. Returns the newly formatted slice.
func FormatPlayers(players []Player) []Player {
	for i := range players {
		players[i].Format()
	}
	return players
}

// Format runs every player formatting function on the player.
func (p *Player) Format() {
	p.FootStr = p.GetFootStr()
	p.HeightFeet = p.GetHeightInFeet()
	p.NationCode = p.GetNationalityCode()
	p.Nation = p.GetNationalityStr()
	p.PositionStr = p.GetPositionStr()
	p.FirstNameShort = p.GetShortenedFirstName()
}
