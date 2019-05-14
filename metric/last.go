package metric

import "github.com/iPush/aggs/model"

type LastAggs struct{}

func (a *LastAggs) Do(input model.Points) model.Point {
	idx := len(input) - 1
	return input[idx]
}
