package domain

type Question struct {
	Id      int64
	Answers []int32
	// NumFavorite   int
	// NumNeutral    int
	// NumUnFavorite int
	// Invalid       int
}
