package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	a   = app.New()
	w   = a.NewWindow("Groupie Tracker")
	api = map[string]string{
		"artists":   "https://groupietrackers.herokuapp.com/api/artists",
		"locations": "https://groupietrackers.herokuapp.com/api/locations",
		"dates":     "https://groupietrackers.herokuapp.com/api/dates",
		"relation":  "https://groupietrackers.herokuapp.com/api/relation",
	}

	artistsData []Artist
)

// Artist represents an artist retrieved from the API
type Artist struct {
	ID           int64               `json:"id"`
	Name         string              `json:"name"`
	Image        string              `json:"image"`
	Members      []string            `json:"members"`
	CreationDate int64               `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Dates        map[string][]string `json:"datesLocations"`
}

func main() {
	apiGet()

	w.Resize(fyne.NewSize(800, 600))
	w.SetMainMenu(createMainMenu())

	showHomePage()

	w.ShowAndRun()
}

func apiGet() {
	for key, url := range api {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to get %s data: %v", key, err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read %s response body: %v", key, err)
		}

		switch key {
		case "artists":
			if err := json.Unmarshal(body, &artistsData); err != nil {
				log.Fatalf("Failed to unmarshal artists data: %v", err)
			}
		default:
			// Handle other API data if needed
		}
	}
}

func createMainMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				widget.NewModalPopUp(
					fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
						widget.NewLabel("Groupie Tracker"),
						widget.NewLabel("Version: 1.0"),
						widget.NewLabel("Developed by: Your Name"),
					),
					w.Canvas(),
				).Show()
			}),
		),
	)
}

func showHomePage() {
	bgImage := canvas.NewImageFromFile("img.jpg")
	bgImage.FillMode = canvas.ImageFillContain

	bgContainer := fyne.NewContainerWithLayout(layout.NewMaxLayout(), bgImage)

	content := container.NewVBox(
		createToolbar(),
		bgContainer,
		container.NewVBox(
			canvas.NewText("Groupie Tracker", theme.PrimaryColor()),
			layout.NewSpacer(),
			createArtistGrid(),
		),
	)

	scrollContainer := container.NewVScroll(content)

	w.SetContent(scrollContainer)
}

func createToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			showHomePage()
		}),
	)
}

func createArtistGrid() *fyne.Container {
	grid := container.NewGridWithColumns(3)
	for _, artist := range artistsData {
		grid.Add(createArtistCard(artist))
	}
	return grid
}

func createArtistCard(artist Artist) fyne.CanvasObject {
	image := canvas.NewImageFromFile(artist.Image)
	image.FillMode = canvas.ImageFillContain

	nameLabel := widget.NewLabel(artist.Name)
	membersLabel := widget.NewLabel(fmt.Sprintf("Members: %s", strings.Join(artist.Members, ", ")))
	firstAlbumLabel := widget.NewLabel(fmt.Sprintf("First Album: %s", artist.FirstAlbum))
	creationDateLabel := widget.NewLabel(fmt.Sprintf("Creation Date: %d", artist.CreationDate))

	return container.NewVBox(
		image,
		nameLabel,
		membersLabel,
		firstAlbumLabel,
		creationDateLabel,
		widget.NewButton("View Details", func() {
			showArtistDetails(artist)
		}),
	)
}

func showArtistDetails(artist Artist) {
	widget.NewModalPopUp(
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			canvas.NewText(artist.Name, theme.PrimaryColor()),
			canvas.NewImageFromFile(artist.Image),
			widget.NewLabel(fmt.Sprintf("Members: %s", strings.Join(artist.Members, ", "))),
			widget.NewLabel(fmt.Sprintf("First Album: %s", artist.FirstAlbum)),
			widget.NewLabel(fmt.Sprintf("Creation Date: %d", artist.CreationDate)),
		),
		w.Canvas(),
	).Show()
}
