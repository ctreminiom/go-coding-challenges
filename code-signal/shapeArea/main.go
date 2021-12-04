package main

import "fmt"

func main() {
	fmt.Println(solution(2))
}

func solution(n int) int {
	return (1+(n-1)*n)*2 - 1
}
