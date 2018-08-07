package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/apex/log"
	apexJSON "github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
)

// use JSON logging when run by Up (including `up start`).
func init() {
	if os.Getenv("UP_STAGE") == "" {
		log.SetHandler(text.Default)
	} else {
		log.SetHandler(apexJSON.Default)
	}
}

type photo struct {
	AlbumID     string   `json:"albumId"`
	ID          string   `json:"id"`
	URL         string   `json:"url"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
}

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", router)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		create(w, r)
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		// Give an error message.
	}
}

func create(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	ph := new(photo)
	err := decoder.Decode(&ph)
	if err != nil {
		return serverError(w, err)
	}

	if ph.AlbumID == "" || ph.Date == "" {
		return clientError(w, http.StatusBadRequest)
	}

	err = putItem(ph)
	if err != nil {
		return serverError(w, err)
	}
}

func get(w http.ResponseWriter, r *http.Request) {

	ph := &photo{
		URL:         "https://static.allcloud.com/assets/images/blog/golang.png",
		Tags:        []string{"tag-1", "tag-2"},
		Description: "Awesome description",
		Date:        "2008-09-15T15:53:00+05:00",
	}

	sendJSON(w, ph)
}

func sendJSON(w http.ResponseWriter, item interface{}) {
	js, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func clientError(w http.ResponseWriter, statusError int) {
	return http.Error(w, statusError)
}

func serverError(w http.ResponseWriter, err error) {
	return http.Error(w, err.Error(), http.StatusInternalServerError)
}
