package fifa

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

const (
	_ = iota
	normalName
	doubleBarrelName
	middleName
	oneName
)

// init reads the first name and second name files for each country (found in name) and populates
// that country's name string slices. After this process, the name generator is ready to start.
func init() {
	for i, v := range Countries {
		first, err := ioutil.ReadFile("./name/" + Countries[i].string + "/first_names.txt")
		if err != nil {
			log.Fatal(fmt.Errorf("name gen: cannot read first names file for country %s with error %s", Countries[i].string, err.Error()))
		}

		second, err := ioutil.ReadFile("./name/" + Countries[i].string + "/second_names.txt")
		if err != nil {
			log.Fatal(fmt.Errorf("name gen: cannot read second names file for country %s with error %s", Countries[i].string, err.Error()))
		}

		v.FirstNames = strings.Split(string(first), "\n")
		v.SecondNames = strings.Split(string(second), "\n")
		Countries[i] = v
	}
}

// genName populates the players first name and second name with a new generated name.
// Probability of name being one name, double barrel or having a middle name is nationality dependent.
// For example, a Brazilian is more likely to have one name than an English player.
func (p *Player) genName() {
	switch randItemFromSlice(map[int]float64{
		normalName:       Countries[p.Nationality].NormalName,
		doubleBarrelName: Countries[p.Nationality].DoubleBarrel,
		middleName:       Countries[p.Nationality].MiddleName,
		oneName:          Countries[p.Nationality].OneName,
	}) {
	case normalName:
		p.FirstName, p.SecondName = Countries[p.Nationality].FirstNames[rand.Intn(len(Countries[p.Nationality].FirstNames))], Countries[p.Nationality].SecondNames[rand.Intn(len(Countries[p.Nationality].SecondNames))]
	case doubleBarrelName:
		p.FirstName, p.SecondName = Countries[p.Nationality].FirstNames[rand.Intn(len(Countries[p.Nationality].FirstNames))], Countries[p.Nationality].SecondNames[rand.Intn(len(Countries[p.Nationality].SecondNames))]+"-"+Countries[p.Nationality].SecondNames[rand.Intn(len(Countries[p.Nationality].SecondNames))]
	case middleName:
		p.FirstName, p.SecondName = Countries[p.Nationality].FirstNames[rand.Intn(len(Countries[p.Nationality].FirstNames))]+" "+Countries[p.Nationality].FirstNames[rand.Intn(len(Countries[p.Nationality].FirstNames))], Countries[p.Nationality].SecondNames[rand.Intn(len(Countries[p.Nationality].SecondNames))]
	case oneName:
		p.FirstName, p.SecondName = "", Countries[p.Nationality].FirstNames[rand.Intn(len(Countries[p.Nationality].FirstNames))]
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
