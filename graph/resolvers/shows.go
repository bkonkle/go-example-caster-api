package resolvers

import (
	"caster/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateShow(ctx context.Context, input model.CreateShowInput) (*model.MutateShowResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateShow(ctx context.Context, id string, input model.UpdateShowInput) (*model.MutateShowResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteShow(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetShow(ctx context.Context, id string) (*model.Show, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetManyShows(ctx context.Context, where *model.ShowCondition, orderBy []model.ShowsOrderBy, page *int, pageSize *int) (*model.ShowsPage, error) {
	panic(fmt.Errorf("not implemented"))
}
