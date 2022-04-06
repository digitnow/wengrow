package main

import (
	"fmt"
	"github.com/digitnow/wengrow/bubble"
)

func main() {
	array := []int{4,2,7,1,3}
	fmt.Println(array)
	fmt.Println(bubble.BubbleSort(array))
	fmt.Println(array)
}