package repository

import (
	"github.com/user/domain/user"
	"github.com/user/domain/user/repository"
)

const insertUser string = `INSERT INTO user (first_name, last_name, nickname, email, tel) VALUES (?,?,?,?,?)`

type repo struct {
	db repository.DB
}

func (r *repo) Create(user user.User) user.User {
	_, err := r.db.Exec(insertUser)
}
