package schema

import (
	"caster/utils"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Show holds the schema definition for the Show entity.
type Show struct {
	ent.Schema
}

// Fields of the Show.
func (Show) Fields() []ent.Field {
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
	}
}

// Edges of the Show.
func (Show) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("episodes", Episode.Type),
	}
}
