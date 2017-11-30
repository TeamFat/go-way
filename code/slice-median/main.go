package main

import "fmt"

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}
func main() {
	num := []float64{3, 4, 5}
	median := median(num)
	fmt.Println(median)
	//https://plumwine.me/programming-in-go-booleans-and-numbers-implementing-statistics-functions/
}
