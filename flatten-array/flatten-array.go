package flatten

func Flatten(nested any) []any {
	result := []any{} // always return non-nil slice

	switch v := nested.(type) {
	case []any:
		for _, item := range v {
			flattened := Flatten(item)
			for _, f := range flattened {
				if f != nil { // skip nil values
					result = append(result, f)
				}
			}
		}
	case []int:
		for _, item := range v {
			result = append(result, item)
		}
	case []string:
		for _, item := range v {
			result = append(result, item)
		}
	default:
		if v != nil { // skip nil single values
			result = append(result, v)
		}
	}

	return result
}
