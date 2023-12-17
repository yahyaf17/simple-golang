package handlers

import (
	"encoding/json"
	"native-api-go/db"
	"native-api-go/models"
	"native-api-go/utils"
	"net/http"
	"strconv"
)

func GetMovies(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		validateHttpMethod(res)
		return
	}

	var movies []models.Movie

	for _, m := range db.Movies {
		movies = append(movies, *m)
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

func getMovieById(res http.ResponseWriter, req *http.Request) {
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

	idInt, errAtoi := strconv.Atoi(id[0])
	if errAtoi != nil {
		failedParseData(res)
		return
	}

	var movie, found = db.Movies[idInt]
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

func insertNewMovie(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		validateHttpMethod(res)
		return
	}

	var newMovie models.Movie
	err := json.NewDecoder(req.Body).Decode(&newMovie)

	if err != nil {
		failedParseRequest(res)
		return
	}

	id := utils.GetLargestValue(utils.GetMapKeys(db.Movies)) + 1
	newMovie.ID = id
	db.Movies[id] = &newMovie

	respNew, err := json.Marshal(&newMovie)
	if err != nil {
		failedParseData(res)
		return
	}

	utils.ReturnJsonResponse(res, http.StatusCreated, respNew)
}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPut {
		validateHttpMethod(res)
		return
	}

	var payload models.Movie
	err := json.NewDecoder(req.Body).Decode(&payload)

	if err != nil {
		failedParseRequest(res)
		return
	}

	movie, ok := db.Movies[payload.ID]
	if !ok {
		message, _ := json.Marshal(models.ErrorMsg{
			Success: false,
			Message: "ID not found : " + strconv.Itoa(payload.ID),
		})
		utils.ReturnJsonResponse(res, http.StatusInternalServerError, message)
		return
	}

	movie.Title = utils.SetIfPresent(payload.Title, movie.Title)
	movie.Description = utils.SetIfPresent(payload.Description, movie.Description)
	movie.ReleaseYear = utils.SetIfPresent(payload.ReleaseYear, movie.ReleaseYear)

	respEdit, err := json.Marshal(&movie)
	if err != nil {
		failedParseData(res)
		return
	}

	utils.ReturnJsonResponse(res, http.StatusOK, respEdit)
}

func MovieHandlers(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		insertNewMovie(res, req)
	} else if req.Method == http.MethodGet {
		getMovieById(res, req)
	} else if req.Method == http.MethodPut {
		updateMovie(res, req)
	} else {
		validateHttpMethod(res)
	}
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

func failedParseRequest(res http.ResponseWriter) {
	msg, _ := json.Marshal(models.ErrorMsg{
		Success: false,
		Message: "Failed parse movie data",
	})

	utils.ReturnJsonResponse(res, http.StatusInternalServerError, msg)
}
