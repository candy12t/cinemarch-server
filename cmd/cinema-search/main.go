package main

import (
	"fmt"
	"log"

	"github.com/candy12t/cinema-search-server/config"
	"github.com/candy12t/cinema-search-server/infra/database"
	"github.com/candy12t/cinema-search-server/server"
	"github.com/candy12t/cinema-search-server/usecase"
)

func main() {
	db, cleanup, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := cleanup(); err != nil {
			log.Fatal(err)
		}
	}()

	movieRepo := database.NewMovieRepository(db)
	movieUC := usecase.NewMovieUseCase(movieRepo)

	r := server.NewRouter(movieUC)
	addr := fmt.Sprintf(":%s", config.HTTPPort())
	if err := server.NewServer(addr, r).Run(); err != nil {
		log.Fatal(err)
	}
}
