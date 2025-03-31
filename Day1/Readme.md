# Day one - Learning GO

 
- [X] Install Go & set up workspace (GOPATH, GOROOT)  
- [X] Write your first Go program (Hello, World!)  
- [X] Learn variables, constants, data types, operators  
- [ ] Practice control structures (if, switch, loops)  
- [ ] Hands-on: Write a simple calculator CLI  



### Key Point Learned : 
- in Go, program starts with `package main`.
- you can't import a libary if not use, it will throw error, so basically you can't have all the libaries loaded just in case.
- `func main()` this is the main function, where you code the main thing.

## Contants 
- are fixed values that cannot be reassigned.
- They are created at compile time
- even when defined as locals in functions, and can only be numbers, characters (runes), strings or booleans,  Because of the compile-time restriction.

###### how to use : 
```
const(
    pi = 3.14
)
```

> note* *It is also important to note that, only constants can be assigned to other constants.*



## Variables 

There are two ways for delcaring a variable in Go, 

1.
```go
var (
    first_name string = "Adithyan"
    last_name string = "Madhu"
)
```

2. 
```
var first_name, last_name string = "Adithyan","Madhu"
```




## Data Types :

- Strings 
- Bool
- Numeric Types 
    - Integer
    - Unsigned Integer
    - Float Integer
    - Complex
    - Byte and Rune


### String 
In Go, a string is a sequence of bytes. They are declared either using double quotes or backticks which can span multiple lines.

```go
var name string = "My name is Go"

var bio string = `I am statically typed.
```

### Bool 
yes or no, true or false, on or off, it just a toggle switch which is used to store boolean values.

```go
var value bool = false
var isItTrue bool = true
```

### Numberic Types : 

#### Signed and Unsigned integers :

Signed Integer : 
```go 
var i int = 404                     // Platform dependent
var i8 int8 = 127                   // -128 to 127
var i16 int16 = 32767               // -2^15 to 2^15 - 1
var i32 int32 = -2147483647         // -2^31 to 2^31 - 1
var i64 int64 = 9223372036854775807 // -2^63 to 2^63 - 1
```

Unsinged Integer : 
```go 
var ui uint = 404                     // Platform dependent
var ui8 uint8 = 255                   // 0 to 255
var ui16 uint16 = 65535               // 0 to 2^16
var ui32 uint32 = 2147483647          // 0 to 2^32
var ui64 uint64 = 9223372036854775807 // 0 to 2^64
var uiptr uintptr                     // Integer representation of a memory address
```

#### Floating point:
Next, we have floating point types which are used to store numbers with a decimal component.

Go has two floating point types float32 and float64. Both type follows the IEEE-754 standard.
```go 
var f32 float32 = 1.7812 // IEEE-754 32-bit
var f64 float64 = 3.1415 // IEEE-754 64-bit
```


#### Complex:
There are 2 complex types in Go, complex128 where both real and imaginary parts are float64 and complex64 where real and imaginary parts are float32.

We can define complex numbers either using the built-in complex function or as literals.


```go
var c1 complex128 = complex(10, 1)
var c2 complex64 = 12 + 4i
```


#### Byte and Rune
Golang has two additional integer types called byte and rune that are aliases for uint8 and int32 data types respectively.

```go 
type byte = uint8
type rune = int32
```

A rune represents a unicode code point.

```go 
var b byte = 'a'
var r rune = '🍕'
```


### Operators  :


| Type                | Syntax                                                   |
| ------------------- | -------------------------------------------------------- |
| Arithmetic          | `+` `-` `*` `/` `%`                                      |
| Comparison          | `==` `!=` `<` `>` `<=` `>=`                              |
| Bitwise             | `&` `\|` `^` `<<` `>>`                                   |
| Increment/Decrement | `++` `--`                                                |
| Assignment          | `=` `+=` `-=` `*=` `/=` `%=` `<<=` `>>=` `&=` `\|=` `^=` |
| Logical  | `&&` `\|\|` `!` |
| Equality | `==` `!=`       |




<br>


## Control Structures (if,switch,loops) :
Learning about Flow control.

### If / Else

This works pretty much the same as you expect but the expression doesn't need to be surrounded by parentheses ().


```go 
    x := 20 

	if x>10 {
		fmt.Println("it is greater than 10")
		if x == 20 {
			fmt.Println("got it")
		}else {
			fmt.Println("oops oops, no no ")
		}
	} else {
		fmt.Println("YOOO");
	}
```

### Switch 

In Go, the switch case only runs the first case whose value is equal to the condition expression and not all the cases that follow. Hence, unlike other languages, break statement is automatically added at the end of each case.

This means that it evaluates cases from top to bottom, stopping when a case succeeds. Let's take a look at an example:

``` go
func main() {
	day := "monday"

	switch day {
	case "monday":
		fmt.Println("time to work!")
	case "friday":
		fmt.Println("let's party")
	default:
		fmt.Println("browse memes")
	}
}

```

### Loops 
So in Go, we only have one type of loop which is the for loop.

```go 
for i:= 0 ; i < 10 ; i++ {
    fmt.Println(i);
}
```
Break and continue - same as c++ or any other popular language.


#### for forever loop :
```go 
func main(){
    for {
        // this loops runs forever unless used break to break the loop
    }
}
```



## Making of  basic Calculator : 

Operations : 
- addition 
- substraction 
- multiplication 
- division


#### Features : 
- only two variable 
- maybe have history using array ? (will do later)
- cancel option at any point 


```go
func Cal() {
	for {
		var (
			first  int
			second int
			op     string
			ans    int
		)
		fmt.Println("Enter first number:")
		fmt.Scanln(&first)
		fmt.Println("Enter second number:")
		fmt.Scanln(&second)
		fmt.Println("Enter operator (+,-,*,/):")
		fmt.Scanln(&op)

		switch op {
		case "+":
			ans = first + second
		case "-":
			ans = first - second
		case "*":
			ans = first * second
		case "/":
			ans = first / second
		default:
			fmt.Println("Invalid operator!")
			continue
		}

		fmt.Println("Output:", ans)

		fmt.Println("Do you want to continue (yes/no)?")
		var check string 
		fmt.Scanln(&check)
		if check == "no" {
			fmt.Println("thank you for using this dummy calculator")
			break
		}else {
			fmt.Println("CONTINUE")
			continue
		}

	}
}

```

Lets build 









