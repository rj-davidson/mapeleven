package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PSGames holds the schema definition for the PSGames entity.
type PSGames struct {
	ent.Schema
}

// Fields of the PSGames.
func (PSGames) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Appearances"),
		field.Int("Lineups"),
		field.Int("Minutes"),
		field.Int("Number").Default(0),
		field.String("Position").Default(""),
		field.String("Rating").Default(""),
		field.Bool("Captain").Default(false),
		field.Time("lastUpdated").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}

// Edges of the PSGames.
func (PSGames) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playerStats", PlayerStats.Type).Ref("psGames").Unique(),
	}
}
