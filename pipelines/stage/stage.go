package stage

import (
	"context"
	pb "github.com/seidu626/playground/pipelines/proto"
)

type Stage interface {
	Rank(ctx context.Context, req *pb.Request) (*pb.Request, error)
}

type RankFunc func(context.Context, *pb.Request) (*pb.Request, error)

func (f RankFunc) Rank(ctx context.Context, req *pb.Request) (*pb.Request, error) {
	return f(ctx, req)
}
