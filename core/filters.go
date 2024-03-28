package core

import (
	"strconv"
	"strings"
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

// FilterArtistsByMembersCount filtre les artistes par nombre de membres
func FilterArtistsByMembersCount(count int, artists []Artist) []Artist {
	filteredArtists := make([]Artist, 0)
	for _, artist := range artists {
		if len(artist.Members) == count {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

// FilterArtistsByConcertLocation filtre les artistes par lieu de concerts
//func FilterArtistsByConcertLocation(location string, artists []Artist) []Artist {
	//filteredArtists := make([]Artist, 0)
	//for _, artist := range artists {
		//if strings.Contains(strings.ToLower(artist.ConcertLocation), strings.ToLower(location)) {
		//	filteredArtists = append(filteredArtists, artist)
	//	}
	//}
	//return filteredArtists
//}
