# **Day 3: Arrays, Slices, Maps & Ranges**
- [X]Learn arrays, slices, map types & receiver functions  
- [ ]Understand slice manipulation & memory efficiency  
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




