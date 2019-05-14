package metric

import (
	"testing"

	"github.com/iPush/aggs/model"
)

var inputs = []model.Points{
	model.Points{1, 2, 3, 4, 5},
	model.Points{-1, -2, -3, -4, 1000},
	model.Points{1, 1, 1},
	model.Points{},
}

func Test_Max(t *testing.T) {
	aggs := &MaxAggs{}
	for _, input := range inputs {
		m := aggs.Do(input)
		t.Log(m)
	}
}
