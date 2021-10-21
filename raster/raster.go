package raster

import (
	"fmt"
	"strings"
)

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

func (r *Raster) GetWithCell(c *Cell) float64 {
	return r.Data[c.Yindex*r.Xsize+c.Xindex]
}

func (r *Raster) Set(x, y int, val float64) {
	r.Data[y*r.Xsize+x] = val
}

func (r *Raster) SetWithCell(c *Cell, val float64) {
	r.Data[c.Yindex*r.Xsize+c.Xindex] = val
}

func (r *Raster) IsInRegion(x, y int) bool {
	if (x >= 0 && x < r.Xsize) && (y >= 0 && y < r.Ysize) {
		return true
	} else {
		return false
	}
}

func (r *Raster) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Xsize\t\t\t%d\n", r.Xsize))
	sb.WriteString(fmt.Sprintf("Ysize\t\t\t%d\n", r.Ysize))
	sb.WriteString(fmt.Sprintf("Nodata\t\t\t%f\n", r.Nodata))
	sb.WriteString("\n")

	for y := 0; y < r.Ysize; y++ {
		for x := 0; x < r.Xsize; x++ {
			sb.WriteString(fmt.Sprintf("%f", r.Get(x, y)))
			if x != r.Xsize-1 {
				sb.WriteString("\t")
			}
		}
		if y != r.Ysize-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

type Bitmap struct {
	Data  []bool
	Xsize int
	Ysize int
	Size  int
}

func NewBitmap(xsize, ysize int) *Bitmap {
	bm := Bitmap{Data: make([]bool, xsize*ysize), Xsize: xsize, Ysize: ysize,
		Size: xsize * ysize}
	return &bm
}

func NewBitmapWithRaster(r *Raster) *Bitmap {
	bm := NewBitmap(r.Xsize, r.Ysize)
	return bm
}

func (bm *Bitmap) Get(x, y int) bool {
	return bm.Data[y*bm.Xsize+x]
}

func (bm *Bitmap) GetWithCell(c *Cell) bool {
	return bm.Data[c.Yindex*bm.Xsize+c.Xindex]
}

func (bm *Bitmap) Set(x, y int) {
	bm.Data[y*bm.Xsize+x] = true
}

func (bm *Bitmap) SetWithCell(c *Cell) {
	bm.Data[c.Yindex*bm.Xsize+c.Xindex] = true
}

func (bm *Bitmap) Unset(x, y int) {
	bm.Data[y*bm.Xsize+x] = false
}

func (bm *Bitmap) UnsetWithCell(c *Cell) {
	bm.Data[c.Yindex*bm.Xsize+c.Xindex] = false
}

func (bm *Bitmap) IsInRegion(x, y int) bool {
	if (x >= 0 && x < bm.Xsize) && (y >= 0 && y < bm.Ysize) {
		return true
	} else {
		return false
	}
}

func (bm *Bitmap) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Xsize\t\t\t%d\n", bm.Xsize))
	sb.WriteString(fmt.Sprintf("Ysize\t\t\t%d\n", bm.Ysize))
	sb.WriteString("\n")

	for y := 0; y < bm.Ysize; y++ {
		for x := 0; x < bm.Xsize; x++ {
			sb.WriteString(fmt.Sprintf("%t", bm.Get(x, y)))
			if x != bm.Xsize-1 {
				sb.WriteString("\t")
			}
		}
		if y != bm.Ysize-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

type Intmap struct {
	Data  []int
	Xsize int
	Ysize int
	Size  int
}

func NewIntmap(xsize, ysize int) *Intmap {
	bm := Intmap{Data: make([]int, xsize*ysize), Xsize: xsize, Ysize: ysize,
		Size: xsize * ysize}
	return &bm
}

func NewIntmapWithRaster(r *Raster) *Intmap {
	bm := NewIntmap(r.Xsize, r.Ysize)
	return bm
}

func (bm *Intmap) Get(x, y int) int {
	return bm.Data[y*bm.Xsize+x]
}

func (bm *Intmap) GetWithCell(c *Cell) int {
	return bm.Data[c.Yindex*bm.Xsize+c.Xindex]
}

func (bm *Intmap) Set(x, y, val int) {
	bm.Data[y*bm.Xsize+x] = val
}

func (bm *Intmap) SetWithCell(c *Cell, val int) {
	bm.Data[c.Yindex*bm.Xsize+c.Xindex] = val
}

func (bm *Intmap) IsInRegion(x, y int) bool {
	if (x >= 0 && x < bm.Xsize) && (y >= 0 && y < bm.Ysize) {
		return true
	} else {
		return false
	}
}

func (bm *Intmap) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Xsize\t\t\t%d\n", bm.Xsize))
	sb.WriteString(fmt.Sprintf("Ysize\t\t\t%d\n", bm.Ysize))
	sb.WriteString("\n")

	for y := 0; y < bm.Ysize; y++ {
		for x := 0; x < bm.Xsize; x++ {
			sb.WriteString(fmt.Sprintf("%d", bm.Get(x, y)))
			if x != bm.Xsize-1 {
				sb.WriteString("\t")
			}
		}
		if y != bm.Ysize-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
