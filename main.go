package main

import (
	"awesomeProject/Movie/gateway-user/server"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

var Addr = ":8092"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", movieListHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", movieGetHandler).Methods("GET")

	log.Printf("Starting on port %s", Addr)
	log.Fatal(http.ListenAndServe(Addr, r))
}

type Movie struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Poster      string `json:"poster"`
	MovieUrl    string `json:"movie_url"`
	IsPaid      bool   `json:"is_paid"`
}

func movieGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	vars := mux.Vars(r)
	movieID := vars["id"]

	mm := []Movie{
		Movie{0, "Бойцовский клуб", "Description", "/static/posters/fightclub.jpg", "https://youtu.be/qtRKdVHc-cE", true},
		Movie{1, "Крестный отец", "Description", "/static/posters/father.jpg", "https://youtu.be/ar1SHxgeZUc", false},
		Movie{2, "Криминальное чтиво", "Description", "/static/posters/pulpfiction.jpg", "https://youtu.be/s7EdQ4FqbhY", true},
	}

	var resultMovie Movie
	for _, movie := range mm {
		if strconv.Itoa(movie.ID) == movieID {
			resultMovie = movie
		}
	}
	if resultMovie.ID == 0 {
		log.Printf("Render response error: %v", errors.New("movie not found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(&resultMovie)
	if err != nil {
		log.Printf("Render response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func movieListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//mm := []Movie{
	//	Movie{0, "Бойцовский клуб", "Description", "/static/posters/fightclub.jpg", "https://youtu.be/qtRKdVHc-cE", true},
	//	Movie{1, "Крестный отец", "Description", "/static/posters/father.jpg", "https://youtu.be/ar1SHxgeZUc", false},
	//	Movie{2, "Криминальное чтиво", "Description", "/static/posters/pulpfiction.jpg", "https://youtu.be/s7EdQ4FqbhY", true},
	//}
	var Db *sqlx.DB

	err := server.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := Db.Query(`SELECT * FROM movie.films`)
	//err = Db.Select(&films, res)

	posts := make([]*Movie, 0)

	for res.Next() {

		bk := new(Movie)
		err := res.Scan(&bk.ID, &bk.Name, &bk.Description, &bk.Poster, &bk.MovieUrl, &bk.IsPaid)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, bk)

		if err != nil {
			log.Fatal(err)
		}
	}
	jsonData, err := json.Marshal(&posts)
	err = json.NewEncoder(w).Encode(jsonData)
	if err != nil {
		log.Printf("Render response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
