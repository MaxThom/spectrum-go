package led

import "github.com/maxthom/spectrum-go/pkg/utils"

type Position struct {
	position  int
	direction int
}

func NewPosition(start int, end int) *Position {
	dir := utils.RandomInt(0, 2)
	if dir == 0 {
		dir = -1
	}
	return &Position{
		position:  utils.RandomInt(start, end),
		direction: dir,
	}
}
