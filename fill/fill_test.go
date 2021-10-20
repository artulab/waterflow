package fill

import (
	"fmt"
	"testing"

	"github.com/artulab/waterflow/raster"
)

func TestFill(t *testing.T) {
	r := raster.NewRaster(4, 3, -9999)
	r.Data = []float64{
		4, 5, 4, 3,
		4, 1, 2, 7,
		8, 9, 8, 9,
	}

	out, err := Fill(r, 0)

	if err != nil {
		t.Error("error isn't expected")
	}

	fmt.Println(out.String())
}
