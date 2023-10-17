// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/1eedaegon/golang-app-oauth2-system-practice/db/ent/image"
	"github.com/google/uuid"
)

// Image is the model entity for the Image schema.
type Image struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ImageID holds the value of the "image_id" field.
	ImageID uuid.UUID `json:"image_id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID uuid.UUID `json:"tenant_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Image) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case image.FieldID:
			values[i] = new(sql.NullInt64)
		case image.FieldName:
			values[i] = new(sql.NullString)
		case image.FieldCreatedAt, image.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case image.FieldImageID, image.FieldTenantID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Image fields.
func (i *Image) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case image.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case image.FieldImageID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field image_id", values[j])
			} else if value != nil {
				i.ImageID = *value
			}
		case image.FieldTenantID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[j])
			} else if value != nil {
				i.TenantID = *value
			}
		case image.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case image.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case image.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Image.
// This includes values selected through modifiers, order, etc.
func (i *Image) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// Update returns a builder for updating this Image.
// Note that you need to call Image.Unwrap() before calling this method if this Image
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Image) Update() *ImageUpdateOne {
	return NewImageClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Image entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Image) Unwrap() *Image {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Image is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Image) String() string {
	var builder strings.Builder
	builder.WriteString("Image(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("image_id=")
	builder.WriteString(fmt.Sprintf("%v", i.ImageID))
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(fmt.Sprintf("%v", i.TenantID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Images is a parsable slice of Image.
type Images []*Image