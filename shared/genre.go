package shared

import "slices"

type Genre string

const (
	Science    Genre = "science"
	History    Genre = "history"
	Philosophy Genre = "philosophy"
	Art        Genre = "art"
	Cooking    Genre = "cooking"
	Fantasy    Genre = "fantasy"
)

var Genres = []Genre{Science, History, Philosophy, Art, Cooking, Fantasy}

func IsValidGenre(genre string) bool {
	return genre != "" && slices.Contains(Genres, Genre(genre))
}

func getValidGenreStr() string {
	var genreStr string
	genreLen := len(Genres)
	for i, genre := range Genres {
		if i < genreLen-1 {
			genreStr += string(genre) + ", "
		} else {
			genreStr += string(genre)
		}
	}

	return genreStr
}

var ValidGenreStr = getValidGenreStr()
