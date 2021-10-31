package flow

import (
	"math"

	"github.com/artulab/waterflow/raster"
)

type FlowDirectionType int

const (
	D8 FlowDirectionType = iota
	DInf
)

func FlowDirection(inRaster *raster.Raster, forceFlow bool,
	computeDrop bool) (*raster.Raster, *raster.Raster, error) {

	out := raster.CopyRaster(inRaster)
	closed := raster.NewBitmapWithRaster(out)
	direction := raster.NewIntmapWithRaster(inRaster)

	// fill one-cell sinks
	innerRegionIt := raster.NewInnerRegionIterator(inRaster)

	for innerRegionIt.Next() {
		cell := innerRegionIt.Get()

		if cell.GetValue() == inRaster.Nodata {
			continue
		}

		neighbors := raster.NewNeighborIteratorWithCell(inRaster, cell)

		isSink := true
		filledZ := math.MaxFloat64
		for neighbors.Next() {
			ncell := neighbors.Get()

			if ncell.GetValue() == inRaster.Nodata ||
				ncell.GetValue() <= cell.GetValue() {
				isSink = false
				break
			}

			filledZ = math.Min(filledZ, ncell.GetValue())
		}

		if isSink {
			out.SetWithCell(cell, filledZ)
		}
	}

	// compute flow direction of the cells on the edge
	edges := raster.NewBorderIterator(out)

	for edges.Next() {
		cell := edges.Get()
		closed.SetWithCell(cell)

		if cell.GetValue() == out.Nodata {
			direction.SetWithCell(cell, int(raster.None))
		} else {
			if forceFlow {
				direction.SetWithCell(cell, int(cell.EdgeDirection(out)))
			} else {
				dir, _ := findCellDirection(cell, out)
				direction.SetWithCell(cell, int(dir))
			}
		}
	}

	return nil, nil, nil
}

func findCellDirection(c *raster.Cell, r *raster.Raster) (raster.Direction, float64) {
	if c.GetValue() == r.Nodata {
		return raster.None, -1
	}

	neighbors := raster.NewNeighborIteratorWithCell(r, c)

	dir := raster.None
	steepestSlope := 0.0
	var slope float64

	for neighbors.Next() {
		ncell := neighbors.Get()

		if ncell == nil {
			continue
		}

		if ncell.GetValue() == r.Nodata {
			continue
		}

		nDirection := c.RelativeDirection(ncell)
		elevDiff := c.GetValue() - ncell.GetValue()

		if nDirection.IsDiagonal() {
			slope = elevDiff / math.Sqrt(r.CellXSize*r.CellXSize+
				r.CellYSize*r.CellYSize)
		} else if nDirection == raster.Left || nDirection == raster.Right {
			slope = elevDiff / math.Abs(r.CellXSize)
		} else {
			slope = elevDiff / math.Abs(r.CellYSize)
		}

		if slope > steepestSlope {
			steepestSlope = slope
			dir = nDirection
		}
	}

	return dir, steepestSlope
}
