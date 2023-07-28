package data

import "go-back/greeter/internal/biz"

type UserRepo struct {
	data *Data
}

func NewUserRepo(data *Data) *UserRepo {
	return &UserRepo{data: data}
}

func (u *UserRepo) GetUserByUsername(username string) (bool, error) {
	return false, nil
}

func (u *UserRepo) CreateUser(user *biz.User) error {
	return nil
}
