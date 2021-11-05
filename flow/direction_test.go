package flow

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/artulab/waterflow/raster"
)

func TestFlow(t *testing.T) {
	r := raster.NewRaster(4, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 7,
		8, 5, 1, 8,
		8, 9, 8, 9,
	}

	directions, _, err := FlowDirection(r, false, false)

	if err != nil {
		t.Error("error isn't expected")
	}

	expected := []int{
		2, 4, 4, 8,
		1, 1, 16, 16,
		128, 64, 64, 32,
	}

	if reflect.DeepEqual(expected, directions.Data) != true {
		t.Error("flow result isn't expected")
	}
}

func TestFlowWithForceFlow(t *testing.T) {
	r := raster.NewRaster(4, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 7,
		8, 5, 1, 8,
		8, 9, 8, 9,
	}

	directions, _, err := FlowDirection(r, true, false)

	if err != nil {
		t.Error("error isn't expected")
	}

	expected := []int{
		32, 64, 64, 128,
		16, 1, 16, 1,
		8, 4, 4, 2,
	}

	if reflect.DeepEqual(expected, directions.Data) != true {
		t.Error("flow result isn't expected")
	}
}

func TestFlowWithNoData(t *testing.T) {
	r := raster.NewRaster(4, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 7,
		8, 5, -9999, 8,
		8, 9, 8, 9,
	}

	directions, _, err := FlowDirection(r, false, false)

	if err != nil {
		t.Error("error isn't expected")
	}

	fmt.Println(directions.String())

	// mine:
	// 2, 4, 8, 0,
	// 1, 0, 0, 64,
	// 128, 64, 32, 80,
	// 80 is wrong and its cell is not in sink, compute based on heuristic

	expected := []int{
		2, 4, 8, 128,
		1, 1, 0, 64,
		128, 64, 32, 16,
	}

	if reflect.DeepEqual(expected, directions.Data) != true {
		t.Error("flow result isn't expected")
	}
}
