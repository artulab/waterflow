package raster

type Raster struct {
	Data   []float32
	Xsize  int
	Ysize  int
	Size   int
	Nodata float32
}

type Cell struct {
	Value  *float32
	Xindex int
	Yindex int
}

func New(xsize, ysize int, noData float32) *Raster {
	r := Raster{Data: make([]float32, xsize*ysize), Xsize: xsize, Ysize: ysize,
		Size: xsize * ysize, Nodata: noData}
	return &r
}

func (r *Raster) Get(x, y int) float32 {
	return r.Data[y*r.Xsize+x]
}

func (r *Raster) Set(x, y int, val float32) {
	r.Data[y*r.Xsize+x] = val
}

func (r *Raster) IsInRegion(x, y int) bool {
	if (x >= 0 && x < r.Xsize) && (y >= 0 && y < r.Ysize) {
		return true
	} else {
		return false
	}
}
