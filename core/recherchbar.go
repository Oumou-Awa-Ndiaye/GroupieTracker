package core

import (
	"fmt"
	"strings"
)

func Searchbar(query string, artistsData []Artist) []Artist {
	searchArtists := make([]Artist, 0)
	for _, artist := range artistsData {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			searchArtists = append(searchArtists, artist)
		}
	}
	return searchArtists
}

func GenerateSuggestions(research string, results []Artist) []Artist {
	var suggestions []Artist
	for _, artist := range results {
		suggestions = append(suggestions, Artist{
			Name: fmt.Sprintf("%s -Member", strings.ToLower(artist.Name)),
		})
		for _, member := range artist.Members {
			suggestions = append(suggestions, Artist{
				Name: fmt.Sprintf("%s - Member", strings.ToLower(member)),
			})
		}
	}
	return suggestions
}
