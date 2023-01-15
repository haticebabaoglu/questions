package main

import"fmt"

func main(){
	arr := []int{90, 50, 80, 70, 30, 30, 10, 80, 40, 50, 40, 30}
	frequency := make(map[int]int)
	for _ , num := range arr {
		frequency[num] = frequency[num]+1
	}
	fmt.Println("Frequency of the Array is : ", frequency)
}