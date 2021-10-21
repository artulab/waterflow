package fill

import (
	"errors"
	"math"

	"github.com/artulab/waterflow/container"
	"github.com/artulab/waterflow/raster"

	"container/heap"
)

// Fill attempts to correct cells of gi`ven inRaster by filling sinks/pits.
// zLimit refers to the maximum elevation different between the sink and its
// pour point. Those sinks whose elevation difference is greater than zLimit
// will not be filled. If the zLimit is zero, all sinks will be filled.
func Fill(inRaster *raster.Raster, zLimit float64) (*raster.Raster, error) {
	if zLimit < 0 {
		return nil, errors.New("zLimit is expected to be non-negative")
	}

	pq := make(container.PriorityQueue, 0)
	out := raster.CopyRaster(inRaster)
	closed := raster.NewBitmapWithRaster(out)
	edges := raster.NewBorderIterator(out)

	var sinks *raster.Bitmap = nil
	if zLimit > 0 {
		sinks = raster.NewBitmapWithRaster(out)
	}

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

			val := math.Max(*cell.Value, *ncell.Value)
			*ncell.Value = val

			if zLimit > 0 && val != inRaster.GetWithCell(ncell) {
				sinks.SetWithCell(ncell)
			}

			closed.SetWithCell(ncell)
			heap.Push(&pq, ncell)
		}
	}

	if zLimit > 0 {
		// find connected cells in the sink
		label := 0
		maxelev := make([]float64, 0)
		heads := make([]*raster.Cell, 0)
		closed = raster.NewBitmapWithRaster(out)

		it := raster.NewAllIterator(out)
		for it.Next() {
			cell := it.Get()

			// if cell is in the sink and not visited,
			// found a new sink component
			if sinks.GetWithCell(cell) && !closed.GetWithCell(cell) {
				label++

				stack := container.NewStack()
				stack.Push(cell)
				heads = append(heads, cell)

				for {
					scell, _ := stack.Pop()

					if scell == nil {
						break
					}

					closed.SetWithCell(scell)

					// store the max elevation difference for the sink
					if len(maxelev) < label {
						maxelev = append(maxelev,
							*scell.Value-inRaster.GetWithCell(scell))
					} else {
						maxelev[label-1] = math.Max(maxelev[label-1],
							*scell.Value-inRaster.GetWithCell(scell))
					}

					neighbors := raster.NewNeighborIteratorWithCell(out, scell)

					for neighbors.Next() {
						ncell := neighbors.Get()

						if closed.GetWithCell(ncell) {
							continue
						}

						if !sinks.GetWithCell(ncell) {
							continue
						}

						stack.Push(ncell)
					}
				}
			}
		}

		for i, head := range heads {
			maxdiff := maxelev[i]

			// undo elevations for the sinks whose maximum elevation difference
			// is greater zLimit
			if maxdiff >= zLimit {
				stack := container.NewStack()
				stack.Push(head)

				closed = raster.NewBitmapWithRaster(out)

				for {
					scell, _ := stack.Pop()

					if scell == nil {
						break
					}

					if closed.GetWithCell(scell) {
						continue
					}

					*scell.Value = inRaster.GetWithCell(scell)
					closed.SetWithCell(scell)

					neighbors := raster.NewNeighborIteratorWithCell(out, scell)

					for neighbors.Next() {
						ncell := neighbors.Get()

						if !sinks.GetWithCell(ncell) {
							continue
						}

						stack.Push(ncell)
					}
				}
			}
		}
	}

	return out, nil
}
