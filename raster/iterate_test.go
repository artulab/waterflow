package raster

import (
	"testing"
)

func TestAllIterator(t *testing.T) {
	r := New(2, 1, -9999)
	r.Data = []float32{0, 1}

	expected := []float32{0, 1}

	iter := NewAllIterator(r)

	testIterator(t, iter, expected)
}

func TestNeighborIterator(t *testing.T) {
	r := New(3, 3, -9999)
	r.Data = []float32{
		10, 11, 12,
		13, 14, 15,
		16, 17, 18,
	}

	iter := NewNeighborIterator(r, 1, 1)

	expected := []float32{10, 11, 12, 13, 15, 16, 17, 18}

	testIterator(t, iter, expected)
}

func TestBorderIterator(t *testing.T) {
	r := New(4, 3, -9999)
	r.Data = []float32{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
	}

	iter := NewBorderIterator(r)

	expected := []float32{0, 1, 2, 3, 7, 11, 10, 9, 8, 4}

	testIterator(t, iter, expected)
}

func testIterator(t *testing.T, iter Iterator, expected []float32) {
	for _, v := range expected {
		if iter.Next() != true {
			t.Error("False returned from Next isn't expected")
		}

		if *iter.Get().Value != v {
			t.Error("Value returned from Get isn't expected")
		}
	}

	if iter.Next() != false {
		t.Error("True returned from Next isn't expected")
	}

	if iter.Error() != nil {
		t.Error("Iterator error isn't expected")
	}

	// try to Get an item from the consumed iterator
	if iter.Get() != nil {
		t.Error("Value returned from Get isn't expected")
	}

	if iter.Error() == nil {
		t.Error("Iterator error is expected")
	}
}
