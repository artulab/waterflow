package raster

type Raster struct {
	Data   []float64
	Xsize  int
	Ysize  int
	Size   int
	Nodata float64
}

type Cell struct {
	Value  *float64
	Xindex int
	Yindex int
}

func NewRaster(xsize, ysize int, noData float64) *Raster {
	r := Raster{Data: make([]float64, xsize*ysize), Xsize: xsize, Ysize: ysize,
		Size: xsize * ysize, Nodata: noData}
	return &r
}

func CopyRaster(r *Raster) *Raster {
	cr := Raster{Data: make([]float64, r.Xsize*r.Ysize), Xsize: r.Xsize, Ysize: r.Ysize,
		Size: r.Xsize * r.Ysize, Nodata: r.Nodata}
	copy(cr.Data, r.Data)

	return &cr
}

func (r *Raster) Get(x, y int) float64 {
	return r.Data[y*r.Xsize+x]
}

func (r *Raster) Set(x, y int, val float64) {
	r.Data[y*r.Xsize+x] = val
}

func (r *Raster) IsInRegion(x, y int) bool {
	if (x >= 0 && x < r.Xsize) && (y >= 0 && y < r.Ysize) {
		return true
	} else {
		return false
	}
}

type BitMap struct {
	Data  []bool
	Xsize int
	Ysize int
	Size  int
}

func NewBitmap(xsize, ysize int) *BitMap {
	bm := BitMap{Data: make([]bool, xsize*ysize), Xsize: xsize, Ysize: ysize,
		Size: xsize * ysize}
	return &bm
}

func NewBitmapWithRaster(r *Raster) *BitMap {
	bm := NewBitmap(r.Xsize, r.Ysize)
	return bm
}

func (bm *BitMap) Get(x, y int) bool {
	return bm.Data[y*bm.Xsize+x]
}

func (bm *BitMap) GetWithCell(c *Cell) bool {
	return bm.Data[c.Yindex*bm.Xsize+c.Xindex]
}

func (bm *BitMap) Set(x, y int) {
	bm.Data[y*bm.Xsize+x] = true
}

func (bm *BitMap) SetWithCell(c *Cell) {
	bm.Data[c.Yindex*bm.Xsize+c.Xindex] = true
}

func (bm *BitMap) Unset(x, y int) {
	bm.Data[y*bm.Xsize+x] = false
}

func (bm *BitMap) UnsetWithCell(c *Cell) {
	bm.Data[c.Yindex*bm.Xsize+c.Xindex] = false
}
