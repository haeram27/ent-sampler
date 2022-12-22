package schema

import "entgo.io/ent"

// Bugs holds the schema definition for the Bugs entity.
type Bugs struct {
	ent.Schema
}

// Fields of the Bugs.
func (Bugs) Fields() []ent.Field {
	return nil
}

// Edges of the Bugs.
func (Bugs) Edges() []ent.Edge {
	return nil
}
