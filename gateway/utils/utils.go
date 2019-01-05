package utils

// ToMap parses given key-value pairs into a map
func ToMap(args ...interface{}) (fields map[string]interface{}) {
	l := len(args)
	if args != nil && l > 0 && l%2 == 0 {
		fields = make(map[string]interface{})
		for i := 0; i < l; i += 2 {
			if key, ok := args[i].(string); !ok {
				continue
			} else {
				fields[key] = args[i+1]
			}
		}
		if len(fields) == 0 {
			fields = nil
		}
	}
	return
}
