package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Diary struct {
	ent.Schema
}

// Fields of the Diary.
func (Diary) Fields() []ent.Field {
	return []ent.Field{
		field.String("emotion"),
		field.String("condition"),
		field.String("texts"),
		field.Int("user_id"),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Diary.
func (Diary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("diaries").
			Unique().
			Field("user_id").Required(),
	}
}
