package flow

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/artulab/waterflow/raster"
)

func TestFlowWithSink(t *testing.T) {
	r := raster.NewRaster(4, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 7,
		8, 5, 1, 8,
		8, 9, 8, 9,
	}

	params := FlowDirectionParameters{InRaster: r, ForceFlow: false,
		ComputeDrop: false}
	out, err := FlowDirection(params)

	if err != nil {
		t.Error("error isn't expected")
	}

	expected := []int{
		2, 4, 4, 8,
		1, 1, 16, 16,
		128, 64, 64, 32,
	}

	if reflect.DeepEqual(expected, out.FlowDirectionRaster.Data) != true {
		t.Error("flow result isn't expected")
	}
}

func TestFlowWithSinkAndForceFlow(t *testing.T) {
	r := raster.NewRaster(4, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 7,
		8, 5, 1, 8,
		8, 9, 8, 9,
	}

	params := FlowDirectionParameters{InRaster: r, ForceFlow: true,
		ComputeDrop: false}
	out, err := FlowDirection(params)

	if err != nil {
		t.Error("error isn't expected")
	}

	expected := []int{
		32, 64, 64, 128,
		16, 1, 16, 1,
		8, 4, 4, 2,
	}

	if reflect.DeepEqual(expected, out.FlowDirectionRaster.Data) != true {
		t.Error("flow result isn't expected")
	}
}

func TestFlowWithNoSinkAndNoDataAndMultiDirEdge(t *testing.T) {
	r := raster.NewRaster(4, 3, 1, 1, -9999)
	r.Data = []float64{
		9, 8, 8, 7,
		8, 5, -9999, 8,
		8, 9, 8, 9,
	}

	params := FlowDirectionParameters{InRaster: r, ForceFlow: false,
		ComputeDrop: false}
	out, err := FlowDirection(params)

	if err != nil {
		t.Error("error isn't expected")
	}

	fmt.Println(out.FlowDirectionRaster.String())

	// mine:
	// 2, 4, 8, 128,
	// 1, 0, 0, 64,
	// 128, 64, 32, 80,
	// 80 is wrong and its cell is not in sink, compute based on heuristic

	expected := []int{
		2, 4, 8, 128,
		1, 1, 0, 64,
		128, 64, 32, 16,
	}

	if reflect.DeepEqual(expected, out.FlowDirectionRaster.Data) != true {
		t.Error("flow result isn't expected")
	}
}

func TestFlowWithNoSinkAndNoDataAndMultiDirEdge2(t *testing.T) {
	r := raster.NewRaster(4, 4, 1, 1, 0)
	r.Data = []float64{
		9, 8, 8, 7,
		8, 0, 6, 8,
		8, 9, 5, 9,
		8, 9, 8, 9,
	}

	params := FlowDirectionParameters{InRaster: r, ForceFlow: false,
		ComputeDrop: false}
	out, err := FlowDirection(params)

	if err != nil {
		t.Error("error isn't expected")
	}

	fmt.Println(out.FlowDirectionRaster.String())

	// mine:
	// 5       2       4       8
	// 132     0       4       8
	// 68      1       0       16
	// 64      128     64      32

	expected := []int{
		4, 2, 4, 8,
		16, 0, 4, 8,
		16, 1, 32, 16,
		8, 128, 64, 32,
	}

	if reflect.DeepEqual(expected, out.FlowDirectionRaster.Data) != true {
		t.Error("flow result isn't expected")
	}
}
