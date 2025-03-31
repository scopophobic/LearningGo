
package main

 

import (
	"fmt"
	"learningGoLang/Day1/calculator"
)

const (
	AD_age string = "22"
)
// func main(){
// 	fmt.Println("Hello - World");
// 	// fmt.Println("I wanna ")
// 	var name,username string = "Adithyan","Scopophobic"
// 	fmt.Println(name + " - >> " + username )
// 	cool_name := "AD"
	

// 	fmt.Println("better short name : " + cool_name + " and age :" + AD_age)

	
// 	x := 20 

// 	if x>10 {
// 		fmt.Println("it is greater than 10")
// 		if x == 20 {
// 			fmt.Println("got it")
// 		}else {
// 			fmt.Println("oops oops, no no ")
// 		}
// 	} else {
// 		fmt.Println("YOOO");
// 	}


// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}

// }


func main(){
	fmt.Println("Welcome to the dummy calculator")
	calculator.Cal()
}