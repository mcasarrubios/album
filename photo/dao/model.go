package dao

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
}

// Photo model
type Photo struct {
	BasicPhoto
	ID  string `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

// CreateInput data to create a photo
type CreateInput struct {
	BasicPhoto
}

// FilterInput fields
type FilterInput struct {
	AlbumID     string   `json:"albumId"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
}

// QueryInput params to query photos
type QueryInput struct {
	Filter  FilterInput       `json:"filter"`
	Project []string          `json:"project"`
	Limit   int               `json:"limit"`
	LastKey map[string]string `json:"lastKey"`
}

// QueryOutput results of querying photos
type QueryOutput struct {
	Items   []Photo
	LastKey KeyPhoto
}

// DAO access to DB
type DAO struct {
	db DBProvider
}

// DataAccessor accesor to DB
type DataAccessor interface {
	Create(input CreateInput, URL string) (*Photo, error)
	// Get(input GetInput) (*Photo, error)
	List(query QueryInput) ([]Photo, error)
	// Delete(input CreateInput) (error)
}
