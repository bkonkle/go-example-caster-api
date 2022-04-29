package schema

import (
	"caster/utils"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
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
		field.Text("email").
			NotEmpty(),
		field.Text("display_name").
			NotEmpty().
			Optional(),
		field.Text("picture").
			NotEmpty().
			Optional(),
		field.JSON("content", &utils.Content{}).
			Optional(),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return nil
}
