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
	computeDrop bool) (*raster.Intmap, *raster.Raster, error) {

	out := raster.CopyRaster(inRaster)
	slopes := raster.NewRasterWithRaster(inRaster)
	directions := raster.NewIntmapWithRaster(inRaster)

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

		if cell.GetValue() == out.Nodata {
			directions.SetWithCell(cell, int(raster.None))
		} else {
			if forceFlow {
				directions.SetWithCell(cell, int(cell.EdgeDirection(out)))
				slopes.SetWithCell(cell, 0)
			} else {
				dir, slope := findCellDirection(cell, out)
				directions.SetWithCell(cell, int(dir))
				slopes.SetWithCell(cell, slope)
			}
		}
	}

	// compute flow direction of the cell inside of the raster
	innerRegionIt = raster.NewInnerRegionIterator(out)

	for innerRegionIt.Next() {
		cell := innerRegionIt.Get()

		dir, slope := findCellDirection(cell, out)
		directions.SetWithCell(cell, int(dir))
		slopes.SetWithCell(cell, slope)
	}

	return directions, slopes, nil
}

func findCellDirection(c *raster.Cell, r *raster.Raster) (raster.Direction, float64) {
	if c.GetValue() == r.Nodata {
		return raster.None, -1
	}

	neighbors := raster.NewNeighborIteratorWithCell(r, c)

	dir := 0
	steepestSlope := 0.0
	var slope float64

	downSlopes := [8]float64{-1, -1, -1, -1, -1, -1, -1, -1}
	nIndex := -1

	for neighbors.Next() {
		ncell := neighbors.Get()

		nIndex++

		if ncell == nil || ncell.GetValue() == r.Nodata {
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

		if slope >= 0 {
			downSlopes[nIndex] = slope
		}

		if slope >= steepestSlope {
			steepestSlope = slope
		}
	}

	// combine flow directions if there's multiple
	// neighbor downslope cells with same elevation
	for i, val := range downSlopes {
		if val == steepestSlope {
			dir |= (1 << i)
		}
	}

	return raster.Direction(dir), steepestSlope
}
