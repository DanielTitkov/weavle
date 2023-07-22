package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Sentence holds the schema definition for the Sentence entity.
type Sentence struct {
	ent.Schema
}

// Fields of the Sentence.
func (Sentence) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("text").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Int("order").Positive(),
	}
}

// Edges of the Sentence.
func (Sentence) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("story", Story.Type).
			Ref("sentences").
			Required().Unique(),
		edge.From("author", User.Type).
			Ref("sentences").
			Required(),
	}
}
