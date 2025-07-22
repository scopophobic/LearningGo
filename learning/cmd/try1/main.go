

package main

import (
	"fmt"
)


func main(){
	var intNum int
	fmt.Println("i'm learning to code here in go ehhe")
	// int8 int16 int32 int64
	// uint8 uint16 uint32 uint64 

	fmt.Println("Hello World")
	intNum = 10;
	printme(intNum)


}


func printme(a any) {
	fmt.Println(a)
}
