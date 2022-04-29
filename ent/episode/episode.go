// Code generated by entc, DO NOT EDIT.

package episode

import (
	"time"
)

const (
	// Label holds the string label denoting the episode type in the database.
	Label = "episode"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldSummary holds the string denoting the summary field in the database.
	FieldSummary = "summary"
	// FieldPicture holds the string denoting the picture field in the database.
	FieldPicture = "picture"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldShowID holds the string denoting the show_id field in the database.
	FieldShowID = "show_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the episode in the database.
	Table = "episodes"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "episodes"
	// OwnerInverseTable is the table name for the Show entity.
	// It exists in this package in order to avoid circular dependency with the "show" package.
	OwnerInverseTable = "shows"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "show_id"
)

// Columns holds all SQL columns for episode fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTitle,
	FieldSummary,
	FieldPicture,
	FieldContent,
	FieldShowID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// SummaryValidator is a validator for the "summary" field. It is called by the builders before save.
	SummaryValidator func(string) error
	// PictureValidator is a validator for the "picture" field. It is called by the builders before save.
	PictureValidator func(string) error
	// ShowIDValidator is a validator for the "show_id" field. It is called by the builders before save.
	ShowIDValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
