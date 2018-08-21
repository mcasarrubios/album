package dao

// BasicPhoto info of a photo photo
type BasicPhoto struct {
	AlbumID     string   `json:"albumId"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
}

// Photo model
type Photo struct {
	BasicPhoto
	ID  string `json:"id"`
	URL string `json:"url"`
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
	Filter  FilterInput
	Project []string
	Limit   int
	LastKey map[string]string
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
