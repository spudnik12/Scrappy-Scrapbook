//package leap evaluates a given year to test whether it is a leap year
package leap

//func IsLeapYear takes an int type value, year, and returns a bool
func IsLeapYear(year int) bool {
	//tests conditional (group) to determine if year is a leap
	if (year%4 == 0 && year%100 != 0) || (year%100 == 0 && year%400 == 0) {
		return true
	}
	//otherwise, it is not a leap year
	return false
}
