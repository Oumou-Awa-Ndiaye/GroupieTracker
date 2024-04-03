package core

import (
	"strconv"
	"strings"
)

func Searchbar(query string, artistsData []Artist) []Artist {
	searchArtists := make([]Artist, 0)
	for _, artist := range artistsData {
		// Vérifiez si le terme de recherche correspond au nom de l'artiste, aux membres, à la date de création ou à la date du premier album
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) ||
			containsMember(artist.Members, query) ||
			strconv.Itoa(artist.DateCreation) == query ||
			strings.ToLower(artist.FirstAlbum) == strings.ToLower(query) {
			searchArtists = append(searchArtists, artist)
		}
	}
	return searchArtists
}


// Fonction utilitaire pour vérifier si les membres de l'artiste contiennent le terme de recherche
func containsMember(members []string, query string) bool {
	for _, member := range members {
		if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
			return true
		}
	}
	return false
}

// Fonction pour générer des suggestions en fonction du terme de recherche
func generateSuggestions(query string, artistsData []Artist) []Artist {
	suggestions := make([]Artist, 0)
	for _, artist := range Searchbar(query, artistsData) {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			suggestions = append(suggestions, artist)
		}
	}
	return suggestions
}

