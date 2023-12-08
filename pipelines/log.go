package pipelines

import (
	"context"
	pb "github.com/seidu626/playground/pipelines/proto"
	"github.com/seidu626/playground/pipelines/stage"
	"log"
)

func Log(next stage.Stage) stage.Stage {
	return stage.RankFunc(func(ctx context.Context, req *pb.Request) (resp *pb.Request, err error) {
		defer func() {
			if err != nil {
				log.Errorw(
					"stage failed",
					"error", err,
					"request", req.JSON(),
					"response", resp.JSON(),
					"stage", stage.Name(stage),
				)
			}
		}()
		return Rank(ctx, req)
	})
}
