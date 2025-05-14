package jwt_test

import (
	"github.com/jamal23041989/go_short_links/pkg/jwt"
	"testing"
)

func TestJWTSign(t *testing.T) {
	const email = "a@a.ru"
	jwtService := jwt.NewJWT("$2a$10$a+fNj3apaWOOkO=rEu0Cs/dMcKF7ruPuvfdn3A1mT9/GUDGznVhX2")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}

	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("token is valid")
	}
	if data.Email != email {
		t.Fatalf("email %s not equal %s", data.Email, email)
	}
}
