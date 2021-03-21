package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const Addr = ":8081"

func main() {
	http.HandleFunc("/movies", movieListHandler)
	http.HandleFunc("/movies/", movieListHandlerId)
	log.Printf("Starting on port %s", Addr)
	log.Fatal(http.ListenAndServe(Addr, nil))
}

type Movie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Poster   string `json:"poster"`
	MovieUrl string `json:"movie_url"`
	IsPaid   bool   `json:"is_paid"`
}

func movieListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	mm := []Movie{
		Movie{0, "Бойцовский клуб", "/static/posters/fightclub.jpg", "https://youtu.be/qtRKdVHc-cE", true},
		Movie{1, "Крестный отец", "/static/posters/father.jpg", "https://youtu.be/ar1SHxgeZUc", false},
		Movie{2, "Криминальное чтиво", "/static/posters/pulpfiction.jpg", "https://youtu.be/s7EdQ4FqbhY", true},
	}
	err := json.NewEncoder(w).Encode(mm)
	if err != nil {
		log.Printf("Render response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

// Обработчик для отображения содержимого заметки.
func movieListHandlerId(w http.ResponseWriter, r *http.Request) {
	mm := []Movie{
		Movie{0, "Бойцовский клуб", "/static/posters/fightclub.jpg", "https://youtu.be/qtRKdVHc-cE", true},
		Movie{1, "Крестный отец", "/static/posters/father.jpg", "https://youtu.be/ar1SHxgeZUc", false},
		Movie{2, "Криминальное чтиво", "/static/posters/pulpfiction.jpg", "https://youtu.be/s7EdQ4FqbhY", true},
	}
	// Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// конвертировать в integer, или значение меньше 0 или больше длины содержимого, возвращаем ответ
	// 404 - страница не найдена!
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 && id >= len(mm){
		http.NotFound(w, r)
		return
	}

	// Используем функцию fmt.Fprintf() для вставки значения из id в строку ответа
	// и записываем его в http.ResponseWriter.
	fmt.Fprintf(w, "%v",mm[id])
}