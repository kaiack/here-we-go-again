package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbers ...[]int) []int {
	var totals []int

	for _, numList := range numbers {
		var tailSum int
		if len(numList) == 0 {
			tailSum = 0
		} else {
			tailSum = Sum(numList[1:])
		}
		totals = append(totals, tailSum)
	}

	return totals
}
