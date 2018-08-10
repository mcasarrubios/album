package dao

import "github.com/aws/aws-sdk-go/service/dynamodb"

// CreateInput data to create a photo
type CreateInput struct {
	AlbumID     string   `json:"albumId"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
}

// QueryInput params to query photos
type QueryInput struct {
	CreateInput
	ID string `json:"id"`
}

// Photo model
type Photo struct {
	URL string `json:"url"`
	QueryInput
}

// DAO access to DB
type DAO struct {
	db *dynamodb.DynamoDB
}

// DataAccessor accesor to DB
type DataAccessor interface {
	Create(input CreateInput, URL string) (*Photo, error)
	// Get(input GetInput) (*Photo, error)
	// List(query QueryInput) ([]Photo, error)
	// Delete(input CreateInput) (error)
}
