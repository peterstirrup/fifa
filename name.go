package fifa

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

const (
	normalName = iota + 1
	doubleBarrelName
	middleName
	oneName
)

// init reads the first name and second name files for each country (found in name) and populates
// that country's name string slices. After this process, the name generator is ready to start.
func init() {
	for i, v := range countries {
		first, err := ioutil.ReadFile("name/" + countries[i].string + "/first_names.txt")
		if err != nil {
			log.Fatal(fmt.Errorf("name gen: cannot read first names file for country %s with error: %s", countries[i].string, err.Error()))
		}

		second, err := ioutil.ReadFile("name/" + countries[i].string + "/second_names.txt")
		if err != nil {
			log.Fatal(fmt.Errorf("name gen: cannot read second names file for country %s with error: %s", countries[i].string, err.Error()))
		}

		v.firstNames = strings.Split(string(first), "\n")
		v.secondNames = strings.Split(string(second), "\n")
		countries[i] = v
	}
}

// genName populates the players first name and second name with a new generated name.
// Probability of name being one name, double barrel or having a middle name is nationality dependent.
// For example, a Brazilian is more likely to have one name than an English player.
func (p *Player) genName() {
	switch randItemFromSlice(map[int]float64{
		normalName:       countries[p.Nationality].normalName,
		doubleBarrelName: countries[p.Nationality].doubleBarrel,
		middleName:       countries[p.Nationality].middleName,
		oneName:          countries[p.Nationality].oneName,
	}) {
	case normalName:
		p.FirstName, p.SecondName = countries[p.Nationality].firstNames[rand.Intn(len(countries[p.Nationality].firstNames))], countries[p.Nationality].secondNames[rand.Intn(len(countries[p.Nationality].secondNames))]
	case doubleBarrelName:
		p.FirstName, p.SecondName = countries[p.Nationality].firstNames[rand.Intn(len(countries[p.Nationality].firstNames))], countries[p.Nationality].secondNames[rand.Intn(len(countries[p.Nationality].secondNames))]+"-"+countries[p.Nationality].secondNames[rand.Intn(len(countries[p.Nationality].secondNames))]
	case middleName:
		p.FirstName, p.SecondName = countries[p.Nationality].firstNames[rand.Intn(len(countries[p.Nationality].firstNames))]+" "+countries[p.Nationality].firstNames[rand.Intn(len(countries[p.Nationality].firstNames))], countries[p.Nationality].secondNames[rand.Intn(len(countries[p.Nationality].secondNames))]
	case oneName:
		p.FirstName, p.SecondName = "", countries[p.Nationality].firstNames[rand.Intn(len(countries[p.Nationality].firstNames))]
	}
}

// GetShortenedFirstName converts the first name to just the first letter,  unless player name is one name
// (in which case it is left blank). This is returned as a string.
// e.g. John Doe --> "J"
//      Doelinho --> ""
func (p *Player) GetShortenedFirstName() string {
	if len(p.FirstName) > 0 {
		return string(p.FirstName[0])
	}

	return ""
}
