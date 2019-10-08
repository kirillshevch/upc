package upc

func Valid(s string) bool {
	if len(s) == 12 {
		return true
	} else {
		return false
	}
}
