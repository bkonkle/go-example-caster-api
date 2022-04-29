package resolvers

import (
	"caster/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateEpisode(ctx context.Context, input model.CreateEpisodeInput) (*model.MutateEpisodeResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, id string, input model.UpdateEpisodeInput) (*model.MutateEpisodeResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEpisode(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetEpisode(ctx context.Context, id string) (*model.Episode, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetManyEpisodes(ctx context.Context, where *model.EpisodeCondition, orderBy []model.EpisodesOrderBy, page *int, pageSize *int) (*model.EpisodesPage, error) {
	panic(fmt.Errorf("not implemented"))
}
