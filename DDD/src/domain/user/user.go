package user

type user struct {
	first_name string
	last_name  string
	email      string
	phone      string
	age        int
	gender     string
}

type User interface {
	New(firstName, lastName, email, phone, gender string, age int) user
	IsAdult() bool
	IsChild() bool
	IsBaby() bool
	IsTeenager() bool
}

func (u *user) New(firstName, lastName, email, phone, gender string, age int) user {
	return user{
		first_name: firstName,
		last_name:  lastName,
		email:      email,
		phone:      phone,
		age:        age,
		gender:     gender,
	}
}

func (u *user) IsAdult() bool {
	if u.Age >= 18 {
		return true
	}
	return false
}

func (u *user) IsChild() bool {
	if u.Age < 13 && u.Age > 5 {
		return true
	}
	return false
}

func (u *user) IsBaby() bool {
	if u.Age >= 0 && u.Age < 3 {
		return true
	}
	return false
}

func (u *user) IsTeenager() bool {
	if u.Age > 13 && u.Age < 18 {
		return true
	}
	return false
}
