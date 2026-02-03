package model

type Desk struct {
	desk [][]rune
	move
	Game
}

type Player struct {
	Name   string
	Rating int
}
type deskSize struct {
	column int
	line   int
}

type Game struct {
	deskSize
	Player
}
type move struct {
	source string
	dest   string
	user   Player
}
