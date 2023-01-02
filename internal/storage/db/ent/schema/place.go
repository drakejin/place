package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/drakejin/place/internal/storage/db/ent/validate"
)

// Place holds the schema definition for the Keyword entity.
type Place struct {
	ent.Schema
}

// Fields of the Place.
// Place 의 메인정보
func (Place) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),

		field.String("name").
			Annotations(entsql.Annotation{
				Size: 200,
			}).
			Validate(validate.MaxRuneCount(200)).
			Comment("name"),

		field.Time("created_at").
			Default(time.Now().UTC).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Comment("first indexed time"),

		field.String("created_by").
			Annotations(entsql.Annotation{
				Size: 300,
			}).
			Validate(validate.MaxRuneCount(100)).
			Comment("first indexed time by which system"),

		field.Time("updated_at").
			Default(time.Now().UTC).
			UpdateDefault(time.Now().UTC).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
			}).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Comment("modified time"),

		field.String("updated_by").
			Annotations(entsql.Annotation{
				Size: 300,
			}).
			Validate(validate.MaxRuneCount(100)).
			Comment("modified by which system"),
		field.Enum("status").
			Values("IDLE", "DELETE", "WAIT", "BLOCK").
			Default("VALID").
			Comment("row의 데이터 상태를 의미합니다 IDLE => 사용가능, DELETE => 삭제, WAIT => 대기, BLOCK => 차단, 정지 상태"),

		field.Enum("type").
			Values("NORMAL").
			Default("NORMAL").
			Comment("row의 데이터 타입을 의미합니다."),

		field.Enum("country_code").
			Values("kr").
			Default("kr").
			Comment("row의 데이터 타입을 의미합니다."),

		field.String("geo_point").
			SchemaType(map[string]string{
				dialect.MySQL: "TEXT",
			}).
			Comment("geo_point = { latitude: number, longitude: number }"),

		// https://en.wikipedia.org/wiki/GeoJSON
		field.String("geo_coordinates").
			SchemaType(map[string]string{
				dialect.MySQL: "TEXT",
			}).
			Comment("geo_coordinates = {}"),

		field.String("geo_h3").
			Annotations(entsql.Annotation{
				Size: 300,
			}).
			Validate(validate.MaxRuneCount(100)).
			Comment("modified by which system"),

		field.String("city").
			Annotations(entsql.Annotation{
				Size: 300,
			}).
			Validate(validate.MaxRuneCount(100)).
			Comment("city"),

		field.String("postal_code").
			Annotations(entsql.Annotation{
				Size: 300,
			}).
			Validate(validate.MaxRuneCount(100)).
			Comment("postal_code / zipcode"),

		field.String("state").
			Annotations(entsql.Annotation{
				Size: 300,
			}).
			Validate(validate.MaxRuneCount(100)).
			Comment("state / province / region"),

		field.String("street_address").
			Annotations(entsql.Annotation{
				Size: 600,
			}).
			Validate(validate.MaxRuneCount(200)).
			Comment("street address 도로명 주소"),

		field.String("building_address").
			Annotations(entsql.Annotation{
				Size: 600,
			}).
			Validate(validate.MaxRuneCount(200)).
			Comment("빌딩 및 건물에서의 정보"),
	}
}

// Edges of the Place.
func (Place) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the Place
func (Place) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").StorageKey("ux_name").Unique(),
	}
}

// Annotations of the Place.
func (Place) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "place",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_0900_ai_ci",
		},
	}
}
