package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/linkedin", linkedinHandler)
	http.HandleFunc("/github", githubHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func linkedinHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.linkedin.com", http.StatusSeeOther)
}

func githubHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com", http.StatusSeeOther)
}
