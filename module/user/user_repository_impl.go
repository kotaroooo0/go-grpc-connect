package user

import (
	"context"
	"fmt"
	"go-grpc-connect/ent"
	"log"
)

type UserRepositoryImpl struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		client: client,
	}
}

func (ur *UserRepositoryImpl) CreateUser(ctx context.Context, user User) (*User, error) {
	u, err := ur.client.User.
		Create().
		SetAge(user.Age).
		SetName(user.Name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return &User{
		u.ID,
		u.Name,
		u.Age,
	}, nil
}

func (ur *UserRepositoryImpl) GetUser(ctx context.Context, id int) (*User, error) {
	u, err := ur.client.User.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed getting user: %w", err)
	}
	log.Println("user was fetched: ", u)
	return &User{
		u.ID,
		u.Name,
		u.Age,
	}, nil
}
