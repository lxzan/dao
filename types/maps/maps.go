package maps

func Keys[K comparable, V any](m map[K]V) []K {
	var keys = make([]K, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	var values = make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
