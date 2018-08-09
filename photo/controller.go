package photo

import (
	"net/http"

	"github.com/mcasarrubios/album/errors"
	"github.com/mcasarrubios/album/photo/dao"
)

// Create a photo
func Create(input dao.CreateInput) (dao.Photo, errors.HT) {
	if input.AlbumID == "" || input.Date == "" {
		return nil, &errors.HTTP{&error{"Fill required fields"}, http.StatusBadRequest}
	}
	ph, err := dao.Create(input)
	if err != nil {
		return nil, &errors.HTTP{&error{"Fill required fields"}, http.StatusBadRequest}
	}
	return ph, nil
}
