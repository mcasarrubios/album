package common

// Contains returns if an array contains an element
func Contains(items []string, e string) bool {
	for _, a := range items {
		if a == e {
			return true
		}
	}
	return false
}
