package fifa

import (
	"math/rand"
	"time"
)

// randItemFromSlice picks a random int (index) from a map m using the probability assigned to that value.
// e.g. m = {1: 0.1, 2: 0.7, 3: 0.2}
// Probability of picking 2 = 0.7.
func randItemFromSlice(m map[int]float64) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()

	var tot float64
	for i, p := range m {
		tot += p

		if r <= tot {
			return i
		}
	}

	return 0
}

// randRange generates a random int between the min and max passed in.
func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

// capNum caps the value of an int to make sure it's between min and max.
func capNum(n, min, max float64) float64 {
	if n > max {
		return max
	} else if n < min {
		return min
	}
	return n
}

// weightedRand returns a random n (-1 <= n <= 1) (depending on talent coefficient) with bias towards
// numbers closer to 0 (a sudo normal distribution).
func weightedRand(max int, talentCo float64) float64 {
	num := posWeightedRand(max)

	if rand.Float64() > talentCo {
		return num - (2 * num)
	}

	return num
}

// posWeightedRand returns a random number bias towards 0 - always positive. 0 < n <= max.
func posWeightedRand(max int) float64 {
	var num float64
	for i := 0; i < repeatRand; i++ {
		num += rand.Float64() * float64(max)/float64(repeatRand)
	}

	return num
}

// randNegative generates a random negative number.
// TODO - Doc this......
func randNegative() float64 {
	r := float64(rand.Intn(3)) - rand.Float64()
	if r < 0 {
		return r * 10
	}
	return float64(r - (r*2)*10)
}