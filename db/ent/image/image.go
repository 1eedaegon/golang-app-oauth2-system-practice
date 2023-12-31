// Code generated by ent, DO NOT EDIT.

package image

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldImageID holds the string denoting the image_id field in the database.
	FieldImageID = "image_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// Table holds the table name of the image in the database.
	Table = "images"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "images"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "tenant_image"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldImageID,
	FieldName,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTenantID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"tenant_image",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultImageID holds the default value on creation for the "image_id" field.
	DefaultImageID func() uuid.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Image queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByImageID orders the results by the image_id field.
func ByImageID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByTenantField orders the results by tenant field.
func ByTenantField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTenantStep(), sql.OrderByField(field, opts...))
	}
}
func newTenantStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TenantInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TenantTable, TenantColumn),
	)
}
