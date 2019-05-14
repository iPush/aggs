package model

type UnaryFn func(Point) Point

// Map apply fn for each element of Points
func Map(fn UnaryFn, points Points) Points {
	if len(points) == 0 {
		return []Point{}
	}

	res := make([]Point, len(points))
	for i, pt := range points {
		res[i] = fn(pt)
	}
	return res
}

type UnaryPredicater func(Point) bool

func Filter(fn UnaryPredicater, pts Points) Points {
	var res Points
	for _, pt := range pts {
		if fn(pt) {
			res = append(res, pt)
		}
	}
	return res
}

type BinaryFn func(x, y Point) Point

func Reduce(fn BinaryFn, pts Points, init Point) Point {
	if pts == nil || len(pts) == 0 {
		return init
	}

	it := pts.CreateIterator()
	for it.HasNext() {
		cur := it.Next()
		init = fn(init, cur)
	}

	return init

}
