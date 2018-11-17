package utils

// FirstString returns the first element from an array of strings
func FirstString(strs []interface{}) string {
	if strs != nil && len(strs) > 0 {
		return strs[0].(string)
	}
	return ""
}

// ToMap parses given key-value pairs into a map
func ToMap(args ...interface{}) map[string]interface{} {
	var fields map[string]interface{}
	l := len(args)
	if args != nil && l > 0 && l%2 == 0 {
		fields = make(map[string]interface{})
		for i := 0; i < l; i += 2 {
			fields[args[i].(string)] = args[i+1]
		}
	}
	return fields
}
