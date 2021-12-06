package main

func addition(x, y int) int {
	return x + y
}

func substraction(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func multiply(x, y int) int {
	return x * y
}
