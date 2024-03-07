/*Partie Lisa*/

package core

import (
	"fmt"
	"strings"
)

func Searchbar(research string, artists []Artist) []Artist {
	var results []Artist
	for _, artist := range artists {
		//1_On va déja verifier si le nom d'artiste correspond a la recherche ou pas
		if strings.Contains((artist.Name), research) { //si le nom correspond on le rajoute a la liste
			results = append(results, artist)
			continue
		}
		// 2_Vérifier si un membre  du groupe des l'artiste correspond à la recherche
		for _, member := range artist.Members {
			if strings.Contains(member, research) {
				results = append(results, artist)
				break
			}
		}
		// 3-Vérifier si l'emplacement de l'artiste correspond à la recherche
		if strings.Contains((artist.Locations), research) {
			results = append(results, artist)
			continue
		}
		//3-Verifier la date du premier album de l'artiste correspond a la requete
		if strings.Contains((artist.FirstAlbum), research) {
			results = append(results, artist)
			continue
		}
		//5_Verifier la date de creation si ca correspond a la recherche
		if fmt.Sprintf("%d", artist.DateCreation) == research {
			results = append(results, artist)
			continue
		}
	}
	return results
}
func GenerateSuggestions(research string, results []Artist) []Artist {
	var suggestions []Artist
	for _, artist := range results {
		suggestions = append(suggestions, Artist{
			Name: fmt.Sprintf("%s -Member", artist.Name),
		})
		for _, member := range artist.Members {
			suggestions = append(suggestions, Artist{
				Name: fmt.Sprintf("%s - Member", member),
			})
		}
	}
	return suggestions
}
