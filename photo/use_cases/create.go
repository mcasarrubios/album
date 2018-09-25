package usecases

import (
	"github.com/mcasarrubios/album/errors"
	"github.com/mcasarrubios/album/photo/entities"
)

// CreatePhoto struct
type CreatePhoto struct {
	dao *entities.DataAccessor
}

// Execute creates a photo
func (cmd *CreatePhoto) Execute(input CreateInput) (*Photo, error) {
	return cmd.dao.Create(input)
}

// CanExecute checks if everything is ok before executing the command
func (cmd *CreatePhoto) CanExecute(context interface{}) error {
	err := authorize(data)
	if err != nil {
		return nil, errors.New(errors.AuthorizationError, errorMsg(cmd, err))
	}

	err = validate(data)
	if err != nil {
		return nil, errors.New(errors.ValidationError, errorMsg(cmd, err))
	}
	return nil
}

func authorize(input CreateInput) error {
	return nil
}

func validate(input CreateInput) error {
	return input.Validate()
}

func errorMsg(cmd Command, err error) string {
	return "Can't create a photo: " + err.Error()
}

// // Create a photo
// func (dao *entities.DataAccessor) Create(input CreateInput, URL string) (*Photo, error) {

// 	// Generate Id
// 	// Save file
// 	// Save in store

// 	id, err := uuid.NewV4()
// 	if err != nil {
// 		return nil, err
// 	}

// 	ph := input.photo(id.String(), URL)
// 	// putItemInput, err := ph.dbPutItemInput()
// 	dao.Create()
// 	fmt.Println(putItemInput)
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = dao.db.PutItem(putItemInput)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ph, nil
// }
