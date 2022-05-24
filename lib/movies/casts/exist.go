package casts


func CastExists(name string) bool {
	for _, Cast := range Casts {
		if Cast.Name == name {
			return true
		}
	}
	return false
}