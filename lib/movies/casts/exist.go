package casts

import "errors"


func CastExists(name string) bool {
	for _, Cast := range Casts {
		if Cast.Name == name {
			return true
		}
	}
	return false
}


func (cast *Cast) Index() (int, error) {
	for index, Cast := range Casts {
		if Cast.ID == cast.ID {
			return index, nil
		}
	}
	return 0, errors.New("cast does not exist")
}