package main 

import "fmt"

func main() {
  var evenOddNumber int 
	fmt.Println("Please enter the number you want to trade : ")
	fmt.Scanln(&evenOddNumber)
   if evenOddNumber %2 == 0 {
      fmt.Println("That's an even number.")
   } else {
	fmt.Println("That's an odd number.")
   }
}