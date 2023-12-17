package main

import (
	"fmt"
	"log"
	"native-api-go/db"
	"native-api-go/handlers"
	"native-api-go/models"
	"net/http"
	"os"
)

func main() {
	db.Movies[01] = &models.Movie{ID: 1, Title: "A Space Odyssey", Description: "Science fiction", ReleaseYear: 1998}
	db.Movies[02] = &models.Movie{ID: 2, Title: "Citizen Kane", Description: "Drama", ReleaseYear: 2021}
	db.Movies[3] = &models.Movie{ID: 3, Title: "Raiders of the Lost Ark", Description: "Action and adventure", ReleaseYear: 2021}
	db.Movies[4] = &models.Movie{ID: 4, Title: "66. The General", Description: "Comedy", ReleaseYear: 1978}

	// get movies
	http.HandleFunc("/movies", handlers.GetMovies)
	http.HandleFunc("/movie", handlers.MovieHandlers)

	var port = "3000"

	err := http.ListenAndServe(":"+port, nil)

	log.Println("Movies Simple Application running on port " + port)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
