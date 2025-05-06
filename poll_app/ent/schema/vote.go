package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("votes").Unique(),
		// edge.From("option", PollOption.Type).Ref("votes"),
		edge.To("options", PollOption.Type),
	}
}
