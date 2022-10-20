package main

import (
	"strconv"
)

func thirdExamSolution(S string) int {
	countDivisibleByThree := 0
	numbersDivisionByThree := map[int]int{}

	for i := 0; i < len(S); i++ {
		for j := 0; j < 10; j++ {
			isDivisibleByThree, _ := strconv.Atoi(S[:i] + strconv.Itoa(j) + S[i+1:])

			if isDivisibleByThree%3 == 0 {
				if numbersDivisionByThree[isDivisibleByThree] > 0 {
					numbersDivisionByThree[isDivisibleByThree]++
				} else {
					numbersDivisionByThree[isDivisibleByThree] = 1
				}
				countDivisibleByThree++
			}
		}
	}

	for _, frequency := range numbersDivisionByThree {
		if frequency > 1 {
			countDivisibleByThree = countDivisibleByThree - frequency + 1
		}
	}

	return countDivisibleByThree
}
