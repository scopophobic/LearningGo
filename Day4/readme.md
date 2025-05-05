# **Day 4: Pointers, Defer & Interfaces**
- [ ] Learn pointers & memory allocation  
- [ ] Understand defer, panic, recover (error recovery)  
- [ ] Learn interfaces and their importance in Go  
- [ ] Hands-on: Create a logger system using interfaces  



## Pointers and Memory Allocation :
Go has pointers like other high level language, pointers holds memory address of a value. Pointer is mainly used to access the addess of a value and modify the actually value from the memory address.

>The type *T is a pointer to a T value. Its zero value is nil.

`var p *int` is how the pointer is used.


The & operator generates a pointer to its operand. `&` is used to point a value's address and store it in a variable and access and modify the value with `*`
```go
i := 42
p = &i
```

The * operator denotes the pointer's underlying value.

```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```
This is known as "dereferencing" or "indirecting".

>Unlike C, Go has no pointer arithmetic.



