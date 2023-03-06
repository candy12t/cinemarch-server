package main

import (
	"fmt"
	"log"

	"github.com/candy12t/cinemarch-server/config"
	"github.com/candy12t/cinemarch-server/infra/database"
	"github.com/candy12t/cinemarch-server/server"
	"github.com/candy12t/cinemarch-server/usecase"
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

	cinemaRepo := database.NewCinemaRepository(db)
	cinemaUC := usecase.NewCinemaUseCase(cinemaRepo)

	r := server.NewRouter(movieUC, cinemaUC)
	addr := fmt.Sprintf(":%s", config.HTTPPort())
	if err := server.NewServer(addr, r).Run(); err != nil {
		log.Fatal(err)
	}
}
