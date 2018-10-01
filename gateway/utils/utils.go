package utils

// FirstString returns the first element from an array of strings
func FirstString(strs []string) string {
	if strs != nil && len(strs) > 0 {
		return strs[0]
	}
	return ""
}
