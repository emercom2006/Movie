package main

import (
	"awesomeProject/GoArchitecture/gb-go-architecture-master/Movie/Movie/server"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const Addr = ":8081"

func main() {

	//инициализируем подключение к базе данных
	err := server.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/movies", movieListHandler)
	http.HandleFunc("/movies/", movieListHandlerId)
	log.Printf("Starting on port %s", Addr)
	log.Fatal(http.ListenAndServe(Addr, nil))
}

func movieListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	mm, v := server.GetAllFilms()
	err := json.NewEncoder(w).Encode(mm)
	if v != nil {
		log.Printf("Render response error: %v", v)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err != nil {
		log.Printf("Render response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

// Обработчик для отображения содержимого.
func movieListHandlerId(w http.ResponseWriter, r *http.Request) {
	mm, v := server.GetAllFilms()
	if v != nil {
		log.Printf("Render response error: %v", v)
		w.WriteHeader(http.StatusInternalServerError)
	}
	// Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// конвертировать в integer, или значение меньше 0 или больше длины содержимого, возвращаем ответ
	// 404 - страница не найдена!
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	// Используем функцию fmt.Fprintf() для вставки значения из id в строку ответа
	// и записываем его в http.ResponseWriter.
	fmt.Fprintf(w, "%v", mm[id])
}
