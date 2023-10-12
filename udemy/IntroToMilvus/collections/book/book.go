package book

import "github.com/milvus-io/milvus-sdk-go/v2/entity"

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
