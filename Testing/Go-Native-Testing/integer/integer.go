package integer

func Add(num1, num2 int) int {
	return num1 + num2
}
func Minus(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}
