package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	PerformGetRequest()

	/*app := app.New()
	window := app.NewWindow("Groupie Tracker App")
	window.Resize(fyne.NewSize(700, 400))
	window.SetContent(widget.NewLabel("Welcome to Groupie Tracker !"))

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for your favorit artists ...")

	searchResults := widget.NewLabel("")

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
		widget.NewLabel("Welcome to Groupie Tracker !"),
		searchEntry,
		searchResults,
		artButton,
		datButton,
		locButton,
	))

	window.ShowAndRun()*/

}

func PerformGetRequest() {
	const myurl = "https://groupietrackers.herokuapp.com/api/artists"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount is : ", byteCount)
	fmt.Println(responseString.String())

	//fmt.Println(content)
	//fmt.Println(string(content))
}
