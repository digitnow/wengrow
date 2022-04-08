package quicks

import (
	"testing"
	//"reflect"
)

func TestPartition(t *testing.T) {
	testData := []struct{
		inputArray []int
		leftPointer int
		rightPointer int
		wanted int
	}{
		{inputArray: []int{0,5,2,1,6,3}, leftPointer: 0, rightPointer: 5, wanted: 3},
		{inputArray: []int{0,1,2,3,6,5}, leftPointer: 3, rightPointer: 5, wanted: 4},
		{inputArray: []int{0,1,2,3,5,6}, leftPointer: 4, rightPointer: 5, wanted: 5},
	}
	for _, td := range testData {
		got := Partition(td.inputArray, td.leftPointer, td.rightPointer)
		if td.wanted != got {
		//if !reflect.DeepEqual(td.wanted, got) {
			t.Fatalf("expected %v, got: %v", td.wanted, got)
		}
	}
}

//Skriv testen til Quicksort selv her:
/*
func TestQuicksort(t *testing.T) {
	testData := []struct {...}
	...
	for _, td := range testData {
		Quicksort(td.inputArray, ...)
		got := td.inputArray
		if !reflect.DeepEqual(td.wanted, got) {
			...
		}
	} 
}
*/