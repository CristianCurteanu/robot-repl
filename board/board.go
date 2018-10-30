package board

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Width  int64
	Height int64
}

type BoardSizeDefinitionError struct {
	arg  string
	prob string
}

func (e *BoardSizeDefinitionError) Error() string {
	return fmt.Sprintf("%s: %s", e.prob, e.arg)
}

func NewBoard(size string) (Board, error) {
	splitted := strings.Split(size, "x")
	x, xerr := strconv.ParseInt(splitted[0], 10, 64)
	y, yerr := strconv.ParseInt(splitted[1], 10, 64)

	if xerr != nil || yerr != nil {
		return Board{}, &BoardSizeDefinitionError{size, "Size is not properly defined"}
	}
	return Board{Width: x, Height: y}, nil
}
