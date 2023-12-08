package pipelines

import (
	"context"
	pb "github.com/seidu626/playground/pipelines/proto"
	"github.com/seidu626/playground/pipelines/stage"
)

type Series interface {
	Rank(ctx context.Context, req *pb.Request) (*pb.Request, error)
}

type series struct {
	stages []stage.Stage
}

func NewSeries(stages ...stage.Stage) Series {
	return &series{stages: stages}
}

func (s *series) Rank(ctx context.Context, req *pb.Request) (*pb.Request, error) {
	var err error
	resp := req

	for _, stg := range s.stages {
		resp, err = stg.Rank(ctx, req)
		if err != nil {
			return nil, err
		}
		req = resp
	}

	return resp, nil
}
