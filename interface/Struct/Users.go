package user

type User struct {
	Id      int
	Name    string
	Surname string
	Email   string
	Age     int
}

var Users = []User{
	{Id: 1, Name: "Kadir", Surname: "Deniz", Email: "kadirdenz@hotmail.com", Age: 22},
	{Id: 2, Name: "Yunus Emre", Surname: "Deniz", Email: "yunusemre@hotmail.com", Age: 18},
	{Id: 3, Name: "Esin", Surname: "Deniz", Email: "esindeniz@hotmail.com", Age: 46},
	{Id: 4, Name: "Abdullah", Surname: "Deniz", Email: "abdullahdeniz@hotmail.com", Age: 46},
}
