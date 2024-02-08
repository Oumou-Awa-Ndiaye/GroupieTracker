package functions

import (
	"fmt"
	"net/http"
)

func searchBar(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		fmt.Fprintf(w, "Please provide a search query.")
		return
	}

	fmt.Fprintf(w, "Search results for: %s", query)
}

func main() {
	http.HandleFunc("/search", searchBar)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
