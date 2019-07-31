//Package hamming compares the lengths of DNA strands
//and counts the number of point mutations.
package hamming

import "errors"

//Distance compares two DNA string values,
//and returns both an error and an int types value that
//represents pass/fail length match and point mutation count.
func Distance(a, b string) (int, error) {
	//Testing for a match in the lengths of string values.
	if len(a) != len(b) {
		return 0, errors.New("mismatching DNA strand lengths")
	}
	//Calls on the function CheckPairs to compare point mutations and returns a value.
	//Calling function Checkpairs to initiate recursion.
	return CheckPairs(a, b), nil

}

//CheckPairs compares point mutations for matching lengths of DNA,
//and returns the number of occurances.
func CheckPairs(a, b string) int {
	if len(a) > 0 {
		//Testing for point mutations at the first element of the string,
		//and recursion call the rest of the elements.
		if a[0] != b[0] {
			return 1 + CheckPairs(a[1:len(a)], b[1:len(b)])
		}

		if a[0] == b[0] {
			return 0 + CheckPairs(a[1:len(a)], b[1:len(b)])
		}
	}
	//This is the recursion terminator.
	return 0
}
