package flow

import (
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

	expected := []float64{
		9, 8, 8, 7,
		8, 7, 7, 8,
		8, 9, 8, 9,
	}

	if reflect.DeepEqual(expected, directions.Data) != true {
		t.Error("flow result isn't expected")
	}
}
