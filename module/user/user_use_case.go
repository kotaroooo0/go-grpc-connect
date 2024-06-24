package user

import "context"

type UserRepository interface {
	CreateUser(context.Context, User) (*User, error)
	GetUser(context.Context, int) (*User, error)
}

type UserUseCase struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, name string, age int) (*User, error) {
	u := User{
		Name: name,
		Age:  age,
	}
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}
