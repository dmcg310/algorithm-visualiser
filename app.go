package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

type App struct {
	grid             Grid
	screen           Screen
	isRunning        bool
	isPaused         bool
	currentAlgorithm string
}

func NewApp(s Screen, g Grid) *App {
	return &App{
		grid:             g,
		screen:           s,
		isRunning:        true,
		isPaused:         true,
		currentAlgorithm: "Bubble",
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
	defer quit(a.screen)

	go func() {
		for {
			event := a.screen.PollEvent()
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
				if !a.isPaused {
					a.update()
				}
			case <-quitQueue:
				return
			}
		}
	}()

	for {
		if a.grid.NeedsRefreshed {
			a.render()
			a.screen.Show()
			a.grid.NeedsRefreshed = false
		}

		select {
		case event := <-eventQueue:
			switch event := event.(type) {
			case *tcell.EventKey:
				rune := event.Rune()
				key := event.Key()

				if rune == 'p' {
					a.isPaused = !a.isPaused
					a.render()
					a.screen.Show()
				}

				if rune == ' ' && a.isPaused {
					a.update()
					a.render()
					a.screen.Show()
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

func (a *App) update() {
}

func (a *App) render() {
	a.screen.Clear()

	width, _ := a.screen.Size()
	status := fmt.Sprintf("Paused: %v | Algorithm: %s", a.isPaused, a.currentAlgorithm)
	if len(status) > width {
		status = status[:width-3] + "..."
	}

	for x := 0; x < width; x++ {
		if x < len(status) {
			a.screen.SetContent(x, 0, rune(status[x]), nil, tcell.StyleDefault.Foreground(tcell.ColorYellow))
		} else {
			a.screen.SetContent(x, 0, ' ', nil, tcell.StyleDefault)
		}
	}

	for x, value := range a.grid.SortArray.Values {
		for y := 0; y < value; y++ {
			a.screen.SetContent(x, a.grid.Height-y-1, ' ', nil,
				tcell.StyleDefault.Background(tcell.ColorWhite))
		}
	}
}
