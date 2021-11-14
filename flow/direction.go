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

type FlowDirectionParameters struct {
	InRaster    *raster.Raster
	ForceFlow   bool
	ComputeDrop bool
}

type FlowDirectionResult struct {
	FlowDirectionRaster *raster.Intmap
	SlopeRaster         *raster.Raster
}

func FlowDirection(param FlowDirectionParameters) (*FlowDirectionResult, error) {

	out := raster.CopyRaster(param.InRaster)
	slopes := raster.NewRasterWithRaster(param.InRaster)
	directions := raster.NewIntmapWithRaster(param.InRaster)

	// fill one-cell sinks
	innerRegionIt := raster.NewInnerRegionIterator(param.InRaster)

	for innerRegionIt.Next() {
		cell := innerRegionIt.Get()

		if cell.GetValue() == param.InRaster.Nodata {
			continue
		}

		neighbors := raster.NewNeighborIteratorWithCell(param.InRaster, cell)

		isSink := true
		filledZ := math.MaxFloat64
		for neighbors.Next() {
			ncell := neighbors.Get()

			if ncell.GetValue() == param.InRaster.Nodata ||
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
			if param.ForceFlow {
				directions.SetWithCell(cell, int(cell.EdgeDirection(out)))
				slopes.SetWithCell(cell, 0)
			} else {
				dir, slope := findCellDirection(cell, out)

				// if can not determine the flow on the edge
				// force flow outward from the raster
				if dir == raster.None {
					directions.SetWithCell(cell, int(cell.EdgeDirection(out)))
				} else {
					directions.SetWithCell(cell, int(dir))
				}
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

	return &FlowDirectionResult{FlowDirectionRaster: directions,
		SlopeRaster: slopes}, nil
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
