package git

// match 0 value
func remove0(it []string, param []string) bool {
	if len(it) >= 1 && it[0] == param[0] {
		return true
	}
	return false
}
