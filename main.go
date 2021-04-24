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
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Printf("Render response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
