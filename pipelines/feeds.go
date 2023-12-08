package pipelines

import (
	"context"
	"github.com/seidu626/playground/pipelines/stage"
)

func (s *Service) GetPopularFeed(ctx context.Context, req *pb.FeedRequest) (*pb.PopularFeed, error) {
	posts, err := s.fetchPopularAndVideoPosts(ctx)
	if err != nil {
		return nil, err
	}
	posts = s.filterPosts(posts)
	posts, scores, err := s.model.ScorePosts(ctx, req.UserID, posts)
	if err != nil {
		return nil, err
	}
	posts = s.sortPosts(posts, scores)
	return pb.NewPopularFeed(posts), nil
}

func PopularFeed(d *service.Dependencies) stage.Stage {
	return stage.Series(
		stage.Parallel(merger.MergeCandidates,
			stage.FetchPopularPosts(d.PostCache),
			stage.FetchVideoPosts(d.PostCache),
			stage.FetchImagePosts(d.PostCache),
		),
		stage.FetchRecentlyViewedPosts(d.UserPostViews),
		stage.FilterRecentlyViewedPosts(),
		stage.ScoreCandidates(d.RankingModel),
		stage.SortCandidates(),
	)
}
