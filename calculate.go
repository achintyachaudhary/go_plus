package main

func main() {
	println("Division", calculate(10000, 5, DIVIDE))
	println("Multiply", calculate(10000, 5, MULTIPLY))
	println("Subtract", calculate(10000, 5, SUB))
	println("Addition", calculate(10000, 5, ADD))
}

const MULTIPLY = "multiply"
const DIVIDE = "divide"
const ADD = "add"
const SUB = "subtract"

func calculate(x float64, y float64, operands string) float64 {
	if operands == MULTIPLY {
		return x * y
	} else if operands == DIVIDE {
		return x / y
	} else if operands == ADD {
		return x + y
	} else if operands == SUB {
		return x - y
	}
	// Weird as must return at end is required
	return 0
}
