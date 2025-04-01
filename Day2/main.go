package main 

import (
	"fmt"
	"time"
	"learningGoLang/Day2/bookshandle"
)


func welcome(name string,age int){
	fmt.Println("Welcome " + name)
	fmt.Println("i heard you are ", age, " Year Old")
}

func getAge(DOY int) int {
	var age int = time.Now().Year() - DOY 
	return age
}
func last_name_func(name string) (string,int){
	var (
		last_name string
		len_name int
	)
	fmt.Println("what is your last name")
	fmt.Scan(&last_name)

	len_name = len(last_name) + len(name)
	return last_name,len_name
}
func main(){
	fmt.Println("hello world")
	// welcome("Adithyan")
	// var(
	// 	 name string
	// 	 DOY int
	// 	 last_name string
	// 	 len_name int
	// )

	// fmt.Println("what is yout name?")
	// fmt.Scan(&name)
	// fmt.Println("which is your year of brith ?")
	// fmt.Scan(&DOY)
	// var age int = getAge(DOY)

	// welcome(name,age)
	// last_name,len_name = last_name_func(name)

	// fmt.Println("your full name is "+ name + " " + last_name+ " " + "damn your name is ", len_name, "letters long")


	bookshandle.Bookshandle()
}


