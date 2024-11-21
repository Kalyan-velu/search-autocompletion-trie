package main

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"
)

// Initialize the Trie
var trie = NewTrie()

func main() {
	fmt.Println("Starting server on port 8080")
	initialWords := []string{"hello", "world", "hi", "hey", "how", "are", "you", "doing", "today"}

	for _, word := range initialWords {
		trie.Insert(strings.ToLower(word))
	}

	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/autocomplete", autocompleteHandler)
	http.ListenAndServe(":8080", nil)
}

// Serve the index.html file
func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/index.html")
}

// Suggestion represents a single autocomplete suggestion
type Suggestion struct {
	Value string `json:"value"`
}

// handle  autocomplete requests
// autocompleteHandler handles HTTP requests for the autocomplete feature.
// It reads the query parameter "q" from the request URL, processes the input text,
// and returns a list of suggestions in HTML format.
//
// The function performs the following steps:
// 1. Trims whitespace from the query and splits it into words.
// 2. If there are multiple words, it adds valid words (longer than 2 characters) to the trie.
// 3. Searches the trie for suggestions based on the last word in the query.
// 4. Constructs an HTML response with the suggestions and writes it to the response writer.
//
// Parameters:
// - w: The HTTP response writer.
// - r: The HTTP request.
func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	text := strings.TrimSpace(query)
	words := strings.Fields(text)
	if len(words) == 0 {
		return
	}
	lastWord := words[len(words)-1]

	// add  new word to trie
	if len(words) > 1 {
		for _, word := range words[:len(words)-1] {
			word = strings.ToLower(word)
			if isValidWord(word) && len(word) > 2 { // Only add words longer than 2 characters
				trie.Insert(word)
			}
		}
	}

	suggestions := trie.Search(lastWord)

	var response strings.Builder

	for _, suggestion := range suggestions {
		// send a list of suggestions in HTML format
		response.WriteString(fmt.Sprintf("<li>%s</li>", suggestion))
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response.String()))
}

func isValidWord(word string) bool {
	for _, r := range word {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Handle autocomplete requests
// func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query().Get("q")
// 	fmt.Printf("Query: %s\n", query)
// 	suggestions := trie.Search(strings.ToLower(query))
// 	var response strings.Builder
// 	for _, suggestion := range suggestions {
// 		response.WriteString(fmt.Sprintf("<li>%s</li>", suggestion))
// 	}
// 	w.Header().Set("Content-Type", "text/html")
// 	w.Write([]byte(response.String()))
// }
