/*package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("MUUUSIICCC")
	w.Resize(fyne.NewSize(400, 400))
	w.SetContent(widget.NewLabel("MUUUUSICCCC!"))

	artButton := widget.Newbutton("Artists", func() {

	})

	w.SetContent(artButton)

	w.ShowAndRun()
}*/

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("MUUUSIICCC")
	w.Resize(fyne.NewSize(400, 400))
	w.SetContent(widget.NewLabel("MUUUUSICCCC!"))

	artButton := widget.NewButton("Artists", func() {
		println("Artists button clicked")
	})

	datButton := widget.NewButton("Dates", func() {
		println("Dates button clicked")
	})

	locButton := widget.NewButton("Locations", func() {
		println("Locations button clicked")
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("MUUUUSICCCC!"),
		artButton,
		datButton,
		locButton,
	))

	w.ShowAndRun()

}
