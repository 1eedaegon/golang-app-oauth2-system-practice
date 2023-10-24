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
		field.String("name"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
		field.UUID("tenant_id", uuid.UUID{}).
			Optional(),
	}
}

// Edge를 정의할 때 UUID를 PK로 보장하지 않도록 주의한다.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("tenant", Tenant.Type).
			Ref("image").
			Unique(),
	}
}
