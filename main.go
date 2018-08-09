package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/apex/log"
	apexJSON "github.com/apex/log/handlers/json"
	apexText "github.com/apex/log/handlers/text"
	"github.com/mcasarrubios/album/errors"
	"github.com/mcasarrubios/album/photo"
)

// use JSON logging when run by Up (including `up start`).
func init() {
	if os.Getenv("UP_STAGE") == "" {
		log.SetHandler(apexText.Default)
	} else {
		log.SetHandler(apexJSON.Default)
	}
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
		// get(w, r)
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

func create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := new(photo.dao.CreateInput)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ph, err = photo.Create(body)
	if err != nil {
		sendError(w, err)
		return
	}
	sendJSON(w, ph)
}

// func get(w http.ResponseWriter, r *http.Request) {
// 	ph := &photo{
// 		URL:         "https://static.allcloud.com/assets/images/blog/golang.png",
// 		Tags:        []string{"tag-1", "tag-2"},
// 		Description: "Awesome description",
// 		Date:        "2008-09-15T15:53:00+05:00",
// 	}

// 	sendJSON(w, ph)
// }

func sendJSON(w http.ResponseWriter, item interface{}) {
	js, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func sendError(w http.ResponseWriter, err errors.HTTP) {
	if err.StatusCode == 500 {
		log.Error(err.Error())
		errMsg := "Internal server error"
	} else {
		errMsg := err.Error()
	}
	http.Error(w, errMsg, err.statusCode)
}
