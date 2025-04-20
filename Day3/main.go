package main

import(
	"fmt"
)



func main(){
	fmt.Println("hello world!")
	// var n int = 4
	var arr [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	a := arr[1:5]
	fmt.Println(a)
	fmt.Println(arr)




	var m  = map[int]string {
		1 : "adithyan",
		4 : "love",
		3 : "priyanka",
		2 : "pia",
	}
	fmt.Println(1," ",m[1])
	fmt.Println(4," ",m[4])
	fmt.Println(3," ",m[3])

	m[10] = "ronaldo"
	m[18] = "virat"
	
	fmt.Println(m)


}


