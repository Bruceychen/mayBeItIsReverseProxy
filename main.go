package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/linkedin", linkedinHandler)
	http.HandleFunc("/github", githubHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func linkedinHandler(w http.ResponseWriter, r *http.Request) {
	// Validate that the input URL is a valid URL
	inputURL := r.URL.Query().Get("url")
	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Validate that the input URL is the expected LinkedIn homepage URL
	if !strings.HasPrefix(inputURL, "https://www.linkedin.com") {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Redirect the user to the input URL
	http.Redirect(w, r, inputURL, http.StatusSeeOther)
}

func githubHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com", http.StatusSeeOther)
}
