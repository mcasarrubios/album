package dao

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/mcasarrubios/album/common"
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
	if in.Filter.AlbumID == "" {
		return nil, errors.New("Missing required fields")
	}
	expr := expression.NewBuilder().WithKeyCondition(keyCondExpression(in.Filter))

	if len(in.Filter.Tags) > 0 || in.Filter.Description != "" {
		expr = expr.WithFilter(filterExpression(in.Filter))
	}

	if len(in.Project) > 0 {
		expr = expr.WithProjection(projectExpression(in.Project))
	}

	exprBuild, err := expr.Build()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query := &dynamodb.QueryInput{
		KeyConditionExpression:    exprBuild.KeyCondition(),
		FilterExpression:          exprBuild.Filter(),
		ProjectionExpression:      exprBuild.Projection(),
		ExpressionAttributeNames:  exprBuild.Names(),
		ExpressionAttributeValues: exprBuild.Values(),
		TableName:                 aws.String("Photo"),
	}

	if in.Limit > 0 {
		query.SetLimit(int64(in.Limit))
	}

	if in.StartKey != "" {
		startKey, err := decodeStartKey(in.StartKey)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		query.SetExclusiveStartKey(startKey)
	}
	return query, nil
}

func (in GetInput) dbQueryInput() (*dynamodb.QueryInput, error) {
	if in.AlbumID == "" || in.ID == "" {
		return nil, errors.New("Missing required fields")
	}
	expr := expression.NewBuilder().
		WithKeyCondition(expression.Key("albumId").Equal(expression.Value(in.AlbumID))).
		WithFilter(expression.Name("id").Equal(expression.Value(in.ID)))

	if len(in.Project) > 0 {
		expr = expr.WithProjection(projectExpression(in.Project))
	}

	exprBuild, err := expr.Build()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &dynamodb.QueryInput{
		KeyConditionExpression:    exprBuild.KeyCondition(),
		FilterExpression:          exprBuild.Filter(),
		ProjectionExpression:      exprBuild.Projection(),
		ExpressionAttributeNames:  exprBuild.Names(),
		ExpressionAttributeValues: exprBuild.Values(),
		TableName:                 aws.String("Photo"),
	}, nil
}

func keyCondExpression(filter FilterInput) expression.KeyConditionBuilder {
	conditions := expression.Key("albumId").Equal(expression.Value(filter.AlbumID))
	if filter.StartDate != "" && filter.EndDate != "" {
		conditions = expression.KeyAnd(conditions,
			expression.KeyBetween(expression.Key("date"),
				expression.Value(filter.StartDate),
				expression.Value(filter.EndDate)))
	}
	return conditions
}

func filterExpression(filter FilterInput) expression.ConditionBuilder {
	isNew := true
	conditions := expression.ConditionBuilder{}
	for _, tag := range filter.Tags {
		conditions = setCondition(conditions, expression.Contains(expression.Name("tags"), tag), isNew)
		isNew = false
	}
	if filter.Description != "" {
		conditions = setCondition(conditions, expression.Contains(expression.Name("description"), filter.Description), isNew)
		isNew = false
	}
	return conditions
}

func projectExpression(projection []string) expression.ProjectionBuilder {
	if !common.Contains(projection, "id") {
		projection = append(projection, "id")
	}
	proj := expression.NamesList(expression.Name(projection[0]))
	for _, name := range projection[1:] {
		proj = expression.AddNames(proj, expression.Name(name))
	}
	return proj
}

func setCondition(origin expression.ConditionBuilder, add expression.ConditionBuilder, isNew bool) expression.ConditionBuilder {
	if isNew {
		return add
	}
	return expression.And(origin, add)
}

func queryOutput(dbOutput *dynamodb.QueryOutput) (*QueryOutput, error) {
	photos := []Photo{}
	err := dynamodbattribute.UnmarshalListOfMaps(dbOutput.Items, &photos)
	if err != nil {
		return nil, err
	}

	encoded, err := encodeLastKey(dbOutput.LastEvaluatedKey)
	if err != nil {
		return nil, err
	}
	return &QueryOutput{
		Items:   photos,
		LastKey: encoded,
	}, nil
}

func encodeLastKey(lastKey map[string]*dynamodb.AttributeValue) (string, error) {
	key := KeyPhoto{}
	err := dynamodbattribute.UnmarshalMap(lastKey, &key)
	if err != nil || key.AlbumID == "" {
		return "", err
	}
	js, err := json.Marshal(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(js), nil
}

func decodeStartKey(encoded string) (map[string]*dynamodb.AttributeValue, error) {
	decoded, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	lastKey := KeyPhoto{}
	err = json.Unmarshal(decoded, &lastKey)
	if err != nil {
		return nil, err
	}
	return dynamodbattribute.MarshalMap(lastKey)
}
