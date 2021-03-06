package resolvers

import (
	"caster/graph/generated"
)

type Resolver struct{}

func NewResolver() generated.Config {
	r := Resolver{}

	return generated.Config{
		Resolvers: &r,
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
