package core

import (
	"strconv"
	"strings"
	"fyne.io/fyne/v2/widget"
)

// FilterArtistsByFirstAlbum filtre les artistes par date de premier album
func FilterArtistsByFirstAlbum(year int, artists []Artist) []Artist {
	filteredArtists := make([]Artist, 0)
	for _, artist := range artists {
		artistFirstAlbumYear, err := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[0])
		if err != nil {
			continue
		}
		if artistFirstAlbumYear == year {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}


//Filtre par membres
/*func filterArtistsByMember(members string, artists []Artist) []Artist {
	filteredArtists := make([]Artist, 0)
	for _, artist := range artists {
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(members)) {
				filteredArtists = append(filteredArtists, artist)
				break
			}
		}
	}
	return filteredArtists
}*/

// Fonction pour filtrer les artistes en fonction du nombre de membres sélectionnés
func filterArtistsByNumMembers(numMembers int, artists []Artist, membersChecks []*widget.Check) []Artist {
    filteredArtists := make([]Artist, 0)
    for _, artist := range artists {
        if len(artist.Members) == numMembers {
            filteredArtists = append(filteredArtists, artist)
        }
    }
    return filteredArtists
}