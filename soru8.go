package main 

import "fmt"

var weight float32
var height float32

func main(){
	fmt.Println("Please enter your height:")
	fmt.Scanln(&weight)
	fmt.Println("Please enter your weight:")
	fmt.Scanln(&height)

    weight /= (height * height) 

     
	if weight < 18.5 {
		fmt.Println("Thin. Your fat persentage: ", weight)
	} else if weight <= 18.5 && weight <= 24.9 {
		fmt.Println("Normal. Your fat persentage: ", weight)
	} else if weight <= 25 && weight <= 29.9 {
		fmt.Println("Fatty. Your fat persentage: " ,weight)
	} else if weight <= 30 && weight <= 34.9 {
		fmt.Println("1st degree obese. Your fat persentage: ", weight)
	} else if weight <= 35 && weight <= 39.9 {
		fmt.Println("2st degree obese. Your fat persentage: " ,weight)
	} else  if weight  > 40{
		fmt.Println("3st degree obese. Your fat persentage: ", weight)
	} else {
		fmt.Println("Please enter a valid Height and Weight.")
	}

}
