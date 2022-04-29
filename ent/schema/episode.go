package schema

import (
	"caster/utils"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Episode holds the schema definition for the Episode entity.
type Episode struct {
	ent.Schema
}

// Fields of the Episode.
func (Episode) Fields() []ent.Field {
	return []ent.Field{
		field.Text("id").
			Annotations(&entsql.Annotation{
				Default: "gen_random_ulid()",
			}).
			Unique().
			NotEmpty().
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
		field.Time("updated_at").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),

		field.Text("title").
			NotEmpty(),
		field.Text("summary").
			NotEmpty().
			Optional(),
		field.Text("picture").
			NotEmpty().
			Optional(),
		field.JSON("content", &utils.Content{}).
			Optional(),

		field.Text("show_id").
			NotEmpty(),
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Show.Type).
			Ref("episodes").
			Unique().
			Field("show_id").
			Required(),
	}
}
