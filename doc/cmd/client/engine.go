package main

import (
	"docs/cmd/client/editor"
	dpb "docs/pkg/proto/docs"
	"log"

	"github.com/nsf/termbox-go"
)

const (
	OperationInsert = iota
	OperationDelete
)

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
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			performLocalOperation(OperationDelete, ev)
		case termbox.KeyDelete:
			performLocalOperation(OperationDelete, ev)

		// The Tab key inserts 4 spaces to simulate a "tab".
		case termbox.KeyTab:
			for i := 0; i < 4; i++ {
				ev.Ch = ' '
				performLocalOperation(OperationInsert, ev)
			}

		// The Enter key inserts a newline character to the editor's content.
		case termbox.KeyEnter:
			ev.Ch = '\n'
			performLocalOperation(OperationInsert, ev)

		// The Space key inserts a space character to the editor's content.
		case termbox.KeySpace:
			ev.Ch = ' '
			performLocalOperation(OperationInsert, ev)

		// Every other key is eligible to be a candidate for insertion.
		default:
			if ev.Ch != 0 {
				performLocalOperation(OperationInsert, ev)
			}
		}
	case termbox.EventResize:
		e.SetSize(termbox.Size())
	}
	return true
}

// performLocalOperation performs an insert or delete operation on the editor's content.

func performLocalOperation(opType int, ev termbox.Event) {
	ch := string(ev.Ch)
	var msg *dpb.Message
	switch opType {
	case OperationInsert:
		text, err := doc.Insert(e.Cursor+1, ch)
		if err != nil {
			e.SetText(text)
		}
		e.SetText(text)
		e.MoveCursor(1, 0)
		msg = &dpb.Message{
			MessageType: dpb.Message_OPERATION,
			Operation:   &dpb.Operation{OperationType: dpb.Operation_INSERT, Position: int32(e.Cursor), Value: ch},
		}
	case OperationDelete:
		if e.Cursor-1 < 0 {
			e.Cursor = 0
		}
		text := doc.Delete(e.Cursor)
		e.SetText(text)
		msg = &dpb.Message{
			MessageType: dpb.Message_OPERATION,
			Operation:   &dpb.Operation{OperationType: dpb.Operation_DELETE, Position: int32(e.Cursor)},
		}
		e.MoveCursor(-1, 0)
	}
	err := client.sendMessage(msg)
	if err != nil {
		log.Printf("error sendmsg %v", err)
		e.StatusChan <- "lost connection!"
	}

}

func drawLoop(e *editor.Editor) {
	for {
		<-e.DrawChan
		e.Draw()
	}
}
