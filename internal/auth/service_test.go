package auth_test

import (
	"github.com/jamal23041989/go_short_links/internal/auth"
	"github.com/jamal23041989/go_short_links/internal/user"
	"testing"
)

type MockUserRepository struct {
}

func (m *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.ru",
	}, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "a@a.ru"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1", "Jamal")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("initial email %s not equal email %s", initialEmail, email)
	}
}
