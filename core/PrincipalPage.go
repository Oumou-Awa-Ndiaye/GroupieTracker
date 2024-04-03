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
	suggestionsContainer *widget.PopUp // Container for suggestions that will be added under the search bar
)

// CreateMainMenu creates the main menu for the application.
func CreateMainMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		// File menu with a Quit option
		fyne.NewMenu("File",
			fyne.NewMenuItem("Quit", func() {
				A.Quit()
			}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				widget.NewModalPopUp(
					fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
						widget.NewLabel("Groupie Tracker by Ynov Paris"),
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
	searchEntry.SetPlaceHolder("Search your artists...")

	// Initialize search button
	searchButton := widget.NewButton("Search", func() {
		searchTerm := searchEntry.Text
		results := Searchbar(searchTerm, artistsData)
		updateArtistGrid(results, W) // Update the artist grid with found artists
	})

	// Initialize filter button
	filterButton := widget.NewButton("Filter Search", func() {
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
	// Function to show or hide suggestions
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
	artistGrid = createArtistGrid(W)
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
					searchEntry.SetText(name)                         // Sets the search bar text
					searchEntry.Refresh()                             // Refreshes the search bar to display the updated text
					showSuggestions(false)                            // Hides the suggestions
					updateArtistGrid(Searchbar(name, artistsData), W) // Updates the grid with the search results
				})

				suggestionButton.Importance = widget.LowImportance
				suggestionButton.Resize(suggestionButton.MinSize())
				suggestionButton.Alignment = widget.ButtonAlignLeading // Left align

				suggestionsBox.Add(suggestionButton)
			}

			if suggestionsContainer == nil {
				suggestionsContainer = widget.NewPopUp(suggestionsBox, W.Canvas())
			} else {
				suggestionsContainer.Content = suggestionsBox
				suggestionsContainer.Refresh()
			}
			// Move suggestionsContainer below the search bar
			suggestionsContainer.Move(fyne.NewPos(searchEntry.Position().X, searchEntry.Position().Y+searchEntry.Size().Height))
			suggestionsContainer.Resize(fyne.NewSize(searchEntry.Size().Width, suggestionsBox.MinSize().Height))
			showSuggestions(true)
		} else {
			showSuggestions(false)
		}
	}

	W.SetOnClosed(func() {
		A.Quit()
	})
	W.CenterOnScreen()
	W.Resize(fyne.NewSize(1000, 600))
	W.Show()
}

// updateArtistGrid refreshes the artist grid with new data.
func updateArtistGrid(artists []Artist, W fyne.Window) {
	artistGrid.Objects = nil // Clear the existing grid
	for _, artist := range artists {
		artistGrid.Add(createArtistCard(artist, W))
	}
	artistGrid.Refresh() // Refresh the grid with updated artists
}

// createArtistGrid initializes and returns a grid layout containing artist cards.
func createArtistGrid(W fyne.Window) *fyne.Container {
	artistsData := GetArtists()
	grid := container.NewGridWithColumns(4) // Create a grid with 4 columns.

	for _, artist := range artistsData {
		artistCard := createArtistCard(artist, W) // Create a card for each artist.
		grid.Add(artistCard)                      // Add the card to the grid.
	}

	return grid // Return a grid layout with artists cards
}

func ArtistsPage(artist Artist) {
	W := A.NewWindow("Groupie Tracker")
	homeButton := widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		ShowHomePage(A)
		W.Hide()
	})
	image := getImageFromURL(artist.Image)

	content := container.NewVBox(
		container.NewHBox(
			homeButton,
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			canvas.NewText(artist.Name, theme.PrimaryColor()),
			layout.NewSpacer(),
		),

		container.NewVBox(
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewLabel("List of members:"),
				layout.NewSpacer(),
			),
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewLabel(strings.Join(artist.Members, ", ")),
				layout.NewSpacer(),
			),
		),

		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel(fmt.Sprintf("Date of first album release: %s", artist.FirstAlbum)),
			layout.NewSpacer(),
		),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel(fmt.Sprintf("Date of creation: %d", artist.DateCreation)),
			layout.NewSpacer(),
		),
		image,
	)

	W.SetOnClosed(func() {
		A.Quit()
	})

	W.SetContent(container.NewBorder(content, nil, nil, nil))
	W.CenterOnScreen()
	W.Resize(fyne.NewSize(1000, 600))
	W.Show()
}

// createArtistCard creates a visual card representing an artist, for use in the grid.
func createArtistCard(artist Artist, W fyne.Window) fyne.CanvasObject {
	// Create a label with the artist name
	nameLabel := widget.NewLabelWithStyle(artist.Name, fyne.TextAlignCenter, fyne.TextStyle{})
	nameLabel.Wrapping = fyne.TextTruncate // To ensure the text does not exceed the width of the card.

	image := getImageFromURL(artist.Image)
	image.FillMode = canvas.ImageFillContain
	viewDetailsButton := widget.NewButton(artist.Name, func() {
		ArtistsPage(artist)
		W.Hide()
	})
	artistCard := container.NewVBox(
		image,
		viewDetailsButton,
	)
	artistCard = container.NewPadded(artistCard)
	// Returns a container with artist's name, image...
	return artistCard
}
