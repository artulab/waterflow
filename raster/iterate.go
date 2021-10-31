package raster

import "io"

type Iterator interface {
	Next() bool
	Get() *Cell
	Error() error
}

type AllIterator struct {
	raster *Raster
	idx    int
	err    error
}

func NewAllIterator(r *Raster) *AllIterator {
	i := AllIterator{raster: r, idx: -1}
	return &i
}

func (it *AllIterator) Next() bool {
	if it.idx >= it.raster.Size-1 {
		it.idx = it.raster.Size
		return false
	} else {
		it.idx++
		return true
	}
}

func (it *AllIterator) Get() *Cell {
	if it.idx >= 0 && it.idx < it.raster.Size {
		return &Cell{Value: &it.raster.Data[it.idx],
			Xindex: it.idx % it.raster.Xsize,
			Yindex: it.idx / it.raster.Xsize,
		}
	} else {
		it.err = io.EOF
		return nil
	}
}

func (it *AllIterator) Error() error {
	return it.err
}

type NeighborIterator struct {
	raster  *Raster
	xCenter int
	yCenter int
	idx     int
	err     error
}

// the same order as the Direction starting from Right up to TopRight
var yNeighborIndices = [8]int{0, 1, 1, 1, 0, -1, -1, -1}
var xNeighborIndices = [8]int{1, 1, 0, -1, -1, -1, 0, 1}

func NewNeighborIterator(r *Raster, xCenter int, yCenter int) *NeighborIterator {
	i := NeighborIterator{raster: r, idx: -1, xCenter: xCenter,
		yCenter: yCenter}
	return &i
}

func NewNeighborIteratorWithCell(r *Raster, c *Cell) *NeighborIterator {
	i := NewNeighborIterator(r, c.Xindex, c.Yindex)
	return i
}

func (it *NeighborIterator) Next() bool {
	if it.idx >= 7 {
		it.idx = 8
		return false
	} else {
		it.idx++
		return true
	}
}

func (it *NeighborIterator) Get() *Cell {
	if it.idx >= 0 && it.idx < 8 {
		x := it.xCenter + xNeighborIndices[it.idx]
		y := it.yCenter + yNeighborIndices[it.idx]

		if it.raster.IsInRegion(x, y) {
			return &Cell{Value: &it.raster.Data[y*it.raster.Xsize+x],
				Xindex: x, Yindex: y}
		} else {
			return nil
		}
	} else {
		it.err = io.EOF
		return nil
	}
}

func (it *NeighborIterator) Error() error {
	return it.err
}

type BorderIterator struct {
	raster *Raster
	x      int
	y      int
	idx    int
	size   int
	err    error
}

func NewBorderIterator(r *Raster) *BorderIterator {
	i := BorderIterator{raster: r, idx: -1, x: -1, y: -1,
		size: 2*(r.Xsize+r.Ysize) - 4}
	return &i
}

func (it *BorderIterator) Next() bool {
	if it.idx >= it.size-1 {
		it.idx = it.size
		return false
	} else {
		it.idx++

		if it.idx < it.raster.Xsize {
			it.x++
			it.y = 0
		} else if it.idx >= it.raster.Xsize-1 &&
			it.idx < it.raster.Xsize+it.raster.Ysize-1 {
			it.x = it.raster.Xsize - 1
			it.y++
		} else if it.idx >= it.raster.Xsize+it.raster.Ysize-1 &&
			it.idx < 2*it.raster.Xsize+it.raster.Ysize-2 {
			it.y = it.raster.Ysize - 1
			it.x--
		} else {
			it.x = 0
			it.y--
		}
		return true
	}
}

func (it *BorderIterator) Get() *Cell {
	if it.idx >= 0 && it.idx < it.size {
		return &Cell{Value: &it.raster.Data[it.y*it.raster.Xsize+it.x],
			Xindex: it.x, Yindex: it.y}
	} else {
		it.err = io.EOF
		return nil
	}
}

func (it *BorderIterator) Error() error {
	return it.err
}

type InnerRegionIterator struct {
	raster *Raster
	x      int
	y      int
	idx    int
	size   int
	err    error
}

func NewInnerRegionIterator(r *Raster) *InnerRegionIterator {
	i := InnerRegionIterator{raster: r, x: 0, y: 1, idx: -1,
		size: (r.Xsize * r.Ysize) - 2*(r.Xsize+r.Ysize) + 4}
	return &i
}

func (it *InnerRegionIterator) Next() bool {
	if it.idx >= it.size-1 {
		it.idx = it.size
		return false
	} else {
		it.idx++

		if it.x < it.raster.Xsize-2 {
			it.x++
		} else {
			it.x = 1
			it.y++
		}

		return true
	}
}

func (it *InnerRegionIterator) Get() *Cell {
	if it.idx >= 0 && it.idx < it.size {
		return &Cell{Value: &it.raster.Data[it.y*it.raster.Xsize+it.x],
			Xindex: it.x, Yindex: it.y}
	} else {
		it.err = io.EOF
		return nil
	}
}

func (it *InnerRegionIterator) Error() error {
	return it.err
}
