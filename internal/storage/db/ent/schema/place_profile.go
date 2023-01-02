package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/drakejin/place/internal/storage/db/ent/validate"
)

// PlaceProfile holds the schema definition for the Keyword entity.
type PlaceProfile struct {
	ent.Schema
}

// Fields of the PlaceProfile.
func (PlaceProfile) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),

		field.String("description").
			Annotations(entsql.Annotation{
				Size: 300,
			}).
			Validate(validate.MaxRuneCount(100)).
			Comment("poi 소개 및 설명"),

		field.String("contact_number").
			Annotations(entsql.Annotation{
				Size: 300,
			}).
			Validate(validate.MaxRuneCount(100)).
			Comment("대표전화번호"),
	}
}

// Edges of the PlaceProfile.
func (PlaceProfile) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the PlaceProfile
func (PlaceProfile) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations of the PlaceProfile.
func (PlaceProfile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "page_source",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_0900_ai_ci",
		},
	}
}
