package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// every fyne app will have at least one windows

	w := a.NewWindow("Hello world")

	// put some content inside the window
	// we will put something called container inside window
	// can also put a widget
	// set the content of my window

	w.SetContent(widget.NewLabel("Hello Wordl"))

	w.ShowAndRun()
}
