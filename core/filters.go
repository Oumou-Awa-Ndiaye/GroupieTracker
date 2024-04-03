package core

import (
	"fyne.io/fyne/v2/widget"
)

// FilterArtistsByCreationDate filters artists by the date of their first album
func FilterArtistsByCreationDate(startyear, endyear int, artists []Artist) []Artist {
	filteredArtists := make([]Artist, 0)
	for _, artist := range artists {
		if artist.DateCreation >= startyear && artist.DateCreation <= endyear {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	return filteredArtists
}

// Function to filter artists based on the number of selected members
func filterArtistsByNumMembers(artists []Artist, membersChecks []*widget.Check) []Artist {
	filteredArtists := make([]Artist, 0)
	checkedNumbers := getCheckedNumbers(membersChecks...)

	for _, artist := range artists {
		for _, nbr := range checkedNumbers {
			if len(artist.Members) == nbr {
				filteredArtists = append(filteredArtists, artist)
				// No need to continue the inner loop once a match is found
				break
			}
		}
	}
	return filteredArtists
}

func getCheckedNumbers(checks ...*widget.Check) []int {
	var checkedNumbers []int
	for i, check := range checks {
		if check.Checked {
			checkedNumbers = append(checkedNumbers, i+1) // Add 1 since member numbers start from 1
		}
	}
	return checkedNumbers
}
