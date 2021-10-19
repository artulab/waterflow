package fill

import (
	"github.com/artulab/waterflow/raster"
)

// Fill attempts to correct cells of given inRaster by filling sinks/pits.
// zLimit refers to the maximum value between the original cell value and its
// filled value. Those sinks whose elevation difference is greater than zLimit
// will not be filled. If the zLimit is zero, all sinks will be filled.
func Fill(inRaster *raster.Raster, zLimit float32) (*raster.Raster, error) {
	return nil, nil
}
