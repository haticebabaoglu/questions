package main 

import "fmt"

func main() {	
var finalGrade float32 
var midtermGrade float32  	
	fmt.Println("Vize notunu giriniz :")
	fmt.Scanln(&midtermGrade)
	fmt.Println("Final notunu giriniz :")
	fmt.Scanln(&finalGrade)
     finalGradeAverage := (finalGrade *70 ) / 100
     midtermGradeAverage := (midtermGrade *30) / 100
     total := finalGradeAverage + midtermGradeAverage
     fmt.Println("Not ortalaması:")
     fmt.Println(total)


}
