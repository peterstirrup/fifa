package player

import "math/rand"

const (
	_ = iota
	Argentina
	Belgium
	Brazil
	England
	France
	Germany
	Italy
	Netherlands
	Portugal
	Spain
)

type Country struct {
	// string is the countries name ("argentina")
	string
	// FirstNames is that countries first names
	FirstNames []string
	// SecondNames is that countries second names
	SecondNames []string
	// NormalName is the probability of a normal name from that country (first and second name).
	NormalName float64
	// DoubleBarrel is the probability of a name from that country
	// being double barrelled (0 <= n < 1)
	DoubleBarrel float64
	// MiddleName is the probability of a name from that country
	// having a middle name (0 <= n < 1)
	MiddleName float64
	// OneName is the probability of a name from that country
	// having only one word (like Ronaldinho) (0 <= n < 1)
	OneName float64
}

var Countries = map[int]Country{
	Argentina: {
		string:       "Argentina",
		NormalName:   0.65,
		DoubleBarrel: 0,
		MiddleName:   0.3,
		OneName:      0.05,
	},
	Belgium: {
		string:       "Belgium",
		NormalName:   0.85,
		DoubleBarrel: 0.05,
		MiddleName:   0.1,
		OneName:      0,
	},
	Brazil: {
		string:       "Brazil",
		NormalName:   0.6,
		DoubleBarrel: 0,
		MiddleName:   0.2,
		OneName:      0.2,
	},
	England: {
		string:       "England",
		NormalName:   0.9,
		DoubleBarrel: 0.05,
		MiddleName:   0.05,
		OneName:      0,
	},
	France: {
		string:       "France",
		NormalName:   0.9,
		DoubleBarrel: 0.05,
		MiddleName:   0.05,
		OneName:      0,
	},
	Germany: {
		string:       "Germany",
		NormalName:   0.7,
		DoubleBarrel: 0.1,
		MiddleName:   0.2,
		OneName:      0,
	},
	Italy: {
		string:       "Italy",
		NormalName:   0.55,
		DoubleBarrel: 0.2,
		MiddleName:   0.2,
		OneName:      0.05,
	},
	Netherlands: {
		string:       "Netherlands",
		NormalName:   0.45,
		DoubleBarrel: 0.2,
		MiddleName:   0.3,
		OneName:      0.05,
	},
	Portugal: {
		string:       "Portugal",
		NormalName:   0.75,
		DoubleBarrel: 0.05,
		MiddleName:   0.1,
		OneName:      0.1,
	},
	Spain: {
		string:       "Spain",
		NormalName:   0.65,
		DoubleBarrel: 0.1,
		MiddleName:   0.2,
		OneName:      0.05,
	},
}

// genNationality generates the players random nationality.
func (p *Player) genNationality() {
	p.Nationality = rand.Intn(Spain) + 1
}

// GetNationalityCode returns the users nationality as a nationality code (for CSS and frontend class formatting).
// e.g. Argentina --> "AR". If not a valid country ID, return "NaN".
// Note: England --> "GB" is correct (and not "EN") because the CSS package doesn't have an
// England flag (go figure).
func (p *Player) GetNationalityCode() string {
	switch p.Nationality {
	case Argentina:
		return "AR"
	case Belgium:
		return "BE"
	case Brazil:
		return "BR"
	case England:
		return "GB"
	case Italy:
		return "IT"
	case France:
		return "FR"
	case Germany:
		return "DE"
	case Netherlands:
		return "NL"
	case Portugal:
		return "PT"
	case Spain:
		return "ES"
	}

	return "NaN"
}

// GetNationalityStrings returns the users nationality as a string. Nationality is generated as an ID - this function
// provides the nationality as a human readable string.
// e.g. 0 --> "Argentina"
func (p *Player) GetNationalityStr() string {
	switch p.Nationality {
	case Argentina:
		return "Argentina"
	case Belgium:
		return "Belgium"
	case Brazil:
		return "Brazil"
	case England:
		return "England"
	case Italy:
		return "Italy"
	case France:
		return "France"
	case Germany:
		return "Germany"
	case Netherlands:
		return "Netherlands"
	case Portugal:
		return "Portugal"
	case Spain:
		return "Spain"
	}

	return "NaN"
}
