package main 

import "fmt"

func main() {
	arr := [10]int{20, 0, 5, 7, 11, 13, 45, 52, 12, 42}
		sum := 0
		for i := 0; i < len(arr); i++ {
			sum += arr[i]
		}
		average := (float32(sum)) / 10
		fmt.Println("Sum = ", sum, "Average = ", average)
}
