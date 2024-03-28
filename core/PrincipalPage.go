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

// ShowHomePage setups and displays the home page of the application.
func ShowHomePage() {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher ...")
	var suggestionsPopup *widget.PopUp // A popup for showing search suggestions.

	searchEntry.OnChanged = func(term string) {
		// Hide existing popup if the search term is empty or update it with new suggestions.
		if suggestionsPopup != nil {
			suggestionsPopup.Hide()
		}

		if term == "" {
			return
		}

		suggestions := generateSuggestions(term, artistsData)
		if len(suggestions) == 0 {
			return
		}

		// List to display suggestions.
		list := widget.NewList(
			func() int {
				return len(suggestions)
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("")
			},
			func(i widget.ListItemID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(suggestions[i].Name)
			},
		)
		// When a suggestion is selected, set the search entry's text to it and hide the popup.
		list.OnSelected = func(i widget.ListItemID) {
			selectedArtist := suggestions[i]
			showArtistDetails(selectedArtist)
			searchEntry.SetText(suggestions[i].Name)
			if suggestionsPopup != nil {
				suggestionsPopup.Hide()
			}
		}

		// Show suggestions in a popup just below the search entry.
		suggestionsPopup = widget.NewPopUp(list, W.Canvas())
		suggestionsPopup.Show()

		entryPos := fyne.CurrentApp().Driver().AbsolutePositionForObject(searchEntry)
		suggestionsPopup.Move(fyne.NewPos(entryPos.X, entryPos.Y+searchEntry.Size().Height))
		suggestionsPopup.Resize(fyne.NewSize(searchEntry.Size().Width, list.MinSize().Height))
	}

	artistsData = GetArtists()
	artistGrid := createArtistGrid()
	// Setup the main content of the homepage.
	content := container.NewVBox(
		createToolbar(),
		canvas.NewText("Groupie Tracker", theme.PrimaryColor()),
		layout.NewSpacer(),
		searchEntry,
		artistGrid,
	)

	scrollContainer := container.NewVScroll(content)
	W.SetContent(scrollContainer)
}

func updateArtistGrid(artists []Artist) {
	artistGrid.Objects = nil // Effacer la grille existante
	for _, artist := range artists {
		artistGrid.Add(createArtistCard(artist))
	}
	artistGrid.Refresh() // Rafraîchir la grille pour afficher les modifications
}

func createToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			ShowHomePage()
		}),
	)
}

var artistGrid *fyne.Container

func createArtistGrid() *fyne.Container {
	artistsData := GetArtists()

	numColumns := 1

	grid := container.NewGridWithColumns(numColumns)

	for _, artist := range artistsData {
		artistLabel := widget.NewLabel(artist.Name)

		grid.Add(artistLabel)
	}

	// Retourner la grille contenant les artistes
	return grid
}

func showArtistDetails(artist Artist) {
	widget.NewModalPopUp(
		fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			canvas.NewText(artist.Name, theme.PrimaryColor()),
			canvas.NewImageFromFile(artist.Image),
			widget.NewLabel(fmt.Sprintf("Members: %s", strings.Join(artist.Members, ", "))),
			widget.NewLabel(fmt.Sprintf("First Album: %s", artist.FirstAlbum)),
			widget.NewLabel(fmt.Sprintf("Creation Date: %d", artist.DateCreation)),
		),
		W.Canvas(),
	).Show()
}
func createArtistCard(artist Artist) fyne.CanvasObject {
	// Créer un label avec le nom de l'artiste
	nameLabel := widget.NewLabel(artist.Name)

	// Créer un label avec les membres de l'artiste
	membersLabel := widget.NewLabel(fmt.Sprintf("Members: %s", strings.Join(artist.Members, ", ")))

	// Créer un label avec le premier album de l'artiste
	firstAlbumLabel := widget.NewLabel(fmt.Sprintf("First Album: %s", artist.FirstAlbum))

	// Créer un label avec la date de création de l'artiste
	creationDateLabel := widget.NewLabel(fmt.Sprintf("Creation Date: %d", artist.DateCreation))

	image := canvas.NewImageFromFile(artist.Image)

	// Créer une carte d'artiste contenant les labels et l'image
	artistCard := container.NewVBox(
		nameLabel,
		image,
		membersLabel,
		firstAlbumLabel,
		creationDateLabel,
		widget.NewButton("View Details", func() {
			showArtistDetails(artist)
		}),
	)

	return artistCard
}
func showArtistImage(artist Artist) {
	image := getImageFromURL(artist.Image)
	window := a.NewWindow(fmt.Sprintf("Image de %s", artist.Name))
	window.SetContent(image)
	window.Show()
}
