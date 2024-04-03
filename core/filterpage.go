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

	// Créer des Labels pour afficher les années
	labelStartDate := widget.NewLabel("1986")
	sliderStartDate := widget.NewSlider(1958, 2015) // 1958 for Bee Gees and 2015 for Juice Wrld
	sliderStartDate.SetValue(1958)

	// Mettre à jour les Labels à chaque fois que la valeur des sliders change
	sliderStartDate.OnChanged = func(value float64) {
		updateLabelYear(labelStartDate, value)
	}

	startDateRange := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Date de Début :"),
			labelStartDate,
		),
		container.NewVBox(
			sliderStartDate,
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
	)

	W.SetOnClosed(func() {
		A.Quit()
	})

	W.SetContent(container.NewBorder(content, nil, nil, nil))
	W.CenterOnScreen()
	W.Resize(fyne.NewSize(1000, 600))
	W.Show()
}
