package user

import (
	"context"
	"go-grpc-connect/ent/enttest"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	userRepo := NewUserRepository(client)

	u := User{
		Name: "Test User",
		Age:  20,
	}

	user, err := userRepo.CreateUser(context.Background(), u)
	require.NoError(t, err)
	assert.Equal(t, u.Name, user.Name)
	assert.Equal(t, u.Age, user.Age)
}

func TestGetUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	client.User.Create().SetName("Test User").SetAge(20).SaveX(context.Background())

	userRepo := NewUserRepository(client)

	user := User{
		Name: "Test User",
		Age:  20,
	}

	result, err := userRepo.GetUser(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, result.Name, user.Name)
	assert.Equal(t, result.Age, user.Age)
}
