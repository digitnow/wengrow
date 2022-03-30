package aen

func AverageOfEvenNumbers(array []int) int {
	var sum int = 0
	var countOfEvenNumbers int = 0
	for _, number := range array {
		if number % 2 == 0 {
			sum += number
			countOfEvenNumbers++
		}
	}
	return sum / countOfEvenNumbers
}