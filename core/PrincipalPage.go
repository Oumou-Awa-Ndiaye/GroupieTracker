package core

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strings"
)

// Global variables for the application and the main window.
var (
	a           = app.New()
	W           = a.NewWindow("Groupie Tracker")
	artistsData []Artist
	artistGrid  *fyne.Container
)

// CreateMainMenu creates the main menu for the application.
func CreateMainMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		//File for a Quit option
		fyne.NewMenu("File",
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				widget.NewModalPopUp(
					fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
						widget.NewLabel("Groupie Tracker  Par Ynov Paris  "),
					),
					W.Canvas(),
				).Show()
			}),
		),
	)
}

func ShowHomePage() {
	// Initialize search bar with a specific size
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search...")

	// Initialize search button
	searchButton := widget.NewButton("Search", func() {
		searchTerm := searchEntry.Text
		results := Searchbar(searchTerm, artistsData)
		updateArtistGrid(results) // Update the artist grid with found artists
	})

	// Container for the search bar without unnecessary spacing
	searchContainer := container.New(layout.NewVBoxLayout(), searchEntry, searchButton)

	// Create the artist grid and set up the home page content
	artistsData = GetArtists()
	artistGrid = createArtistGrid()
	mainContent := container.NewVBox(searchContainer, artistGrid)

	// Create the main navigation toolbar
	toolbar := createToolbar()

	// Set up the main content of the home page
	content := container.NewBorder(toolbar, nil, nil, nil, mainContent)

	// Add a scroll bar if necessary
	scrollContainer := container.NewVScroll(content)
	W.SetContent(scrollContainer)

	// Manage suggestions for the search bar
	searchEntry.OnChanged = func(input string) {
		if input == "" {
			return // Do nothing if the search bar is empty
		}

		// Generate suggestions based on input
		suggestions := generateSuggestions(input, artistsData)

		if len(suggestions) > 0 {
			// Create the suggestions list
			suggestionsList := widget.NewList(
				func() int { return len(suggestions) },
				func() fyne.CanvasObject { return widget.NewLabel("") },
				func(id widget.ListItemID, object fyne.CanvasObject) {
					object.(*widget.Label).SetText(suggestions[id].Name)
				},
			)

			suggestionsList.OnSelected = func(id widget.ListItemID) {
				searchEntry.SetText(suggestions[id].Name)
			}

			// Display the suggestions list in a popup
			showSuggestionsPopup(suggestionsList, searchEntry)
		}
	}

	W.Show()
}

// updateArtistGrid refreshes the artist grid with new data.
func updateArtistGrid(artists []Artist) {
	artistGrid.Objects = nil // Effacer la grille existante
	for _, artist := range artists {
		artistGrid.Add(createArtistCard(artist))
	}
	artistGrid.Refresh() //Refresh the grid with update artistes
}

// createToolbar generates a toolbar for navigation within the app.
func createToolbar() *widget.Toolbar {
	// Returns a toolbar widget with predefined actions, e.g., returning to home page.
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			ShowHomePage()
		}),
	)
}

// createArtistGrid initializes and returns a grid layout containing artist cards.
func createArtistGrid() *fyne.Container {
	artistsData := GetArtists()
	grid := container.NewGridWithColumns(4) // Crée un grille avec 3 colonnes.

	for _, artist := range artistsData {
		artistCard := createArtistCard(artist) // Crée une carte pour chaque artiste.
		grid.Add(artistCard)                   // Ajoute la carte à la grille.
	}

	return grid // Return  a grid layout with artists cards
}

// showArtistDetails displays a popup with detailed information about an artist.
func showArtistDetails(artist Artist) {
	image := getImageFromURL(artist.Image)
	widget.NewPopUp(
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(),

			canvas.NewText(artist.Name, theme.PrimaryColor()),
			widget.NewLabel(fmt.Sprintf("Members: %s", strings.Join(artist.Members, ", "))),
			widget.NewLabel(fmt.Sprintf("First Album: %s", artist.FirstAlbum)),
			widget.NewLabel(fmt.Sprintf("Creation Date: %d", artist.DateCreation)),
			image,
		),
		W.Canvas(),
	).Show()
}

// createArtistCard creates a visual card representing an artist, for use in the grid.
func createArtistCard(artist Artist) fyne.CanvasObject {
	// Create a  label with the artist name
	nameLabel := widget.NewLabelWithStyle(artist.Name, fyne.TextAlignCenter, fyne.TextStyle{})
	nameLabel.Wrapping = fyne.TextTruncate // Pour s'assurer que le texte ne dépasse pas la largeur de la carte.

	image := getImageFromURL(artist.Image)
	image.FillMode = canvas.ImageFillContain
	viewDetailsButton := widget.NewButton(artist.Name, func() {
		showArtistDetails(artist)
	})
	artistCard := container.NewVBox(
		nameLabel,
		image,
		viewDetailsButton,
	)
	artistCard = container.NewPadded(artistCard)
	// Returns a container with artist's name, image,...
	return artistCard
}
