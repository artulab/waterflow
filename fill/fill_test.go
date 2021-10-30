package fill

import (
	"reflect"
	"testing"

	"github.com/artulab/waterflow/raster"
)

func TestFill(t *testing.T) {
	r := raster.NewRaster(4, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 7,
		8, 5, 1, 8,
		8, 9, 8, 9,
	}

	out, err := Fill(r, 0)

	if err != nil {
		t.Error("error isn't expected")
	}

	expected := []float64{
		9, 8, 8, 7,
		8, 7, 7, 8,
		8, 9, 8, 9,
	}

	if reflect.DeepEqual(expected, out.Data) != true {
		t.Error("fill result isn't expected")
	}
}

func TestFillWithZLimit(t *testing.T) {
	r := raster.NewRaster(4, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 7,
		8, 1, 6, 8,
		8, 9, 8, 9,
	}

	out, err := Fill(r, 4)

	if err != nil {
		t.Error("error isn't expected")
	}

	expected := []float64{
		9, 8, 8, 7,
		8, 1, 6, 8,
		8, 9, 8, 9,
	}

	if reflect.DeepEqual(expected, out.Data) != true {
		t.Error("fill result isn't expected")
	}
}

func TestFillWithZLimitAnd2Sinks(t *testing.T) {
	r := raster.NewRaster(7, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 9, 8, 9, 8,
		7, 2, 6, 8, 6, 6, 9,
		8, 9, 8, 9, 8, 9, 7,
	}

	out, err := Fill(r, 4)

	if err != nil {
		t.Error("error isn't expected")
	}

	expected := []float64{
		9, 8, 8, 9, 8, 9, 8,
		7, 2, 6, 8, 7, 7, 9,
		8, 9, 8, 9, 8, 9, 7,
	}

	if reflect.DeepEqual(expected, out.Data) != true {
		t.Error("fill result isn't expected")
	}
}

func TestFillWithZLimitAnd2SinksEdgeCase(t *testing.T) {
	r := raster.NewRaster(7, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 9, 8, 9, 8,
		7, 3, 6, 8, 6, 6, 9,
		8, 9, 8, 9, 8, 9, 7,
	}

	out, err := Fill(r, 4)

	if err != nil {
		t.Error("error isn't expected")
	}

	expected := []float64{
		9, 8, 8, 9, 8, 9, 8,
		7, 3, 6, 8, 7, 7, 9,
		8, 9, 8, 9, 8, 9, 7,
	}

	if reflect.DeepEqual(expected, out.Data) != true {
		t.Error("fill result isn't expected")
	}
}

func TestFillWithZLimitAnd2SinksAllFilled(t *testing.T) {
	r := raster.NewRaster(7, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 9, 8, 9, 8,
		7, 4, 6, 8, 6, 6, 9,
		8, 9, 8, 9, 8, 9, 7,
	}

	out, err := Fill(r, 4)

	if err != nil {
		t.Error("error isn't expected")
	}

	expected := []float64{
		9, 8, 8, 9, 8, 9, 8,
		7, 7, 7, 8, 7, 7, 9,
		8, 9, 8, 9, 8, 9, 7,
	}

	if reflect.DeepEqual(expected, out.Data) != true {
		t.Error("fill result isn't expected")
	}
}

func TestFillWithZLimitNoChange(t *testing.T) {
	r := raster.NewRaster(4, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 7,
		8, 5, 1, 8,
		8, 9, 8, 9,
	}

	out, err := Fill(r, 4)

	if err != nil {
		t.Error("error isn't expected")
	}

	expected := []float64{
		9, 8, 8, 7,
		8, 5, 1, 8,
		8, 9, 8, 9,
	}

	if reflect.DeepEqual(expected, out.Data) != true {
		t.Error("fill queue result isn't expected")
	}
}
