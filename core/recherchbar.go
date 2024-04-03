package core

import (
	"strconv"
	"strings"
)

// Searchbar searches through artists based on the query and returns matching artists.
func Searchbar(query string, artistsData []Artist) []Artist {
	searchArtists := make([]Artist, 0)
	for _, artist := range artistsData {
		// Check if the search term matches the artist's name, members, creation date, or first album date
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) ||
			containsMember(artist.Members, query) ||
			strconv.Itoa(artist.DateCreation) == query ||
			strings.ToLower(artist.FirstAlbum) == strings.ToLower(query) {
			searchArtists = append(searchArtists, artist)
		}
	}
	return searchArtists
}

// Utility function to check if the artist's members contain the search term
func containsMember(members []string, query string) bool {
	for _, member := range members {
		if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
			return true
		}
	}
	return false
}

// Function to generate suggestions based on the search term
func generateSuggestions(query string, artistsData []Artist) []Artist {
	suggestions := make([]Artist, 0)
	for _, artist := range Searchbar(query, artistsData) {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			suggestions = append(suggestions, artist)
		}
	}
	return suggestions
}
