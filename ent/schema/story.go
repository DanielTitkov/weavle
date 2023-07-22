package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Story struct {
	ent.Schema
}

// Fields of the Story.
func (Story) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Enum("status").Values("open", "closed", "waiting").Default("open"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").UpdateDefault(time.Now),
	}
}

// Edges of the Story.
func (Story) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sentences", Sentence.Type),
	}
}
