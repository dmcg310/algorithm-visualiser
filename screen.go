package main

import "github.com/gdamore/tcell/v2"

type Screen tcell.Screen

type Grid struct {
	Width, Height int
}

func InitScreen() (Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := screen.Init(); err != nil {
		return nil, err
	}

	return screen, nil
}

func NewGrid(width, height int) *Grid {
	return &Grid{
		Width:  width,
		Height: height,
	}
}
