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
	sortingAlgorithm SortingAlgorithm
	sortingStarted   bool
	stepCount        int
	readyToSort      bool
}

func NewApp(s Screen, g Grid) *App {
	return &App{
		grid:             g,
		screen:           s,
		isRunning:        true,
		isPaused:         true,
		sortingAlgorithm: NewSortingAlgorithm("Bubble"),
		sortingStarted:   false,
		stepCount:        0,
		readyToSort:      false,
	}
}

func (a *App) Run() {
	eventQueue := make(chan tcell.Event)
	quitQueue := make(chan struct{})

	fps := 45
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
		select {
		case event := <-eventQueue:
			switch event := event.(type) {
			case *tcell.EventKey:
				rune := event.Rune()
				key := event.Key()

				if rune == 'p' {
					if a.sortingStarted {
						a.isPaused = !a.isPaused
					}
				}

				if rune == 's' && !a.sortingStarted {
					a.sortingStarted = true
					a.readyToSort = true
					a.isPaused = false
				}

				if rune == 'r' {
					a.reset()
				}

				if rune == ' ' && a.isPaused && a.sortingStarted {
					a.update()
					a.render()
					a.screen.Show()
				}

				if rune == '1' {
					if a.sortingAlgorithm.Name != "Bubble" {
						a.switchAlgorithm("Bubble")
					}
				}

				if rune == '2' {
					if a.sortingAlgorithm.Name != "Selection" {
						a.switchAlgorithm("Selection")
					}
				}

				if key == tcell.KeyEscape ||
					key == tcell.KeyCtrlC || rune == 'q' {
					close(quitQueue)
					return
				}
			}
		case <-ticker.C:
			if a.sortingStarted && a.readyToSort && !a.isPaused {
				a.update()
			}

			a.render()
			a.screen.Show()
		case <-quitQueue:
			return
		}
	}
}

func (a *App) update() {
	if a.sortingAlgorithm.IsFinished() {
		return
	}

	a.sortingAlgorithm.Step()
	a.stepCount++
}

func (a *App) render() {
	a.screen.Clear()
	width, height := a.screen.Size()

	status := fmt.Sprintf(`Paused: %v | Started: %v | Algorithm: %s | Steps: %d`,
		a.isPaused, a.sortingStarted, a.sortingAlgorithm.Name, a.stepCount)
	if len(status) > width {
		status = status[:width-3] + "..."
	}

	for x := 0; x < width; x++ {
		if x < len(status) {
			a.screen.SetContent(x, 0, rune(status[x]), nil,
				tcell.StyleDefault.Foreground(tcell.ColorSkyblue))
		} else {
			a.screen.SetContent(x, 0, ' ', nil, tcell.StyleDefault)
		}
	}

	arrayStartY := 2           // Status bar offset
	arrayHeight := height - 10 // Keybindings offset
	minCellWidth := 1
	maxCellWidth := 4
	horizontalPadding := 4
	availableWidth := width - horizontalPadding
	cellWidth := maxCellWidth
	visibleElements := len(a.sortingAlgorithm.Array.Values)

	for cellWidth >= minCellWidth && visibleElements*cellWidth > availableWidth {
		if cellWidth == minCellWidth {
			visibleElements = availableWidth / cellWidth
			break
		}
	}

	maxValue := 0
	for i := 0; i < visibleElements && i < len(a.sortingAlgorithm.Array.Values); i++ {
		if a.sortingAlgorithm.Array.Values[i] > maxValue {
			maxValue = a.sortingAlgorithm.Array.Values[i]
		}
	}

	arrayWidth := visibleElements * cellWidth
	arrayStartX := (width - arrayWidth) / 2

	currentIndex, nextIndex := -1, -1
	if a.sortingStarted && !a.sortingAlgorithm.IsFinished() {
		currentIndex, nextIndex =
			a.sortingAlgorithm.Algorithm.GetCurrentIndices()
	}

	for i := 0; i < visibleElements && i < len(a.sortingAlgorithm.Array.Values); i++ {
		value := a.sortingAlgorithm.Array.Values[i]
		xStart := arrayStartX + (i * cellWidth)
		normalizedHeight := int(float64(value) / float64(maxValue) *
			float64(arrayHeight))

		var cellColor tcell.Color
		if i == currentIndex || i == nextIndex {
			cellColor = tcell.ColorRed
		} else {
			cellColor = tcell.ColorWhite
		}

		for dx := 0; dx < cellWidth; dx++ {
			x := xStart + dx

			for y := 0; y < arrayHeight; y++ {
				if y < normalizedHeight {
					a.screen.SetContent(x,
						arrayStartY+arrayHeight-y-1, ' ',
						nil, tcell.StyleDefault.Background(
							cellColor))
				} else {
					a.screen.SetContent(x,
						arrayStartY+arrayHeight-y-1, ' ',
						nil, tcell.StyleDefault.Background(
							tcell.ColorBlack))
				}
			}
		}
	}

	a.renderKeybindings(height)
}

func (a *App) renderKeybindings(height int) {
	width, _ := a.screen.Size()
	keybindings := []string{
		"s: Start sorting",
		"p: Pause/Resume",
		"r: Reset",
		"space: Step (when paused)",
		"q/esc/ctrl-c: Quit",
		"1: Bubble sort, 2: Selection sort",
	}

	for i, binding := range keybindings {
		y := height - len(keybindings) + i
		for x, ch := range binding {
			if x < width {
				a.screen.SetContent(x, y, ch, nil,
					tcell.StyleDefault.Foreground(
						tcell.ColorSkyblue))
			}
		}
	}
}

func (a *App) switchAlgorithm(name string) {
	a.sortingAlgorithm = NewSortingAlgorithm(name)
	a.reset()
}

func (a *App) reset() {
	a.sortingAlgorithm.Reset()
	a.sortingStarted = false
	a.isPaused = true
	a.stepCount = 0
	a.readyToSort = false
}
