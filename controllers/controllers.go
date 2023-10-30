package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	moviemodel "github.com/jesseemana/moviesapi/model"
	"github.com/jesseemana/moviesapi/services"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := services.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie moviemodel.Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	services.InsertMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	services.UpdateMovie(params["id"])
	json.NewEncoder(w).Encode("Movie marked as watched")
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	services.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Movie deleted")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := services.DeleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
