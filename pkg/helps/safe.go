package helps

import "github.com/tiyee/gokit/pkg/constraints"

func SafePage[T constraints.Integer](i T) T {
	var j T = 1
	if i < j {
		i = j
	}
	return i
}
