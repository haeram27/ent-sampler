// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Node struct {
	ent.Schema
}

func (Node) Fields() []ent.Field {
	return []ent.Field{field.String("id").StorageKey("node_name"), field.String("timestamp"), field.String("pod_info")}
}
func (Node) Edges() []ent.Edge {
	return nil
}
func (Node) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "node"}}
}
