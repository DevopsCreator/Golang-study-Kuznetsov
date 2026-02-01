package model

import "time"



type Player struct{
	Name string
	Rating int
}

type Game struct{
	Users Player
	moveCount move
}

type move struct {
	count int
	user Player
}

type Field struct {
	NumberОfFigures int
	Game
}