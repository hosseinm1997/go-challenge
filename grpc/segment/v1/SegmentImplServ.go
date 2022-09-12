package v1

import (
	shared "arman-estimation-service/grpc/shared/v1"
	"arman-estimation-service/infrastructures/components"
	services "arman-estimation-service/types/interfaces/services"
	"arman-estimation-service/types/structs"
	"context"
)

type SegmentImplServ struct {
	UnimplementedSegmentServer
}

func (s SegmentImplServ) Push(ctx context.Context, request *PushSegmentRequest) (*PushSegmentResponse, error) {
	service := components.Resolve[services.ISegmentService]()
	err := service.Push(ctx, request.Segment, request.UserId)

	if err == nil {
		return &PushSegmentResponse{Response: &shared.Status{Code: 0}},
			nil
	}

	if err.Category == structs.Categories.BusinessLogic {
		return &PushSegmentResponse{
				Response: &shared.Status{
					Code:    int32(structs.Categories.BusinessLogic),
					Message: err.Err,
				},
			},
			nil
	}

	return &PushSegmentResponse{
		Response: &shared.Status{
			Code:    int32(structs.Categories.Internal),
			Message: "Internal server error",
		},
	}, nil
}
