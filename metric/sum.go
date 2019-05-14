package metric

import (
	"github.com/iPush/aggs/model"
)

type SumAggs struct{}

func (sa *SumAggs) Do(input model.Points) model.Point {
	var res model.Point
	for _, i := range input {
		res += i
	}
	return res
}
