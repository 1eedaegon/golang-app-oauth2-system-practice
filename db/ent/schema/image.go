package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("image_id", uuid.UUID{}).
			Unique().
			Default(uuid.New),
		field.UUID("tenant_id", uuid.UUID{}),
		field.String("name"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

func (Image) Edge() []ent.Edge {
	return []ent.Edge{
		edge.
			From("tenant", Tenant.Type).
			Ref("resource"),
	}
}
