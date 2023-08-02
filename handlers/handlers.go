package handlers

import (
	"encoding/json"
	"native-api-go/native-api-go/db"
	"native-api-go/native-api-go/models"
	"native-api-go/native-api-go/utils"
	"net/http"
)

func GetMovies(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		message := []byte(`{
			"success": false,
			"message": "Invalid HTTP method executed",
		   }`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, message)
		return
	}

	var movies []models.Movie

	for _, m := range db.Movies {
		movies = append(movies, m)
	}

	// parse the movie data into json format
	movieJSON, err := json.Marshal(map[string]interface{}{
		"movieList": movies,
	})
	if err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Error parsing the movie data",
	 }`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	utils.ReturnJsonResponse(res, http.StatusOK, movieJSON)

}
