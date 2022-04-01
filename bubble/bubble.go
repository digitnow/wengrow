package bubble

func BubbleSort(array []int) []int {
	unsortedUntilIndex := len(array) - 1

	for sorted := false; !sorted; {
		sorted = true
		for i, _ := range array[0:unsortedUntilIndex] {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				sorted = false
			}
		}
		unsortedUntilIndex--
	}
	return array
}