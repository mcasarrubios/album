package entities

// KeyPhoto key index fields
type KeyPhoto struct {
	AlbumID string `json:"albumId,omitempty"`
	Date    string `json:"date,omitempty"`
}

// BasicPhoto info of a photo photo
type BasicPhoto struct {
	KeyPhoto
	Tags        []string `json:"tags,omitempty"`
	Description string   `json:"description,omitempty"`
	URL         string   `json:"url,omitempty"`
}

// Photo model
type Photo struct {
	BasicPhoto
	ID string `json:"id,omitempty"`
}

// // Create a photo
// func (dao *DataAccessor) Create(input CreateInput) (*Photo, error) {

// 	err := validate(input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return dao.Create(input)
// }
