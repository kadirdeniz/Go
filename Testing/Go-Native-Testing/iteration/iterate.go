package iterate

func Repeat(val string) string {
	response := ""
	for i := 1; i <= 5; i++ {
		response += val
	}
	return response
}
