package main

import (
	"fmt"
	"github.com/jamal23041989/go_short_links/configs"
	"github.com/jamal23041989/go_short_links/internal/auth"
	"github.com/jamal23041989/go_short_links/internal/link"
	"github.com/jamal23041989/go_short_links/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
