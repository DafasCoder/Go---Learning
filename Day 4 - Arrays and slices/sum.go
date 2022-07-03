package main



func Sum(values [6] int) int {

	total := 0

	for _, value := range values {
		total += value
	}


	return total

}

func SumUnknown(values []int ) int {
	total := 0 
	for _, value := range values {
		total += value
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, SumUnknown(numbers))
	}
	
	return sums
}

func SumAllTails(numbersOfTails ...[]int) []int {
	var sums []int

	for _, numbers := range numbersOfTails{
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
		tails := numbers[1:]
		sums = append(sums, SumUnknown(tails))
		}	
	}

	return sums
}