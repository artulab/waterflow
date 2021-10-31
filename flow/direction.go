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

	edges := raster.NewBorderIterator(out)

	for edges.Next() {
		cell := edges.Get()
		closed.SetWithCell(cell)

		if cell.GetValue() != inRaster.Nodata {
			if forceFlow {
				direction.SetWithCell(cell, int(cell.EdgeDirection(inRaster)))
			} else {
				dir, _ := findCellDirection(cell, inRaster)
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
