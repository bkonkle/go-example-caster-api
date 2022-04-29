package resolvers

import (
	"caster/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) GetCurrentUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) GetOrCreateCurrentUser(ctx context.Context, input model.CreateUserInput) (*model.MutateUserResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCurrentUser(ctx context.Context, input model.UpdateUserInput) (*model.MutateUserResult, error) {
	panic(fmt.Errorf("not implemented"))
}
