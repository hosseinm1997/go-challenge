package v1

import (
	shared "arman-estimation-service/grpc/shared/v1"
	"arman-estimation-service/infrastructures/components"
	"arman-estimation-service/types/interfaces/services"
	"arman-estimation-service/types/structs"
	"context"
)

type EstimationImplServ struct {
	UnimplementedEstimationServer
}

func (r EstimationImplServ) Count(ctx context.Context, request *EstimationRequest) (*EstimationResponse, error) {
	service := components.Resolve[services.IEstimationService]()
	count, err := service.Estimate(ctx, request.Segment)

	if err == nil {
		return &EstimationResponse{
			Response: &shared.Status{Code: 0},
			Count:    count,
		}, nil
	}

	if err.Category == structs.Categories.BusinessLogic {
		return &EstimationResponse{
			Response: &shared.Status{Code: int32(structs.Categories.BusinessLogic), Message: err.Err},
			Count:    0,
		}, nil
	}

	return &EstimationResponse{
		Response: &shared.Status{Code: int32(structs.Categories.Internal), Message: "Internal server error"},
		Count:    0,
	}, nil
}
