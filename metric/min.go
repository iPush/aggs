package metric

import (
	"math"

	"github.com/iPush/aggs/model"
)

type MinAggs struct {
}

func (a *MinAggs) Do(input model.Points) model.Point {
	//TODO: init max value to lower bound of float64
	var min model.Point = math.MaxFloat64
	for _, i := range input {
		if i < min {
			min = i
		}
	}
	return min
}
