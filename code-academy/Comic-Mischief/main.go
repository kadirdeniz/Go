package main

import (
	"fmt"
)

func main() {
	var title,
		writer,
		artist,
		publisher string = "Tracey Hatchet", "Mr. GoToSleep", "Jewel Tampson", "DizzyBooks Publishing Inc."

	var year, pageNumber int = 1997, 14
	var grade float32 = 6.5

	fmt.Println(publisher, writer, artist, title)
	fmt.Println(year, pageNumber)
	fmt.Println(grade)

	title, writer, artist, publisher = "Epic Vol. 1", "Ryan N. Shawn", "Phoebe Paperclips", "DizzyBooks Publishing Inc."

	year, pageNumber = 2013, 160
	grade = 9.0

	fmt.Println(publisher, writer, artist, title)
	fmt.Println(year, pageNumber)
	fmt.Println(grade)

}
