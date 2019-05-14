package model

type Point float64
type Points []Point

type PointsIter struct {
	pts Points
	pos int
}

func (pts Points) CreateIterator() Iterator {
	return &PointsIter{
		pts: pts,
	}
}

func (pi *PointsIter) HasNext() bool {
	return pi.pos < len(pi.pts)
}

func (pi *PointsIter) Next() Point {
	res := pi.pts[pi.pos]
	pi.pos += 1
	return res
}
