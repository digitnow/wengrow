package darec

func DoubleArray(array []int, index int) {
	if index >= len(array) { return }
	array[index] *= 2
	DoubleArray(array, index + 1)
}