package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

func main() {
	a := app.New()
	// every fyne app will have at least one windows

	w := a.NewWindow("Hello world")

	// put some content inside the window
	// we will put something called container inside window
	// can also put a widget
	// set the content of my window
	//	output,entry,btn := makeUI()

	w.SetContent(widget.NewLabel("Hello Wordl"))

	w.ShowAndRun()
}

func makeUI(*widget.Label, *widget.Entry, *widget.Button) {

}
