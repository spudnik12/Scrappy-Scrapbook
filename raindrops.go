//Package raindrops converts numbers into rain-talk.
package raindrops

import "strconv"

var soundsOfRain = map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}

//Convert takes an int, checks if it's divisible by any of the keys of rain, and returns associated string values.
func Convert(input int) string {
	var expected string
	for i := 3; i <= 7; i = i + 2 {
		if input%i == 0 {
			expected = expected + soundsOfRain[i]
		}
	}
	//Empty expected value converts the input to a string.
	if expected == "" {
		expected = strconv.Itoa(input)
	}
	return expected
}
