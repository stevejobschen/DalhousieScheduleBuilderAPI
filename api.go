package main

import (
	"fmt"
	"net/http"
	"time"

	"./db"
)

func main() {
	fmt.Printf("API has started https running on port: 8080... \n")
	http.HandleFunc("/api/courses", courses)
	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}

func courses(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	tn := t.Format("2006/01/02 15:04:05")
	fmt.Printf("[" + tn + "] New connection made")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	category := r.URL.Query().Get("s")
	term := r.URL.Query().Get("t")
	fmt.Fprintf(w, string(db.Open(category, term)))
}

