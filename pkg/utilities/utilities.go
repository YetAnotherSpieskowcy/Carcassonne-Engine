package utilities

// Retrieves key and value from a map, and a bool indicating if the map is empty.
// In case of multi-element maps, it is not specified which element is returned.
func GetAnyElementFromMap[K comparable, V any](m map[K]V) (K, V, bool) {
	for key, value := range m {
		return key, value, true
	}
	return *new(K), *new(V), false
}
