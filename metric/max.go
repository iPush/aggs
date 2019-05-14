package metric

import (
	"math"

	"github.com/iPush/aggs/model"
)

type MaxAggs struct {
}

func (a *MaxAggs) Do(input model.Points) model.Point {
	//TODO: init max value to lower bound of float64
	var max model.Point = -math.SmallestNonzeroFloat64
	for _, i := range input {
		if i > max {
			max = i
		}
	}
	return max
}
