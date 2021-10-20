package fill

import (
	"math"

	"github.com/artulab/waterflow/container"
	"github.com/artulab/waterflow/raster"

	"container/heap"
)

// Fill attempts to correct cells of given inRaster by filling sinks/pits.
// zLimit refers to the maximum value between the original cell value and its
// filled value. Those sinks whose elevation difference is greater than zLimit
// will not be filled. If the zLimit is zero, all sinks will be filled.
func Fill(inRaster *raster.Raster, zLimit float64) (*raster.Raster, error) {
	pq := make(container.PriorityQueue, 0)
	out := raster.CopyRaster(inRaster)
	closed := raster.NewBitmapWithRaster(out)
	edges := raster.NewBorderIterator(out)

	for edges.Next() {
		cell := edges.Get()
		closed.SetWithCell(cell)

		if *cell.Value != inRaster.Nodata {
			heap.Push(&pq, cell)
		}
	}

	for pq.Len() > 0 {
		cell := heap.Pop(&pq).(*raster.Cell)
		neighbors := raster.NewNeighborIteratorWithCell(out, cell)

		for neighbors.Next() {
			ncell := neighbors.Get()

			if ncell == nil {
				continue
			}

			if *ncell.Value == inRaster.Nodata {
				continue
			}

			if closed.GetWithCell(ncell) {
				continue
			}

			*cell.Value = math.Max(*cell.Value, *cell.Value)
			closed.SetWithCell(ncell)
			heap.Push(&pq, ncell)
		}
	}

	return out, nil
}
