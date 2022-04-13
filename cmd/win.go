package main

import (
	"github.com/tadvi/winc"
)

func StartWindow() *winc.Form {
	window := winc.NewForm(nil)
	window.SetMaxSize(1440, 900)
	window.SetSize(800, 400)

	FileButton(window)
	EditButton(window)
	FormatButton(window)
	ViewButton(window)
	HelpButton(window)
	ApiButton(window)

	//NewEdit() text form
	NewEdit(window)

	//params
	window.Center()
	window.Show()
	window.OnClose().Bind(wndOnClose)
	winc.RunMainLoop() // must call to start event loop
	return window
}

func Start() {
	StartWindow()
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}
