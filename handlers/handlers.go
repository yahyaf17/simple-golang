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
		validateHttpMethod(res)
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
		message, _ := json.Marshal(models.ErrorMsg{
			Success: false,
			Message: "Failed parse movie data",
		})

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, message)
		return
	}

	utils.ReturnJsonResponse(res, http.StatusOK, movieJSON)
}

func GetMovieById(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		validateHttpMethod(res)
		return
	}

	var id, ok = req.URL.Query()["id"]
	if !ok {
		message, _ := json.Marshal(models.ErrorMsg{
			Success: false,
			Message: "Required request ID",
		})
		utils.ReturnJsonResponse(res, http.StatusInternalServerError, message)
		return
	}

	var movie, found = db.Movies[id[0]]
	if !found {
		message, _ := json.Marshal(models.ErrorMsg{
			Success: false,
			Message: "ID not found : " + id[0],
		})
		utils.ReturnJsonResponse(res, http.StatusInternalServerError, message)
		return
	}

	movieJson, err := json.Marshal(&movie)
	if err != nil {
		failedParseData(res)
		return
	}

	utils.ReturnJsonResponse(res, http.StatusOK, movieJson)
}

func validateHttpMethod(res http.ResponseWriter) {
	message, _ := json.Marshal(models.ErrorMsg{
		Success: false,
		Message: "Invalid HTTP Method",
	})

	utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, message)
}

func failedParseData(res http.ResponseWriter) {
	message, _ := json.Marshal(models.ErrorMsg{
		Success: false,
		Message: "Failed parse movie data",
	})

	utils.ReturnJsonResponse(res, http.StatusInternalServerError, message)
}
