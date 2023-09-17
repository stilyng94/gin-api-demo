package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Album holds the schema definition for the Album entity.
type Album struct {
	ent.Schema
}

// Fields of the Album.
func (Album) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("artist"),
		field.Int("price").Min(1).Positive(),
	}
}

// Edges of the Album.
func (Album) Edges() []ent.Edge {
	return nil
}

func (Album) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title", "artist").
			Unique(),
	}
}

func (Album) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}

}
