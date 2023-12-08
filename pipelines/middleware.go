package pipelines

import (
	"context"
	pb "github.com/seidu626/playground/pipelines/proto"
	"github.com/seidu626/playground/pipelines/stage"
	"time"
)

type Middleware func(stage stage.Stage) stage.Stage

func ExampleMiddleware(next stage.Stage) stage.Stage {
	return stage.RankFunc(func(ctx context.Context, req *pb.Request) (*pb.Request, error) {
		// ...
		return next.Rank(ctx, req)
	})
}

func Monitor(next stage.Stage) stage.Stage {
	return stage.RankFunc(func(ctx context.Context, req *pb.Request) (*pb.Request, error) {
		defer func(startedAt time.Time) {
			stageLatencySeconds.With(prometheus.Labels{
				methodLabel: req.Options.Method,
				stageLabel:  stage.Name(next),
			}).Observe(time.Since(startedAt).Seconds())
		}(time.Now())

		return next.Rank(ctx, req)
	})
}
