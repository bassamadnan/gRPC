package utils

import (
	"docs/cmd/client/editor"
	"time"

	"github.com/nsf/termbox-go"
)

type UIConfig struct {
	EditorConfig editor.EditorConfig
}

// TUI is built using termbox-go.
// termbox allows us to set any content to individual cells, and hence, the basic building block of the editor is a "cell".

// initUI creates a new editor view and runs the main loop.
func (c *Client) InitUI(conf UIConfig) error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	e := editor.NewEditor(conf.EditorConfig)
	e.SetSize(termbox.Size())
	e.SetText(c.sayHello())
	e.SendDraw()
	e.IsConnected = true
	time.Sleep(1 * time.Second)
	return nil
}
