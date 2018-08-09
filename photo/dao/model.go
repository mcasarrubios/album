package dao

import uuid "github.com/satori/go.uuid"

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

type DAO struct {
	db *DynamoDB
}

type DataAccesor interface {
	Create(input CreateInput) (*Photo, error)
	// Get(input GetInput) (*Photo, error)
	// List(query QueryInput) ([]Photo, error)
	// Delete(input CreateInput) (error)
}

func New(db DB) DataAccesor {
	return &DAO{db}
}

func (dao *DAO) Create(input CreateInput) (*Photo, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	ph.ID = input.photo(id)
	return putItem(ph)
}

func (in CreateInput) photo(id string) *Photo {
	p := new(Photo)
	p.AlbumID = in.AlbumID
	p.Tags = in.Tags
	p.Description = in.Description
	p.Date = in.Date
	p.ID = id
	return p
}
