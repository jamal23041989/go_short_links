package main

import (
	"fmt"
	"github.com/jamal23041989/go_short_links/configs"
	"github.com/jamal23041989/go_short_links/internal/auth"
	"github.com/jamal23041989/go_short_links/internal/link"
	"github.com/jamal23041989/go_short_links/pkg/db"
	"github.com/jamal23041989/go_short_links/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	dbNew := db.NewDb(conf)
	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(dbNew)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
