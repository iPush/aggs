package model

import "testing"

var testPoints = Points{0.1, 0.2, 0.3, 0.4, 1.0}

func Test_Iter(t *testing.T) {
	it := testPoints.CreateIterator()
	for it.HasNext() {
		t.Log(it.Next())
	}
}
