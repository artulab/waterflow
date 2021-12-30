# WaterFlow
![Version](https://img.shields.io/badge/version-v0.1.0-blue.svg?cacheSeconds=2592000)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> WaterFlow is a Go program that provides community-maintained free and open-source hydrology tools, implementing state-of-the-art algorithms found in the literature. WaterFlow tools aim to be API-compatible with [ESRI Hydrology toolset](https://desktop.arcgis.com/en/arcmap/latest/tools/spatial-analyst-toolbox/an-overview-of-the-hydrology-tools.htm) but do not necessarily guarantee to produce exactly the same output in the pixel/cell level as ESRI tools.

## Install

```sh
go get -v github.com/artulab/waterflow
```

## Usage

Import the package:
```go
import "github.com/artulab/waterflow"
```

## Tools

### Fill
Fill attempts to correct cells of given *inRaster* by filling sinks/pits.

#### Parameters
| Parameter | Description |
| ----------| ------------ |
| inRaster | Input raster data. |
| zLimit | Maximum elevation difference between the sink and its pour point. Those sinks whose elevation difference is greater than zLimit will not be filled. If the zLimit is zero, all sinks will be filled. |

#### Algorithm
[Barnes, R., Lehman, C., & Mulla, D. (2014). Priority-flood: An optimal depression-filling and watershed-labeling algorithm for digital elevation models. Computers & Geosciences, 62, 117-127](https://www.sciencedirect.com/science/article/pii/S0098300413001337)


### Direction (In Progress)
Computes the hydrologic flow directions on the input DEM that directs flow from each grid cell to one or more of its neighbors.

#### Parameters
| Parameter | Description |
| ----------| ------------ |
| inRaster | Input raster data. Note that the input raster does *not* need to be filled beforehand. |
| forceFlow | Determines if edge cells always flow outward or the direction is computed based on normal flow rules. |
| computeDrop | The ratio of the maximum change in elevation from each cell along the direction of flow to the path length between centers of cells, expressed in percentages. |
| skipFillingOneCellSinks | The tool fills one-cell sinks by default on *inRaster*. Set true if the input raster is already filled or this is not desired. |

#### Algorithm
[Survila, K., Yildirim, A. A., Li, T., Liu, Y. Y., Tarboton, D. G., & Wang, S. (2016, July). A scalable high-performance topographic flow direction algorithm for hydrological information analysis. In Proceedings of the XSEDE16 Conference on Diversity, Big Data, and Science at Scale (pp. 1-7).](https://dl.acm.org/doi/abs/10.1145/2949550.2949571)

## Run tests

```sh
go test
```

## Authors

**Ahmet Artu Yildirim**

* Website: https://www.artulab.com
* E-Mail: ahmet@artulab.com

## Contributing

Contributions, issues and feature requests are welcome!

Feel free to check [issues page](https://github.com/artulab/waterflow/issues). 

## Show your support

Give a ⭐️ if this project helped you!


## License

This project is [MIT](https://opensource.org/licenses/MIT) licensed.