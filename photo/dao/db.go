package dao

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
)

// DBProvider provider to DB
type DBProvider interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	// GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	Query(query *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

// New Creates a DAO
func New(db DBProvider) DataAccessor {
	return &DAO{db}
}

// OpenDB opens a session in the DB
func OpenDB() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(getConfig().session)
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}

// Create a photo
func (dao *DAO) Create(input CreateInput, URL string) (*Photo, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	ph := input.photo(id.String(), URL)
	putItemInput, err := ph.dbPutItemInput()
	fmt.Println(putItemInput)
	if err != nil {
		return nil, err
	}

	_, err = dao.db.PutItem(putItemInput)
	if err != nil {
		return nil, err
	}
	return ph, nil
}

// List query photos
func (dao *DAO) List(query QueryInput) (*QueryOutput, error) {
	queryInput, err := query.dbQueryInput()
	if err != nil {
		return nil, err
	}
	return dao.query(queryInput)
}

// Get a photo
func (dao *DAO) Get(input GetInput) (*Photo, error) {
	queryInput, err := input.dbQueryInput()
	if err != nil {
		return nil, err
	}
	output, err := dao.query(queryInput)

	if len(output.Items) == 1 {
		return &output.Items[0], nil
	}
	return nil, nil
}

func (dao *DAO) query(queryInput *dynamodb.QueryInput) (*QueryOutput, error) {
	output, err := dao.db.Query(queryInput)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return queryOutput(output)
}
