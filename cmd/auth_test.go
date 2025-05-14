package main

import (
	"bytes"
	"encoding/json"
	"github.com/jamal23041989/go_short_links/internal/auth"
	"github.com/jamal23041989/go_short_links/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(user.User{
		Email:    "a2@a.ru",
		Password: "$2a$10$a/fN/3apaWOOkO/rEu0Cs.dMcKF7ruPuvfdn3A1mT9.GUDGznVhX2",
		Name:     "Jamal2",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "a2@a.ru").
		Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {
	db := initDb()
	initData(db)

	testServer := httptest.NewServer(App())
	defer testServer.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "1",
	})

	res, err := http.Post(testServer.URL+"/auth", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("expected %d got %d", 200, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}

	if resData.Token == "" {
		t.Fatal("token empty")
	}
	removeData(db)
}

func TestLoginFail(t *testing.T) {
	db := initDb()
	initData(db)

	testServer := httptest.NewServer(App())
	defer testServer.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "12",
	})

	res, err := http.Post(testServer.URL+"/auth", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("expected %d got %d", 401, res.StatusCode)
	}
	removeData(db)
}
