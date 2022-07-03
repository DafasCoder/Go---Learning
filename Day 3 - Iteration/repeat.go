package main


	const repeatedCount = 5

	func Repeat(character string) int {
		var repeated string
		var counter = 0
		
		for i := 0; i < repeatedCount ; i++{
			repeated += character
			counter ++
		}
		return counter
	}