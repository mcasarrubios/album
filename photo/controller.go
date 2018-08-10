package photo

import (
	"encoding/json"
	"io"
	"net/http"

	DAO "github.com/mcasarrubios/album/photo/dao"
)

type control struct {
	dao DAO.DataAccessor
}

// HTTPController of photo resource
type HTTPController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

// NewController creates a photo controller
func NewController() (HTTPController, error) {
	db, err := DAO.OpenDB()
	dao := DAO.New(db)
	return &control{dao: dao}, err
}

// Create a photo
func (ctrl *control) Create(w http.ResponseWriter, r *http.Request) {
	input, err := decode(r.Body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if input.AlbumID == "" || input.Date == "" {
		http.Error(w, "Fill required fields", http.StatusBadRequest)
		return
	}

	ph, err := ctrl.create(input)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusBadRequest)
		return
	}

	sendJSON(w, ph)
}

func (ctrl *control) create(input DAO.CreateInput) (*DAO.Photo, error) {
	URL := "http://my-album.awesome-photo.jpg"
	return ctrl.dao.Create(input, URL)
}

func (ctrl *control) Get(w http.ResponseWriter, r *http.Request) {
	ph := &DAO.Photo{
		URL: "https://static.allcloud.com/assets/images/blog/golang.png",
		QueryInput: DAO.QueryInput{
			ID: "213",
			CreateInput: DAO.CreateInput{
				Tags:        []string{"tag-1", "tag-2"},
				Description: "Awesome description",
				Date:        "2008-09-15T15:53:00+05:00",
			},
		},
	}

	sendJSON(w, ph)
}

func decode(body io.ReadCloser) (DAO.CreateInput, error) {
	decoder := json.NewDecoder(body)
	payload := new(DAO.CreateInput)
	err := decoder.Decode(&payload)
	return *payload, err
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
