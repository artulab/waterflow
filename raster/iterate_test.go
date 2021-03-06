package raster

import (
	"testing"
)

func TestAllIterator(t *testing.T) {
	r := NewRaster(2, 1, 1, 1, -9999)
	r.Data = []float64{0, 1}

	expected := []float64{0, 1}

	iter := NewAllIterator(r)

	testIterator(t, iter, expected)
}

func TestNeighborIterator(t *testing.T) {
	r := NewRaster(3, 3, 1, 1, -9999)
	r.Data = []float64{
		10, 11, 12,
		13, 14, 15,
		16, 17, 18,
	}

	iter := NewNeighborIterator(r, 1, 1)

	expected := []float64{15, 18, 17, 16, 13, 10, 11, 12}

	testIterator(t, iter, expected)
}

func TestBorderIterator(t *testing.T) {
	r := NewRaster(4, 3, 1, 1, -9999)
	r.Data = []float64{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
	}

	iter := NewBorderIterator(r)

	expected := []float64{0, 1, 2, 3, 7, 11, 10, 9, 8, 4}

	testIterator(t, iter, expected)
}

func TestInnerRegionIterator(t *testing.T) {
	r := NewRaster(5, 4, 1, 1, -9999)
	r.Data = []float64{
		0, 1, 2, 3, 4,
		5, 6, 7, 8, 9,
		10, 11, 12, 13, 14,
		15, 16, 17, 18, 19,
	}

	iter := NewInnerRegionIterator(r)

	expected := []float64{6, 7, 8, 11, 12, 13}

	testIterator(t, iter, expected)
}

func testIterator(t *testing.T, iter Iterator, expected []float64) {
	for _, v := range expected {
		if iter.Next() != true {
			t.Error("false returned from Next isn't expected")
		}

		if *iter.Get().Value != v {
			t.Error("value returned from Get isn't expected")
		}
	}

	if iter.Next() != false {
		t.Error("true returned from Next isn't expected")
	}

	if iter.Error() != nil {
		t.Error("iterator error isn't expected")
	}

	// try to Get an item from the consumed iterator
	if iter.Get() != nil {
		t.Error("value returned from Get isn't expected")
	}

	if iter.Error() == nil {
		t.Error("iterator error is expected")
	}
}
