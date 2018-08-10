package main

import (
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	"github.com/mcasarrubios/album/photo"
)

// App application type
type App struct {
	ctrl photo.HTTPController
}

// use JSON logging when run by Up (including `up start`).
func init() {
	if os.Getenv("UP_STAGE") == "" {
		log.SetHandler(text.Default)
	} else {
		log.SetHandler(json.Default)
	}
}

func main() {
	ctrl, err := photo.NewController()
	if err != nil {
		log.WithError(err).Fatal("error creating photo controller")
	}
	app := &App{ctrl}
	http.HandleFunc("/", app.router)
	addr := ":" + os.Getenv("PORT")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}

func (app *App) router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.ctrl.Get(w, r)
	case http.MethodPost:
		app.ctrl.Create(w, r)
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		// Give an error message.
	}
}
