// Code generated by entc, DO NOT EDIT.

package ent

import (
	"caster/ent/profile"
	"caster/utils"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Profile is the model entity for the Profile schema.
type Profile struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Picture holds the value of the "picture" field.
	Picture string `json:"picture,omitempty"`
	// Content holds the value of the "content" field.
	Content *utils.Content `json:"content,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Profile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case profile.FieldContent:
			values[i] = new([]byte)
		case profile.FieldID, profile.FieldEmail, profile.FieldDisplayName, profile.FieldPicture:
			values[i] = new(sql.NullString)
		case profile.FieldCreatedAt, profile.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Profile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Profile fields.
func (pr *Profile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profile.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case profile.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case profile.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case profile.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pr.Email = value.String
			}
		case profile.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				pr.DisplayName = value.String
			}
		case profile.FieldPicture:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field picture", values[i])
			} else if value.Valid {
				pr.Picture = value.String
			}
		case profile.FieldContent:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Content); err != nil {
					return fmt.Errorf("unmarshal field content: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Profile.
// Note that you need to call Profile.Unwrap() before calling this method if this Profile
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Profile) Update() *ProfileUpdateOne {
	return (&ProfileClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Profile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Profile) Unwrap() *Profile {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Profile is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Profile) String() string {
	var builder strings.Builder
	builder.WriteString("Profile(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", email=")
	builder.WriteString(pr.Email)
	builder.WriteString(", display_name=")
	builder.WriteString(pr.DisplayName)
	builder.WriteString(", picture=")
	builder.WriteString(pr.Picture)
	builder.WriteString(", content=")
	builder.WriteString(fmt.Sprintf("%v", pr.Content))
	builder.WriteByte(')')
	return builder.String()
}

// Profiles is a parsable slice of Profile.
type Profiles []*Profile

func (pr Profiles) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
