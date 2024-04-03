package core

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func FilterPage(A fyne.App) {
	W := A.NewWindow("Groupie Trackers")

	// Fonction pour mettre à jour le Label avec l'année correspondant à la valeur du slider
	updateLabelYear := func(label *widget.Label, value float64) {
		year := int(value)
		label.SetText(strconv.Itoa(year))
	}
	artistGrid = createArtistGrid(W)

	// Créer des Labels pour afficher les années
	labelCreationDateStart := widget.NewLabel("1958")
	labelCreationDateEnd := widget.NewLabel("2015")

	sliderCreationDateStart := widget.NewSlider(1958, 2015) // 1958 for Bee Gees and 2015 for Juice Wrld
	sliderCreationDateEnd := widget.NewSlider(1958, 2015)   // 1958 for Bee Gees and 2015 for Juice Wrld

	sliderCreationDateStart.SetValue(1958)
	sliderCreationDateEnd.SetValue(2015)

	// Mettre à jour les Labels à chaque fois que la valeur des sliders change
	sliderCreationDateStart.OnChanged = func(value float64) {
		updateLabelYear(labelCreationDateStart, value)
	}

	sliderCreationDateEnd.OnChanged = func(value float64) {
		updateLabelYear(labelCreationDateEnd, value)
	}

	startDateRange := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Date de Début :"),
			labelCreationDateStart,
			labelCreationDateEnd,
		),
		container.NewVBox(
			sliderCreationDateStart,
			sliderCreationDateEnd,
		),
		container.NewHBox(),
	)

	// Déclaration d'une slice de widgets Check
	var membersChecks []*widget.Check

	// Boucle pour créer et initialiser les cases à cocher
	for i := 1; i <= 7; i++ {
		memberCheck := widget.NewCheck(strconv.Itoa(i), func(checked bool) {})
		memberCheck.SetChecked(false) // Initialiser à false pour qu'elles ne soient pas cochées par défaut
		membersChecks = append(membersChecks, memberCheck)
	}

	// Convertir membersChecks en []fyne.CanvasObject
	var canvasObjects []fyne.CanvasObject
	for _, check := range membersChecks {
		canvasObjects = append(canvasObjects, check)
	}

	// Création du VBox pour les membres
	numMembers := container.NewVBox(
		widget.NewLabel("Nombre de Membres :"),
		container.NewHBox(canvasObjects...), // Utilisation de l'opérateur spread pour ajouter tous les éléments de la slice
	)

	applyButton := widget.NewButton("Apply Filters", func() {
		artists := GetArtists()
		artists = FilterArtistsByCreationDate(int(sliderCreationDateStart.Value), int(sliderCreationDateEnd.Value), artists)
		artists = filterArtistsByNumMembers(artists, membersChecks)
		// fmt.Println(artists)
		updateArtistGrid(artists, W)

	})

	homeButton := widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		ShowHomePage(A)
		W.Hide()
	})

	content := container.NewVBox(
		container.NewHBox(
			homeButton,
			layout.NewSpacer(),
		),
		startDateRange,
		numMembers,
		applyButton,
	)
	scrollContainer := container.NewVScroll(artistGrid)

	W.SetOnClosed(func() {
		A.Quit()
	})

	W.SetContent(container.NewBorder(content, nil, nil, nil, scrollContainer))
	W.CenterOnScreen()
	W.Resize(fyne.NewSize(1000, 600))
	W.Show()
}
