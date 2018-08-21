package dao

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func (in CreateInput) photo(id string, URL string) *Photo {
	p := new(Photo)
	p.AlbumID = in.AlbumID
	p.Tags = in.Tags
	p.Description = in.Description
	p.Date = in.Date
	p.ID = id
	p.URL = URL
	return p
}

func (ph *Photo) dbPutItemInput() (*dynamodb.PutItemInput, error) {
	item, err := dynamodbattribute.MarshalMap(ph)
	if err != nil {
		return nil, err
	}
	return &dynamodb.PutItemInput{
		TableName: aws.String("Photo"),
		Item:      item,
	}, nil
}

func (in QueryInput) dbQueryInput() (*dynamodb.QueryInput, error) {
	keyCond := expression.Key("AlbumID").Equal(expression.Value(in.Filter.AlbumID))
	filterCond := filterExpression(in.Filter)
	proj := projectExpression(in.Project)

	expression, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithFilter(filterCond).
		WithProjection(proj).
		Build()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &dynamodb.QueryInput{
		KeyConditionExpression:    expression.KeyCondition(),
		ProjectionExpression:      expression.Projection(),
		ExpressionAttributeNames:  expression.Names(),
		ExpressionAttributeValues: expression.Values(),
		TableName:                 aws.String("Photo"),
	}, nil
}

func filterExpression(filter FilterInput) expression.ConditionBuilder {
	var conditions expression.ConditionBuilder

	for _, tag := range filter.Tags {
		conditions = expression.And(conditions, expression.Contains(expression.Name("Tags"), tag))
	}

	if filter.StartDate != "" && filter.EndDate != "" {
		conditions = expression.And(conditions, expression.Between(expression.Name("Date"), expression.Value(filter.StartDate), expression.Value(filter.EndDate)))
	}

	if filter.Description != "" {
		conditions = expression.And(conditions, expression.Contains(expression.Name("Description"), filter.Description))
	}

	return conditions
}

func projectExpression(projection []string) expression.ProjectionBuilder {
	proj := expression.NamesList(expression.Name(projection[0]))
	for _, name := range projection[1:] {
		proj = expression.AddNames(proj, expression.Name(name))
	}
	return proj
}

// func mapStruct(in interface{}) map[string]interface{} {
// 	var inInterface map[string]interface{}
// 	inrec, _ := json.Marshal(in)
// 	return json.Unmarshal(inrec, &inInterface)
// }
