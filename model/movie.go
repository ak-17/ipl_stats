package model

type Movie struct {
	MovieTitle    string `json:"movie_title"`
	Actor1Name    string `json:"actor_1_name"`
	Actor2Name    string `json:"actor_2_name"`
	Budget        string `json:"budget"`
	DirectorName  string `json:"director_name"`
	Genres        string `json:"genres"`
	MovieImdbLink string `json:"movie_imdb_link"`
	TitleYear     string `json:"title_year"`
}
