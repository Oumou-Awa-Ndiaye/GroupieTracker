package core

import (
	"fyne.io/fyne/v2/widget"
)

// FilterArtistsByCreationDate filtre les artistes par date de premier album
func FilterArtistsByCreationDate(startyear, endyear int, artists []Artist) []Artist {
	filteredArtists := make([]Artist, 0)
	for _, artist := range artists {
		if artist.DateCreation >= startyear && artist.DateCreation <= endyear {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	return filteredArtists
}

// Fonction pour filtrer les artistes en fonction du nombre de membres sélectionnés
func filterArtistsByNumMembers(artists []Artist, membersChecks []*widget.Check) []Artist {
	filteredArtists := make([]Artist, 0)
	checkedNumbers := getCheckedNumbers(membersChecks...)

	for _, artist := range artists {
		for _, nbr := range checkedNumbers {
			if len(artist.Members) == nbr {
				filteredArtists = append(filteredArtists, artist)
				// Pas besoin de continuer la boucle intérieure une fois qu'un match est trouvé
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
			checkedNumbers = append(checkedNumbers, i+1) // Ajouter 1 car les nombres de membres commencent à partir de 1
		}
	}
	return checkedNumbers
}
