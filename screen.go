package main

import "github.com/gdamore/tcell/v2"

type Screen tcell.Screen

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
