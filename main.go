package main

import (
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	"github.com/julienschmidt/httprouter"
	"github.com/mcasarrubios/album/photo"
)

// use JSON logging when run by Up (including `up start`).
func init() {
	if os.Getenv("UP_STAGE") == "" {
		log.SetHandler(text.Default)
	} else {
		log.SetHandler(json.Default)
	}
}

func main() {
	photoCtrl, err := photo.NewController()
	if err != nil {
		log.WithError(err).Fatal("error creating photo controller")
	}
	router := httprouter.New()
	router.GET("/", photoCtrl.List)
	router.POST("/", photoCtrl.Create)
	router.GET("/:photoID", photoCtrl.Get)

	addr := ":" + os.Getenv("PORT")
	if err := http.ListenAndServe(addr, router); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}
