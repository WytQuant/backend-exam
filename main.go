package main

import "fmt"

func main() {
	// Second Exam
	fmt.Println("----------------------- SECOND EXAM START HERE ------------------------")
	inputStringOfSecondExam := []string{"We test coders. Give us a try?", "Forget CVs..Save time . x x"}
	for index, inputString := range inputStringOfSecondExam {
		fmt.Println("Test case", index+1, "Result:", secondExamSolution(inputString))
	}

	//Third Exam
	fmt.Println("----------------------- THIRD EXAM START HERE ------------------------")
	inputStringOfThirdExam := []string{"23", "0081", "022"}
	for index, inputString := range inputStringOfThirdExam {
		fmt.Println("Test case", index+1, "Result:", thirdExamSolution(inputString))
	}

}
