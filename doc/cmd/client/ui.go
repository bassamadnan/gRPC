package main

import (
	"docs/cmd/client/editor"
	"docs/crdt"
	"log"

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
	// client.sendError()
	e.SetText(crdt.Content(doc))
	e.SendDraw()

	go drawLoop()

	return mainLoop()
}

// mainLoop is the main update loop for the UI.
func mainLoop() error {
	termboxChan := getTermboxChan()

	for {
		select {
		case ev := <-termboxChan:
			if !handleTermboxEvent(ev) {
				return nil // exit the application
			}
		case msg, ok := <-msgChan:
			if !ok {
				log.Fatalf("not ok server closed?")
				return nil
			}
			handleServerMessage(msg)

		}
	}
}
