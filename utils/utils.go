package utils

func Contains(s []string, e string) bool { // Check if a string is in a slice of strings
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
