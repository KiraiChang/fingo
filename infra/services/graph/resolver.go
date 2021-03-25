package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/KiraiChang/fingo/infra/persistences/ent"
	"github.com/KiraiChang/fingo/infra/services/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
