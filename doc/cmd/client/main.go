package main

import (
	"docs/cmd/client/editor"
)

func main() {

	flags = parseFlags()
	uiConfig := UIConfig{
		EditorConfig: editor.EditorConfig{
			ScrollEnabled: true,
		},
	}
	err := initUI(uiConfig)
}
