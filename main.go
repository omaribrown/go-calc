package main

import (
	"fmt"
	"time"
)

// write a struct with 4 methods (add, sub, mul, div)

type calc struct {
	valOne int
	valTwo int
}

func (c calc) add() {

	fmt.Println(c.valOne, " + ", c.valTwo, " = ", c.valOne+c.valTwo)
}

func (c calc) multiply() {
	fmt.Println(c.valOne, " x ", c.valTwo, " = ", c.valOne*c.valTwo)

}
func (c calc) subtract() {
	fmt.Println(c.valOne, " - ", c.valTwo, " = ", c.valOne-c.valTwo)

}
func (c calc) divide() {
	fmt.Println(c.valOne, " / ", c.valTwo, " = ", c.valOne/c.valTwo)

}

func main() {
	var tempParamOne int
	var tempParamTwo int
	var tempOperator string

	// have the user input three values, val1, operator, val2,

	fmt.Println("Enter your first value: ")
	fmt.Scanln(&tempParamOne)

	fmt.Println("Enter your operator ('add', 'sub', 'mul', 'div')")
	fmt.Scanln(&tempOperator)

	fmt.Println("Enter your second value: ")
	fmt.Scanln(&tempParamTwo)

	params := calc{
		tempParamOne,
		tempParamTwo,
	}

	// depending on the operatur, run the specified equation, return the result

	switch tempOperator {
	case "add":
		fmt.Println("Addition: ")
		params.add()

	case "sub":
		fmt.Println("Subtraction: ")
		params.subtract()

	case "mul":
		fmt.Println("Multiplication: ")
		params.multiply()

	case "div":
		fmt.Println("Division: ")
		params.divide()
	}

	// return the time of the calculation
	currentTime := time.Now()
	fmt.Println("Calculation run at: ", currentTime.Format("2006-01-02 3:4:5 PM"))

}
