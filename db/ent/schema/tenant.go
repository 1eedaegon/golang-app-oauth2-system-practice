package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("tenant_id", uuid.UUID{}).
			Unique().
			Default(uuid.New),
		field.String("name"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edge를 정의할 때 UUID를 PK로 보장하지 않도록 주의한다.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Tenant.Type).
			From("children").
			Unique(),
		edge.To("image", Image.Type),
	}
}
