package shared

type Genre string

const (
	Science    Genre = "science"
	History    Genre = "history"
	Philosophy Genre = "philosophy"
	Art        Genre = "art"
	Cooking    Genre = "cooking"
)

var Genres = []Genre{Science, History, Philosophy, Art, Cooking}
