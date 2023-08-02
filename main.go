package main

import (
	"fmt"
	"native-api-go/native-api-go/db"
	"native-api-go/native-api-go/handlers"
	"native-api-go/native-api-go/models"
	"net/http"
	"os"
)

func main() {
	db.Movies["001"] = models.Movie{ID: "001", Title: "A Space Odyssey", Description: "Science fiction", ReleaseYear: 1998}
	db.Movies["002"] = models.Movie{ID: "002", Title: "Citizen Kane", Description: "Drama", ReleaseYear: 2021}
	db.Movies["003"] = models.Movie{ID: "003", Title: "Raiders of the Lost Ark", Description: "Action and adventure", ReleaseYear: 2021}
	db.Movies["004"] = models.Movie{ID: "004", Title: "66. The General", Description: "Comedy", ReleaseYear: 1978}

	// get movies
	http.HandleFunc("/movies", handlers.GetMovies)

	err := http.ListenAndServe(":3000", nil)
	// print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
