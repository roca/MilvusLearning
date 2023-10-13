package book

import (
	"context"
	"errors"
	"vector-db/collections"

	"slices"

	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

var (
	collectionName = "book"
)
var schema = &entity.Schema{
	CollectionName: collectionName,
	Description:    "Test book search",
	Fields: []*entity.Field{
		{
			Name:       "book_id",
			DataType:   entity.FieldTypeInt64,
			PrimaryKey: true,
			AutoID:     false,
		},
		{
			Name:       "word_count",
			DataType:   entity.FieldTypeInt64,
			PrimaryKey: false,
			AutoID:     false,
		},
		{
			Name:     "book_intro",
			DataType: entity.FieldTypeFloatVector,
			TypeParams: map[string]string{
				"dim": "2",
			},
		},
	},
	EnableDynamicField: true,
}

type Book struct {
	Schema         *entity.Schema
	CollectionName string
}

func (b *Book) CreateCollection() error {
	//defer collections.CloseConnection(collections.MilvusClient)
	schema.CollectionName = b.CollectionName
	b.Schema = schema

	collectionNames, _ := collections.GetCollectionNames(context.Background(), collections.MilvusClient)
	if slices.Contains(collectionNames, collectionName) {
		return errors.New("Book collection already exists!")
	}

	err := collections.CreateCollection(context.Background(), collections.MilvusClient, b.Schema)
	if err != nil {
		return err
	}
	return nil
}
