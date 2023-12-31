// Code generated by ent, DO NOT EDIT.

package tenant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the tenant type in the database.
	Label = "tenant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeImage holds the string denoting the image edge name in mutations.
	EdgeImage = "image"
	// Table holds the table name of the tenant in the database.
	Table = "tenants"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "tenants"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "tenant_parent"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "tenants"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "tenant_parent"
	// ImageTable is the table that holds the image relation/edge.
	ImageTable = "images"
	// ImageInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImageInverseTable = "images"
	// ImageColumn is the table column denoting the image relation/edge.
	ImageColumn = "tenant_image"
)

// Columns holds all SQL columns for tenant fields.
var Columns = []string{
	FieldID,
	FieldTenantID,
	FieldName,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tenants"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"tenant_parent",
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
	// DefaultTenantID holds the default value on creation for the "tenant_id" field.
	DefaultTenantID func() uuid.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Tenant queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
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

// ByChildrenField orders the results by children field.
func ByChildrenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), sql.OrderByField(field, opts...))
	}
}

// ByParentCount orders the results by parent count.
func ByParentCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newParentStep(), opts...)
	}
}

// ByParent orders the results by parent terms.
func ByParent(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByImageCount orders the results by image count.
func ByImageCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newImageStep(), opts...)
	}
}

// ByImage orders the results by image terms.
func ByImage(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImageStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ChildrenTable, ChildrenColumn),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ParentTable, ParentColumn),
	)
}
func newImageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ImageTable, ImageColumn),
	)
}
