package main 

import "fmt"

func main() {

var list []int = make([]int, 0)
fmt.Println("Please enter 5 number :")
 for i := 0; i < 5; i++ {
	number := 0
	fmt.Scanf("%d", &number)
	list= append(list, number)
}

 
 for i := 0; i < len(list); i++ {

	for j:= i; j < len(list); j++ {

      if list[j] > list[i] {
      
	list[j], list[i] = list[i] , list[j]
 	  }
	}
 } 
 fmt.Println(list)

} 



