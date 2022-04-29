package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
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
		field.Text("username").
			NotEmpty(),
		field.Bool("is_active").
			Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
