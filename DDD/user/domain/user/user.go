package user

type User struct {
	first_name string `json="first_name"`
	last_name  string `json="last_name"`
	nickname   string `json="nickname"`
	email      string `json="email"`
	tel        int64  `json="tel"`
}

func New(first_name, last_name, nickname, email string, tel int64) *User {
	return &User{
		first_name: first_name,
		last_name:  last_name,
		nickname:   nickname,
		email:      email,
		tel:        tel,
	}
}
