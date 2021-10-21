package container

import (
	"errors"

	"github.com/artulab/waterflow/raster"
)

type Stack struct {
	data []*raster.Cell
}

func NewStack() *Stack {
	return &Stack{make([]*raster.Cell, 0)}
}

func (s *Stack) Push(c *raster.Cell) {
	s.data = append(s.data, c)
}

func (s *Stack) Pop() (*raster.Cell, error) {
	l := len(s.data)
	if l == 0 {
		return nil, errors.New("empty stack")
	}

	res := s.data[l-1]
	s.data = s.data[:l-1]
	//s.data[l-1] = nil // avoid memory leak
	return res, nil
}
