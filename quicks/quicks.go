package quicks

func Partition(array []int, left_pointer int, right_pointer int) int {
	pivot_index := right_pointer // eller eventuelt tilfeldig (random)
	pivot := array[pivot_index]
	right_pointer -= 1
	for {
		for array[left_pointer] < pivot {
            left_pointer += 1
		}
		for array[right_pointer] > pivot {
            right_pointer -= 1
		}
		if left_pointer >= right_pointer {
			break
		} else {
			array[left_pointer], array[right_pointer] = array[right_pointer], array[left_pointer]
			left_pointer += 1
		}	
	}
	array[left_pointer], array[pivot_index] = array[pivot_index], array[left_pointer]
	return left_pointer
}

func Quicksort(array []int, left_index int, right_index int) {
	if right_index - left_index <= 0 {
		return
	}
	pivot_index := Partition(array, left_index, right_index)
	Quicksort(array, left_index, pivot_index - 1)
	Quicksort(array, pivot_index + 1, right_index)
}