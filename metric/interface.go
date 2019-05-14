package metric

import (
	"github.com/iPush/aggs/model"
)

type MetricAggs interface {
	Do(input model.Points) (ouput model.Point)
}
