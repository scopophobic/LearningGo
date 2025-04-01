# **Day 2: Functions, Structs & Error Handling**

- [X] Learn functions & multiple return values  
- [X] Defers, Stacking Defers and Pointers.
- [X] Understand structs and methods 
- [ ] Learn error handling (errors.New and fmt.Errorf)  
- [ ] Hands-on: Write a CLI tool to manage a list of books (struct-based)  



## Functions and Return values : 

One of unsual features of go is that, we can return multiple values `Multiple return values`.

#### simple declare : 
```go
func myFunction() {}
```

`func <function name> (<parameters>) <return type> {}`
this is the way to make a function in golang.


```go
// normal func usage
func welcome(name string,age int){
	fmt.Println("Welcome " + name)
	fmt.Println("i heard you are ", age, " Year Old")
}

// func with return values
func getAge(DOY int) int {
	var age int = time.Now().Year() - DOY 
	return age
}

// Func with multiple return value
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
```


## Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
func main() {
	defer fmt.Println("world")
	defer fmt.Println("world2")

	fmt.Println("hello")
}
```

### Stacking Defers 
so basically when defer is called, the process goes on to a stack and wait for all other functions to exceute and then one by one from that stack all the defered fucntions are excueted, and this happens in **last in first out** manner since it uses a stack like feature

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```
> Output :
```
$go run main.go

counting
done
9
8
7
6
5
4
3
2
1
0
```




## Pointers
Go has pointers. A pointer holds the memory address of a value.

```go
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer // output = 42
	p = &j         // point to j
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i // output = 21

	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j // output = 73
}
```



## Struct 
A collection of fields.


>Made a struct for student with name(string) and age(int).
```go
type Student struct {
	name string
	age int
}

func main(){
    roll1 := Student{"Adithyan",23}
}
```

## Methods

- Go does not have classes. However, you can define methods on types. 
- A method is a function with a special receiver argument.
- The receiver appears in its own argument list between the func keyword and the method name.

> Methods is bascially providing you with the struct's function to intitialize or for anything
```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```


