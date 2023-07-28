package biz

type User struct {
}

type UserRepo interface {
	GetUserByUsername(string) (bool, error)
	CreateUser(*User) error
}

type UserUsecase struct {
	ur UserRepo
}

func NewUsercase(ur UserRepo) *UserUsecase {
	return &UserUsecase{ur: ur}
}
