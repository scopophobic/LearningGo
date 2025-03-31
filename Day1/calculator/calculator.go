package calculator

import "fmt"

// Exported function (must start with uppercase `C`)

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
