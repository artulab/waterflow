package container

import (
	"container/heap"
	"testing"

	"github.com/artulab/waterflow/raster"
)

func TestPriorityQueue(t *testing.T) {
	r := raster.NewRaster(3, 3, -9999)
	r.Data = []float64{
		7, 5, 8,
		1, 6, 9,
		3, 2, 4,
	}

	pq := make(PriorityQueue, 0)

	iter := raster.NewAllIterator(r)

	for iter.Next() {
		cell := iter.Get()
		heap.Push(&pq, cell)
	}

	expected := [9]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}

	actual := [9]float64{
		-1, -1, -1,
		-1, -1, -1,
		-1, -1, -1,
	}

	i := 0
	for pq.Len() > 0 {
		cell := heap.Pop(&pq).(*raster.Cell)
		actual[i] = *cell.Value
		i++
	}

	if expected != actual {
		t.Error("priority queue result isn't expected")
	}
}
