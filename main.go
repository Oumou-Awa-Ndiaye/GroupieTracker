package main

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
=======
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
>>>>>>> 6cbcf53ca8a7cee16bf07684eba591aa4818555f
)

type Bands struct {
	Index []Artist `json:"index"`
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func main() {

<<<<<<< HEAD
	searchEntry := CreateSearchBar()

	w.SetContent(container.NewVBox(searchEntry, widget.NewLabel("MUUUUSICCCC!")))
	w.ShowAndRun()
=======
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
>>>>>>> 6cbcf53ca8a7cee16bf07684eba591aa4818555f

	var bands Bands
	err := json.Unmarshal([]byte(sampleData), &bands)
	if err != nil {
		log.Fatal(err)
	}

	minCreation := 2000
	maxCreation := 2010
	minMembers := 3
	maxMembers := 5
	concertLocations := []string{"New York", "Los Angeles"}

	filteredByCreationDate := filterByCreationDate(bands, minCreation, maxCreation)
	fmt.Println("Groupes filtrés par date de création :", filteredByCreationDate)

	filteredByMembersCount := filterByMembersCount(bands, minMembers, maxMembers)
	fmt.Println("Groupes filtrés par nombre de membres :", filteredByMembersCount)

	filteredByConcertLocations := filterByConcertLocations(bands, concertLocations)
	fmt.Println("Groupes filtrés par lieux de concerts :", filteredByConcertLocations)

	w.SetContent(widget.NewLabel("MUUUUSICCCC!"))
	w.ShowAndRun()
}

<<<<<<< HEAD
func filterByCreationDate(bands Bands, minCreationDate, maxCreationDate int) []Artist {
	var filteredBands []Artist
	for _, band := range bands.Index {
		if band.CreationDate >= minCreationDate && band.CreationDate <= maxCreationDate {
			filteredBands = append(filteredBands, band)
		}
	}
	return filteredBands
}

func filterByMembersCount(bands Bands, minMembers, maxMembers int) []Artist {
	var filteredBands []Artist
	for _, band := range bands.Index {
		numMembers := len(band.Members)
		if numMembers >= minMembers && numMembers <= maxMembers {
			filteredBands = append(filteredBands, band)
		}
	}
	return filteredBands
}

func filterByConcertLocations(bands Bands, locations []string) []Artist {
	var filteredBands []Artist
	for _, band := range bands.Index {
		for _, loc := range locations {
			if strings.Contains(band.Locations, loc) {
				filteredBands = append(filteredBands, band)
				break
			}
		}
	}
	return filteredBands
}

func CreateSearchBar() *widget.Entry {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Recherche...")
	return searchEntry
}

// Données JSON d'exemple

const sampleData = `
{
	"index": [
		{
			"id": 1,
			"image": "image1.jpg",
			"name": "Groupe 1",
			"members": ["Membre 1", "Membre 2"],
			"creationDate": 2000,
			"firstAlbum": "Album 1",
			"locations": "New York, Londres",
			"concertDates": "2024-02-10, 2024-02-15",
			"relations": "Relation 1"
		},
		{
			"id": 2,
			"image": "image2.jpg",
			"name": "Groupe 2",
			"members": ["Membre 3", "Membre 4", "Membre 5"],
			"creationDate": 2010,
			"firstAlbum": "Album 2",
			"locations": "Los Angeles",
			"concertDates": "2024-02-12, 2024-02-17",
			"relations": "Relation 2"
		}
	]
}
`
=======
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
>>>>>>> 6cbcf53ca8a7cee16bf07684eba591aa4818555f
