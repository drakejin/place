// Code generated by ent, DO NOT EDIT.

package pagereferred

import (
	"time"
)

const (
	// Label holds the string label denoting the pagereferred type in the database.
	Label = "page_referred"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSourceID holds the string denoting the source_id field in the database.
	FieldSourceID = "source_id"
	// FieldTargetID holds the string denoting the target_id field in the database.
	FieldTargetID = "target_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// Table holds the table name of the pagereferred in the database.
	Table = "page_referred"
)

// Columns holds all SQL columns for pagereferred fields.
var Columns = []string{
	FieldID,
	FieldSourceID,
	FieldTargetID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
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
	// CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	CreatedByValidator func(string) error
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	UpdatedByValidator func(string) error
)