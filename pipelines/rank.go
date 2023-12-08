package pipelines

import (
	"context"
	pb "github.com/seidu626/playground/pipelines/proto"
	"github.com/seidu626/playground/pipelines/stage"
)

func Pipeline(d *service.Dependencies) stage.Stage {
	return stage.Series(
		stage.FetchSubscriptions(d.SubscriptionService),
		stage.FetchPosts(d.Cache),
		stage.FilterPostsForCalifornia(),
		stage.ShufflePosts(0.2),
	)
}

func Pipeline(d *service.Dependencies) stage.Stage {
	return stage.Series(
		stage.FetchSubscriptions(d.SubscriptionService),
		stage.FetchPosts(d.Cache),
		stage.RankFunc(func(context.Context, *pb.Request) (*pb.Request, error) {
			if req.Context.Features["geo_region"] == "CA" {
				// ...
			}
			return req, nil
		}),
		stage.ShufflePosts(0.2),
	)
}

func (s *filterRecentlyViewedPosts) Rank(ctx context.Context, req *pb.Request) (*pb.Request, error) {
	seen := req.Context.Features["recently_viewed_post_ids"].GetAsBoolMap()

	var n int
	for _, candidate := range req.Candidates {
		if !seen[candidate.Id] {
			req.Candidate
			s[n] = candidate
			n++
		}
	}
	req.Candidates = req.Candidates[:n] // in-place filtering
	return req, nil
}

func (s *parallel) Rank(ctx context.Context, req *pb.Request) (*pb.Request, error) {
	resps := make([]*pb.Request, len(s.stages))
	g, groupCtx := errgroup.WithContext(ctx)

	for i := range s.stages {
		i := i
		g.Go(func() error {
			defer log.CapturePanic(groupCtx)
			resp, err := s.stages[i].Rank(groupCtx, pb.Copy(req))
			if err != nil {
				return err
			}
			resps[i] = resp
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return s.merge(ctx, req, resps...)
}
