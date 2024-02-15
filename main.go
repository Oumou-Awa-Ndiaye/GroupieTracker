package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	app := app.New()
	window := app.NewWindow("MUUUSIICCC")
	window.Resize(fyne.NewSize(700, 400))
	window.SetContent(widget.NewLabel("MUUUUSICCCC!"))

	// Create a text entry widget for the search bar
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search...")

	// Create a label to display the search results
	searchResults := widget.NewLabel("")

	// Event handler for when the user types in the search bar
	searchEntry.OnChanged = func(query string) {
		// Here you can perform search operations based on the query
		// For this example, we'll just display the query as the search result
		searchResults.SetText("Search results for: " + query)
	}

	artButton := widget.NewButton("Artists", func() {
		println("Artists button clicked")
	})

	datButton := widget.NewButton("Dates", func() {
		println("Dates button clicked")
	})

	locButton := widget.NewButton("Locations", func() {
		println("Locations button clicked")
	})

	window.SetContent(container.NewVBox(
		widget.NewLabel("MUUUUSICCCC!"),
		searchEntry,
		searchResults,
		artButton,
		datButton,
		locButton,
	))

	window.ShowAndRun()

}
