package services

import (
	"arman-estimation-service/types/structs"
	"context"
)

type IEstimationService interface {
	Estimate(ctx context.Context, segment string) (uint64, *structs.CustomError)
}
