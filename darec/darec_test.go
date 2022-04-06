package darec

import (
	"testing"
	"reflect"
)

func TestDoubleArray(t *testing.T) {
	testData := []struct{
		input, wanted []int
	}{
		{input: []int{1,2,3,4,5}, wanted: []int{2,4,6,8,10}},
		{input: []int{-4,-3,-2,-1,0}, wanted: []int{-8,-6,-4,-2,0}},
	}
	for _, td := range testData {
		DoubleArray(td.input, -6)
		got := td.input
		t.Log(td.input)
		if !reflect.DeepEqual(td.wanted, got) {
			t.Fatalf("expected %v, got: %v", td.wanted, got)
		}
	}
}  