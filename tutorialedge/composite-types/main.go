package main

import "fmt"

func main() {
	// declaring an empty array of strings
	// var days []string

	// days[0] = "Monday"
	// days[1] = "Tuesday"
	// days[2] = "Wednesday"
	// days[3] = "Thursday"
	// days[4] = "Wednesday"

	// declaring an array with elements
	var days [7]string = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friyday", "Saturday", "Sunday"}
	days2 := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friyday", "Saturday", "Sunday"}
	days3 := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friyday", "Saturday", "Sunday"}
	fmt.Println(days)
	fmt.Println(days2)
	fmt.Println(days3)

	// Slices
	weekdays := days[0:5]
	fmt.Println(weekdays)

	//Maps
	youtubeSubscribers := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fun Fun Function": 171220,
	}
	fmt.Println(youtubeSubscribers["MKBHD"])

	// Struct
	type Person struct {
		name string
		age  int
	}

	type Team struct {
		name    string
		players [2]Person
	}

	var me Person = Person{
		name: "Kadir",
		age:  22,
	}
	var yunus Person = Person{
		name: "Yunus Emre",
		age:  18,
	}

	var teamPlayers [2]Person = [2]Person{
		me, yunus,
	}
	var fenerbahce Team = Team{
		name:    "Fenerbahçe",
		players: teamPlayers,
	}

	fmt.Println(fenerbahce)
	fmt.Println(fenerbahce.players[0].name)

}
