package fifa

import "math/rand"

// Country IDs
const (
	Argentina = iota + 1
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

// country holds the names associated with a particular country, the string of the countries name (as an embedded type)
// and probabilities of types of name (i.e. double barrel, middle names etc).
type country struct {
	// string is the countries name ("Argentina")
	string
	// firstNames is that countries first names
	firstNames []string
	// secondNames is that countries second names
	secondNames []string
	// normalName is the probability of a normal name from that country (first and second name).
	normalName float64
	// doubleBarrel is the probability of a name from that country
	// being double barrelled (0 <= n < 1)
	doubleBarrel float64
	// middleName is the probability of a name from that country
	// having a middle name (0 <= n < 1)
	middleName float64
	// oneName is the probability of a name from that country
	// having only one word (like Ronaldinho) (0 <= n < 1)
	oneName float64
}

var countries = map[int]country{
	Argentina: {
		string:       "Argentina",
		normalName:   0.65,
		doubleBarrel: 0,
		middleName:   0.3,
		oneName:      0.05,
	},
	Belgium: {
		string:       "Belgium",
		normalName:   0.85,
		doubleBarrel: 0.05,
		middleName:   0.1,
		oneName:      0,
	},
	Brazil: {
		string:       "Brazil",
		normalName:   0.6,
		doubleBarrel: 0,
		middleName:   0.2,
		oneName:      0.2,
	},
	England: {
		string:       "England",
		normalName:   0.9,
		doubleBarrel: 0.05,
		middleName:   0.05,
		oneName:      0,
	},
	France: {
		string:       "France",
		normalName:   0.9,
		doubleBarrel: 0.05,
		middleName:   0.05,
		oneName:      0,
	},
	Germany: {
		string:       "Germany",
		normalName:   0.7,
		doubleBarrel: 0.1,
		middleName:   0.2,
		oneName:      0,
	},
	Italy: {
		string:       "Italy",
		normalName:   0.55,
		doubleBarrel: 0.2,
		middleName:   0.2,
		oneName:      0.05,
	},
	Netherlands: {
		string:       "Netherlands",
		normalName:   0.45,
		doubleBarrel: 0.2,
		middleName:   0.3,
		oneName:      0.05,
	},
	Portugal: {
		string:       "Portugal",
		normalName:   0.75,
		doubleBarrel: 0.05,
		middleName:   0.1,
		oneName:      0.1,
	},
	Spain: {
		string:       "Spain",
		normalName:   0.65,
		doubleBarrel: 0.1,
		middleName:   0.2,
		oneName:      0.05,
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

// GetNationalityStr returns the users nationality as a string. Nationality is generated as an ID - this function
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
