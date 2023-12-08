package pipelines

import (
	"context"
	pb "github.com/seidu626/playground/pipelines/proto"
)

type fetchPopularPosts struct {
	cache *store.PostCache
}

func FetchPopularPosts(cache *store.PostCache) *fetchPopularPosts {
	return &fetchPopularPosts{cache: cache}
}

func (s *fetchPopularPosts) Rank(ctx context.Context, req *pb.Request) (*pb.Request, error) {
	postIDs, err := s.cache.FetchPopularPostIDs(ctx)
	if err != nil {
		return nil, err
	}

	for _, id := range postIDs {
		req.Candidates = append(req.Candidates, pb.NewCandidate(postID))
	}

	return req, nil
}
