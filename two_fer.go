//Package twofer contains function ShareWith which returns a string value based on the input of the function call.
package twofer

//ShareWith returns 3 concatenated string variables.
func ShareWith(name string) string {
	start := "One for "
	end := ", one for me."
	//Name is initialized with a string value when empty.
	if name == "" {
		name = "you"
	}

	return (start + name + end)
}
