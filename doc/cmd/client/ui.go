package main

import (
	"docs/cmd/client/editor"

	"github.com/nsf/termbox-go"
)

type UIConfig struct {
	EditorConfig editor.EditorConfig
}

// initUI creates a new editor view and runs the main loop.
func initUI(client *Client, conf UIConfig) error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	e := editor.NewEditor(conf.EditorConfig)
	e.SetSize(termbox.Size())
	e.SetText(client.sayHello())
	e.SendDraw()

	go drawLoop(e)

	return mainLoop(e)
}

// mainLoop is the main update loop for the UI.
func mainLoop(e *editor.Editor) error {
	termboxChan := getTermboxChan()

	for {
		select {
		case ev := <-termboxChan:
			if !handleTermboxEvent(ev, e) {
				return nil // Exit the application
			}
			e.SendDraw()
		}
	}
}

// getTermboxChan returns a channel of termbox Events repeatedly waiting on user input.
func getTermboxChan() chan termbox.Event {
	termboxChan := make(chan termbox.Event)
	go func() {
		for {
			termboxChan <- termbox.PollEvent()
		}
	}()
	return termboxChan
}

// handleTermboxEvent handles key input by updating the editor state.
func handleTermboxEvent(ev termbox.Event, e *editor.Editor) bool {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyEsc, termbox.KeyCtrlC:
			return false // Exit the application
		case termbox.KeyArrowLeft:
			e.MoveCursor(-1, 0)
		case termbox.KeyArrowRight:
			e.MoveCursor(1, 0)
		case termbox.KeyArrowUp:
			e.MoveCursor(0, -1)
		case termbox.KeyArrowDown:
			e.MoveCursor(0, 1)
		case termbox.KeyEnter:
			e.InsertRune('\n')
		case termbox.KeySpace:
			e.InsertRune(' ')
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			e.DeleteRuneBeforeCursor()
		default:
			if ev.Ch != 0 {
				e.InsertRune(ev.Ch)
			}
		}
	case termbox.EventResize:
		e.SetSize(termbox.Size())
	}
	return true
}

func drawLoop(e *editor.Editor) {
	for {
		<-e.DrawChan
		e.Draw()
	}
}
