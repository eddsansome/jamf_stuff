package main

import (
	"log"
	"net/http"
)

func main() {

	// imagine I am floating in the cloud - maybe as lambda function, OOOOoooOOO
	http.HandleFunc("/jamf_log", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		justification := r.Form.Get("justification")
		user := r.Form.Get("user")

		// this will be logged to S3 rather than stdout
		log.Printf("User %s requested sudo, justification: %s", user, justification)

	})
	log.Println("starting service on port 3000")
	http.ListenAndServe(":3000", nil)
}
