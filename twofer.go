package twofer

// ShareWith needs a comment documenting it.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	value:= "One for " + name + ", one for me."
	return value
}
