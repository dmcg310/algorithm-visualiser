package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type App struct {
	Screen    Screen
	IsRunning bool
	IsPaused  bool
}

func NewApp(s Screen) *App {
	return &App{
		Screen:    s,
		IsRunning: true,
		IsPaused:  true,
	}
}

func (a *App) Run() {
	eventQueue := make(chan tcell.Event)
	quitQueue := make(chan struct{})

	fps := 5
	ticker := time.NewTicker(time.Second / time.Duration(fps))
	defer ticker.Stop()

	quit := func(screen Screen) {
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit(a.Screen)

	go func() {
		for {
			event := a.Screen.PollEvent()
			if event == nil {
				return
			}

			eventQueue <- event
		}
	}()

	go func() {
		for {
			select {
			case <-ticker.C:
				if !a.IsPaused {
					// Run the app
				}
			case <-quitQueue:
				return
			}
		}
	}()

	for {
		select {
		case event := <-eventQueue:
			switch event := event.(type) {
			case *tcell.EventKey:
				rune := event.Rune()
				key := event.Key()

				if rune == 'p' {
					a.IsPaused = !a.IsPaused
					a.Screen.Show()
				}

				if rune == ' ' && a.IsPaused {
					a.Screen.Show()
				}

				if key == tcell.KeyEscape ||
					key == tcell.KeyCtrlC || rune == 'q' {
					close(quitQueue)
					return
				}
			}
		case <-quitQueue:
			break
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}
