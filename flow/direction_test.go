package flow

import (
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
