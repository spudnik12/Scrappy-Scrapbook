//Package proverb takes an a set of words and returns a set of proverbs.
package proverb

//Proverb takes an array of string input values and concatenates two adjacent elements
//with a preset value forming an array of string values to return.
func Proverb(rhyme []string) []string {
	var saying string
	//Created a slice which has no elements.
	var text = make([]string, 0, 1)
	//When rhyme has string values concatenation occurs.
	if len(rhyme) != 0 {
		for i := 0; i+1 < len(rhyme); i++ {
			saying = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
			text = append(text, saying)
		}
		saying = "And all for the want of a " + rhyme[0] + "."
		text = append(text, saying)
	}
	return text
}
