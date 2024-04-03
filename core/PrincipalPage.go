package core

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Global variables for the application and the main window.
var (
	A                    = app.New()
	W                    = A.NewWindow("Groupie Tracker")
	artistsData          []Artist
	artistGrid           *fyne.Container
	suggestionsContainer *widget.PopUp //conteneur pour les suggestions qui sera ajouté sous la barre de recherche
)

// CreateMainMenu creates the main menu for the application.
func CreateMainMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		//File for a Quit option
		fyne.NewMenu("File",
			fyne.NewMenuItem("Quit", func() {
				A.Quit()
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

func ShowHomePage(A fyne.App) {
	// Initialize search bar with a specific size
	W := A.NewWindow("Groupie Tracker")
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Chercher vos artistes ...")

	// Initialize search button
	searchButton := widget.NewButton("Search", func() {
		searchTerm := searchEntry.Text
		results := Searchbar(searchTerm, artistsData)
		updateArtistGrid(results) // Update the artist grid with found artists
	})

	// Initialize filter button

	filterButton := widget.NewButton("Recherche avec Filtre", func() {
		FilterPage(A)
		W.Hide()
	})

	// Container for the search bar without unnecessary spacing
	searchContainer := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		container.New(layout.NewVBoxLayout(),
			searchEntry,
			layout.NewSpacer(), // Add some space between search entry and buttons
			container.New(layout.NewHBoxLayout(),
				layout.NewSpacer(), // Add some space before search button
				searchButton,
				layout.NewSpacer(), // Add some space between search and filter buttons
				filterButton,
				layout.NewSpacer(), // Add some space after filter button
			),
			layout.NewSpacer(), // Add some space at the bottom of the container
		),
		layout.NewSpacer(),
	)
	suggestionsBox := container.NewVBox()
	// Fonction pour montrer ou cacher les suggestions
	showSuggestions := func(show bool) {
		if suggestionsContainer != nil {
			suggestionsContainer.Hide()
			if show {
				suggestionsContainer.Show()
			}
		}
	}
	// Create the artist grid and set up the home page content
	artistsData = GetArtists()
	artistGrid = createArtistGrid()
	mainContent := container.NewVBox(searchContainer, artistGrid)

	// Create the main navigation toolbar
	homeButton := widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		ShowHomePage(A)
		W.Hide()
	})

	// Set up the main content of the home page
	content := container.NewBorder(container.NewHBox(homeButton, layout.NewSpacer()), nil, nil, nil, mainContent)

	// Add a scroll bar if necessary
	scrollContainer := container.NewVScroll(content)
	W.SetContent(scrollContainer)

	// Manage suggestions for the search bar
	searchEntry.OnChanged = func(input string) {
		if input == "" {
			showSuggestions(false)
			return
		}

		// Generate suggestions based on input
		suggestions := generateSuggestions(input, artistsData)
		suggestionsBox.RemoveAll()

		if len(suggestions) > 0 {
			// Create the suggestions list
			for _, suggestion := range suggestions {

				name := suggestion.Name
				suggestionButton := widget.NewButton(name, func() {
					searchEntry.SetText(name)                      // Sets the search bar text
					searchEntry.Refresh()                          // Refreshes the search bar to display the updated text
					showSuggestions(false)                         // Hides the suggestions
					updateArtistGrid(Searchbar(name, artistsData)) // Updates the grid with the search results
				})

				suggestionButton.Importance = widget.LowImportance
				suggestionButton.Resize(suggestionButton.MinSize())
				suggestionButton.Alignment = widget.ButtonAlignLeading // Alignement à gauche

				suggestionsBox.Add(suggestionButton)
			}

			if suggestionsContainer == nil {
				suggestionsContainer = widget.NewPopUp(suggestionsBox, W.Canvas())
			} else {
				suggestionsContainer.Content = suggestionsBox
				suggestionsContainer.Refresh()
			}
			suggestionsContainer.Move(fyne.NewPos(searchEntry.Position().X, searchEntry.Position().Y+searchEntry.Size().Height))
			suggestionsContainer.Resize(fyne.NewSize(searchEntry.Size().Width, suggestionsBox.MinSize().Height))
			showSuggestions(true)
		} else {
			showSuggestions(false)
		}

		// Display the suggestions list in a popup
		//showSuggestionsPopup(suggestionsList, searchEntry)
	}
	W.SetOnClosed(func() {
		A.Quit()
	})
	W.CenterOnScreen()
	W.Resize(fyne.NewSize(1000, 600))
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
			widget.NewLabel(fmt.Sprintf("Listes des members: %s", strings.Join(artist.Members, ", "))),
			widget.NewLabel(fmt.Sprintf("Date de publication du premier Album: %s", artist.FirstAlbum)),
			widget.NewLabel(fmt.Sprintf("Date de creation : %d", artist.DateCreation)),
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
		image,
		viewDetailsButton,
	)
	artistCard = container.NewPadded(artistCard)
	// Returns a container with artist's name, image,...
	return artistCard
}

/*func showSuggestionsPopup(suggestionsList *widget.List, entry *widget.Entry) {
    popup := widget.NewModalPopUp(suggestionsList, W.Canvas())
    popup.Show()
}*/
