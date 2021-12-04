package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(solution([]int{3, 6, -2, -5, 7, 3}))
	fmt.Println(solution2([]int{3, 6, -2, -5, 7, 3}))
}

func solution(inputArray []int) int {

	var pairNumbers [][]int
	for index, value := range inputArray {

		if len(inputArray) > index+1 {

			var (
				actualNumberPosition   = value
				nextPairNumberPosition = inputArray[index+1]
			)

			pairNumbers = append(pairNumbers, []int{actualNumberPosition, nextPairNumberPosition})
		}

	}

	var results []int
	for _, pair := range pairNumbers {
		results = append(results, pair[0]*pair[1])
	}

	sort.Ints(results)
	return results[len(results)-1]
}

func solution2(inputArray []int) int {
	max := inputArray[0] * inputArray[1]

	for i := 1; i < len(inputArray)-1; i++ {
		if max < inputArray[i]*inputArray[i+1] {
			max = inputArray[i] * inputArray[i+1]
		}
	}

	return max
}
