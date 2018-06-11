package domain

type Question struct {
	Id          int64
	Answers     []int32
	Favorable   int
	Unfavorable int
	Neltral     int
	Invalid     int
}
