package main 

import "fmt"

func main() {
 var array [10]int = [10]int{20, 4, 5, 7, 11, 13, 2, 24, 65, 21}
	minimumNumber := array[0]
	for i := 0; i<len(array); i++{
		if array[i] < minimumNumber {
			minimumNumber = array[i]
		}
	}
fmt.Println(minimumNumber)

}