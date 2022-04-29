package resolvers

import (
	"caster/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateProfile(ctx context.Context, input model.CreateProfileInput) (*model.MutateProfileResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, id string, input model.UpdateProfileInput) (*model.MutateProfileResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProfile(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProfile(ctx context.Context, id string) (*model.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetManyProfiles(ctx context.Context, where *model.ProfileCondition, orderBy []model.ProfilesOrderBy, page *int, pageSize *int) (*model.ProfilesPage, error) {
	panic(fmt.Errorf("not implemented"))
}
