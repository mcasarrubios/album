package dao

type CreateInput struct {
	AlbumID     string   `json:"albumId"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
}

type QueryInput struct {
	CreateInput
	ID string `json:"id"`
}

type Photo struct {
	URL string `json:"url"`
	QueryInput
}

func Create(input CreateInput) (*Photo, error) {
	ph := new(Photo)
	ph.ID = "1"
	return ph, nil
}
