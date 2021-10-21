package flow

import "github.com/artulab/waterflow/raster"

type Direction int

const (
	D8 Direction = iota
	MDF
	DInf
)

func flowDirection(r *raster.Raster,
	forceFlow bool, computeDrop bool) (*raster.Raster, *raster.Raster, error) {

	return nil, nil, nil
}
