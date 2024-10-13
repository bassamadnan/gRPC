package main

import (
	"docs/cmd/client/editor"
	"docs/crdt"

	"github.com/nsf/termbox-go"
)

type UIConfig struct {
	EditorConfig editor.EditorConfig
}

// TUI is built using termbox-go.
// termbox allows us to set any content to individual cells, and hence, the basic building block of the editor is a "cell".

// initUI creates a new editor view and runs the main loop.
func initUI(conf UIConfig) error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	e = editor.NewEditor(conf.EditorConfig)
	e.SetSize(termbox.Size())
	e.SetText(crdt.Content(doc))
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
