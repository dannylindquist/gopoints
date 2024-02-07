package gopoints

type Player struct {
	Name    string  `json:"name"`
	Score   int64   `json:"score"`
	History []int64 `json:"history"`
}
