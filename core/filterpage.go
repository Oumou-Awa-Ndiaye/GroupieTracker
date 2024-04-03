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

	// Function to update the Label with the year corresponding to the slider's value
	updateLabelYear := func(label *widget.Label, value float64) {
		year := int(value)
		label.SetText(strconv.Itoa(year))
	}
	artistGrid = createArtistGrid(W)

	// Create Labels to display the years
	labelCreationDateStart := widget.NewLabel("1958")
	labelCreationDateEnd := widget.NewLabel("2015")

	sliderCreationDateStart := widget.NewSlider(1958, 2015) // 1958 for Bee Gees and 2015 for Juice Wrld
	sliderCreationDateEnd := widget.NewSlider(1958, 2015)   // 1958 for Bee Gees and 2015 for Juice Wrld

	sliderCreationDateStart.SetValue(1958)
	sliderCreationDateEnd.SetValue(2015)

	// Update the Labels whenever the slider values change
	sliderCreationDateStart.OnChanged = func(value float64) {
		updateLabelYear(labelCreationDateStart, value)
	}

	sliderCreationDateEnd.OnChanged = func(value float64) {
		updateLabelYear(labelCreationDateEnd, value)
	}

	startDateRange := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Start Date:"),
			labelCreationDateStart,
			labelCreationDateEnd,
		),
		container.NewVBox(
			sliderCreationDateStart,
			sliderCreationDateEnd,
		),
		container.NewHBox(),
	)

	// Declaration of a slice of Check widgets
	var membersChecks []*widget.Check

	// Loop to create and initialize the checkboxes
	for i := 1; i <= 7; i++ {
		memberCheck := widget.NewCheck(strconv.Itoa(i), func(checked bool) {})
		memberCheck.SetChecked(false) // Initialize as false so they are not checked by default
		membersChecks = append(membersChecks, memberCheck)
	}

	// Convert membersChecks into []fyne.CanvasObject
	var canvasObjects []fyne.CanvasObject
	for _, check := range membersChecks {
		canvasObjects = append(canvasObjects, check)
	}

	// Creation of the VBox for members
	numMembers := container.NewVBox(
		widget.NewLabel("Number of Members:"),
		container.NewHBox(canvasObjects...), // Use of the spread operator to add all elements of the slice
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
