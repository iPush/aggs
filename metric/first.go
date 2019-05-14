package metric

import "github.com/iPush/aggs/model"

type FirstAggs struct{}

func (a *FirstAggs) Do(input model.Points) model.Point {
	return input[0]
}
