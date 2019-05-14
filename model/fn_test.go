package model

import "testing"

func add(x, y Point) Point {
	return x + y
}

type Case struct {
	input  Points
	expect Point
}

var addCases []Case = []Case{
	Case{
		input:  Points{1, 2, 3, 4},
		expect: 10,
	},
	Case{
		input:  Points{},
		expect: 0,
	},
}

func Test_Reduce(t *testing.T) {
	for _, cas := range addCases {
		actual := Reduce(add, cas.input, 0)
		t.Log(actual)
		if actual != cas.expect {
			t.Errorf("expect: %f. actual: %f", cas.expect, actual)
		}
	}
}
