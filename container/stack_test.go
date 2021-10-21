package container

import (
	"reflect"
	"testing"

	"github.com/artulab/waterflow/raster"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	actual := []float64{1, 2, 3, 4, 5}

	for i := 0; i < len(actual); i++ {
		stack.Push(&raster.Cell{Value: &actual[i], Xindex: i, Yindex: 0})
	}

	expected := make([]float64, len(actual))

	for i := len(actual) - 1; i >= 0; i-- {
		val, err := stack.Pop()
		if err != nil {
			t.Error("error isn't expected")
		}
		expected[i] = *val.Value
	}

	if reflect.DeepEqual(expected, actual) != true {
		t.Error("stack result isn't expected")
	}

	_, err := stack.Pop()

	if err == nil {
		t.Error("error is expected")
	}
}
