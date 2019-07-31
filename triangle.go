//Package triangle determines the symmetry of a triangle based on the lengths of the sides.
package triangle

import "math"

//Kind assignment carry triangle symmetry identification values.
type Kind string

//Constants contain triangle symmetry values and the non-triangle value.
const NaT, Equ, Iso, Sca Kind = "NaT", "Equ", "Iso", "Sca"

//InputTest tests for any inadmissible values.
func InputTest(s [3]float64) bool {
	for _, i := range s {
		//checking for non-number, -/+ infinite, zero, and negative values
		if i <= 0 || math.IsNaN(i) || i == math.Inf(1) || i == math.Inf(-1) {
			return true
		}
	}
	return false
}

//SideTest tests for triangle inequality.
func SideTest(s [3]float64) bool {

	for i := 0; i < 3; i++ {
		if s[0]+s[1] < s[2] {
			return true
		}
		variableSwap(&s)
	}
	return false
}

//Equilateral checks side lengths for equilateral triangle symmetry.
func Equilateral(s [3]float64) bool {

	if s[0] == s[1] && s[1] == s[2] {
		return true
	}
	return false
}

//Isosceles checks side lengths for isosceles triangle symmetry.
func Isosceles(s [3]float64) bool {

	for i := 0; i < 3; i++ {
		if s[0] == s[1] && s[1] != s[2] {
			return true
		}
		variableSwap(&s)
	}
	return false
}

//Scalene checks side lengths for scalene triangle symmetry type (non-symmetrical triangle).
func Scalene(s [3]float64) bool {
	if s[0] != s[1] && s[1] != s[2] && s[2] != s[0] {
		return true
	}
	return false
}

//variableSwap moves the values of the old-3rd element of the array
//becomes the 1st, old-1st+2nd become 2nd and 3rd element.
func variableSwap(s *[3]float64) {
	var extra float64
	extra = s[2]
	s[2] = s[1]
	s[1] = s[0]
	s[0] = extra
}

//KindFromSides determines triangle symmetry type from the given input
func KindFromSides(a, b, c float64) Kind {

	var k Kind
	s := [3]float64{a, b, c}

	switch { //test to determine triangle symmetry of usable inputs
	case InputTest(s): //test to see if input is a usable number
		k = NaT
	case SideTest(s): //test to see if sides form a triangle
		k = NaT
	case Equilateral(s): //check to see if all three sides are equal
		k = Equ
	case Isosceles(s): //check to see if any two sides are equal
		k = Iso
	case Scalene(s): //check to see if there no symmetry
		k = Sca
	}
	return k //return symmetry of triangle or Not a triangle
}
