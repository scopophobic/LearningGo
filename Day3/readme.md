# **Day 3: Arrays, Slices, Maps & Ranges**
- [X]Learn arrays, slices, map types & receiver functions  
- [X]Understand slice manipulation & memory efficiency  
- [ ]Learn how range works with collections  
- [ ]Hands-on: Write a program that processes student records (map-based)  


## Arrays 
Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation, but primarily they are a building block for slices.
```go 
var a [10]int
```

- Arrays are values. Assigning one array to another copies all the elements.
- In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
- The size of an array is part of its type. The types [10]int and [20]int are distinct.



## Slices : 
Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data.

Except for items with explicit dimension such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.


```go
	var arr [10]int = [10]int{1,2,3,4,5,6,7,8,9,10} // main array
	a := arr[1:5] //sliced from arr
	fmt.Println(a)
	fmt.Println(arr)
```



How different slices work :

```go
var a = [4]string{
		"C++",
		"Go",
		"Java",
		"TypeScript",
	}

	s1 := a[0:2] // Select from 0 to 2
	s2 := a[:3]  // Select first 3
	s3 := a[2:]  // Select last 2

	fmt.Println("Array:", a)
	fmt.Println("Slice 1:", s1)
	fmt.Println("Slice 2:", s2)
	fmt.Println("Slice 3:", s3)
```
>output : 
```output
$ go run main.go
Array: [C++ Go Java TypeScript]
Slice 1: [C++ Go]
Slice 2: [C++ Go Java]
Slice 3: [Java TypeScript]
```


## Maps 
A map is an unordered collection of key-value pairs. 
It maps keys to values. The keys are unique within a map while
 the values may not be.

#### Use case of Map
- It is used for fast lookups
- retrieval of data with keys
- deletion of data based on keys


declaring a map - `var m map[K]V`,  
K is the key's data type 
and v is the value's data type.

i can initialise values like this  
```go
var m = map[string]string{
		"a": "Hello",
		"b": "hi",
	}
```

Retriving Data :
```go
c := m["c"]
fmt.Println("Key c:", c)
```




## Why do we need slices?

- Helps in memory efficientcy.
- Provide more power, flexibility, and convenience.


#### A slice consists of three things:

- A pointer reference to an underlying array.
- The length of the segment of the array that the slice contains.
- And, the capacity, which is the maximum size up to which the segment can grow.



```go 
package main

import "fmt"

func main() {
	a := [5]int{20, 15, 5, 30, 25}

	s := a[1:4]

	// Output: Array: [20 15 5 30 25], Length: 5, Capacity: 5
	fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", a, len(a), cap(a))

	// Output: Slice [15 5 30], Length: 3, Capacity: 4
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d", s, len(s), cap(s))
}
```

This is how slices work, where it is dynamically sliced and used as per the usage, helps you be more memory efficient and compact.


> So, unlike arrays, the zero value of a slice is nil.

### Manupilation on Slices : 

#### 1. Declaration 
Let's see how we can declare a slice.

`var s []T`
As we can see, we don't need to specify any length. Let's declare a slice of integers and see how it works.
```go
func main() {
	var s []string

	fmt.Println(s)
	fmt.Println(s == nil)
}
```
```bash
$ go run main.go
[]
true
```

#### Initialization
There are multiple ways to initialize our slice. One way is to use the built-in make function.
```go
make([]T, len, cap) []T
func main() {
	var s = make([]string, 0, 0)

	fmt.Println(s)
}
```

```bash
$ go run main.go
[]
```
Similar to arrays, we can use the slice literal to initialize our slice.

```go
func main() {
	var a = [4]string{
		"C++",
		"Go",
		"Java",
		"TypeScript",
	}

	s1 := a[0:2] // Select from 0 to 2
	s2 := a[:3]  // Select first 3
	s3 := a[2:]  // Select last 2

	fmt.Println("Array:", a)
	fmt.Println("Slice 1:", s1)
	fmt.Println("Slice 2:", s2)
	fmt.Println("Slice 3:", s3)
}
```
```bash
$ go run main.go
Array: [C++ Go Java TypeScript]
Slice 1: [C++ Go]
Slice 2: [C++ Go Java]
Slice 3: [Java TypeScript]```


