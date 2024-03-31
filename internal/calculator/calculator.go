package calculator

func Calculate(x, y int, operator string) int {
	switch operator {
	case "+":
		return add(x, y)
	case "-":
		return minus(x, y)
	case "*":
		return multyply(x, y)
	case "/":
		return divide(x, y)
	default:
		panic("Invalid operator")
	}
}

func add(x, y int) int {
	return x + y
}

func minus(x, y int) int {
	return x - y
}

func multyply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	if y == 0 {
		panic("Division by zero")
	}

	return x / y
}
