package model

var Users *[]User

type User struct {
	Id       int
	Name     string
	Surname  string
	Email    string
	Password string
}

func (u *User) Create() int {
	err := append(*Users, *u)
	if err != nil {
		panic(err)
	}
	return 1
}

func (u *User) Delete() {
	for i := 0; i <= len(*Users)-1; i++ {

	}
}

func Get() {

}
