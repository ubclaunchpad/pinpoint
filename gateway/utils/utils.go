package utils

func FirstString(strs []string) string {
	if strs != nil && len(strs) > 0 {
		return strs[0]
	}
	return ""
}
