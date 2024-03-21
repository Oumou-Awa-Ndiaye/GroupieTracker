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

var (
	a           = app.New()
	W           = a.NewWindow("Groupie Tracker")
	artistsData []Artist
)

func CreateMainMenu() *fyne.MainMenu {
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
						widget.NewLabel("Groupie Tracker  Par Ynov Paris  "),
					),
					W.Canvas(),
				).Show()
			}),
		),
	)
}

func ShowHomePage() {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for artists...")
	searchEntry.OnChanged = func(term string) {
		results := Searchbar(term, artistsData)
		updateArtistGrid(results)
	}
	artistGrid = createArtistGrid()
	artistsData = GetArtists()
	// Créer le contenu principal de la page
	content := container.NewVBox(
		createToolbar(),
		canvas.NewText("Groupie Tracker", theme.PrimaryColor()),
		layout.NewSpacer(),
		searchEntry, // Ajout DE la barre de recherche
		artistGrid,  // la grille des artistes
	)

	// Créer un conteneur avec défilement pour le contenu principal
	scrollContainer := container.NewVScroll(content)

	// Définir le contenu de la fenêtre
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
func createArtistCard(artist Artist) fyne.CanvasObject {
	image := canvas.NewImageFromFile(artist.Image)
	image.FillMode = canvas.ImageFillContain

	nameLabel := widget.NewLabel(artist.Name)
	membersLabel := widget.NewLabel(fmt.Sprintf("Members: %s", strings.Join(artist.Members, ", ")))
	firstAlbumLabel := widget.NewLabel(fmt.Sprintf("First Album: %s", artist.FirstAlbum))
	creationDateLabel := widget.NewLabel(fmt.Sprintf("Creation Date: %d", artist.DateCreation))

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

func searchArtists(term string) []Artist {
	return Searchbar(term, artistsData)
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
