package server

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//глобальная переменная с подключением к БД
var Db *sqlx.DB

//функция, инициирующая подключение к БД
func InitDB() (err error) {
	//строка, содержащая данные для подключения к БД в следующем формате:
	//login:password@tcp(host:port)/dbname
	var dataSourceName = "root:root@tcp(localhost:8889)/movie"
	//подключаемся к БД, используя нужный драйвер и данные для подключения
	Db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return
	}

	err = Db.Ping()
	return
}

type Movie struct {
	ID       int    `json:"id" db:"ID"`
	Name     string `json:"name" db:"Name"`
	Poster   string `json:"poster" db:"Poster"`
	MovieUrl string `json:"movie_url" db:"MovieUrl"`
	IsPaid   bool   `json:"is_paid" db:"IsPaid"`
}

func GetAllFilms() (films []Movie, err error) {
	query := `SELECT * FROM movie.films`
	err = Db.Select(&films, query)
	return
}

//функция, возвращающая число строк в БД -1
func LenDb() int {
	var count int
	err := Db.QueryRow("SELECT COUNT(*) FROM movie.films").Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		count = count - 1
	}
	return count
}
