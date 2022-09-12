package services

import (
	"arman-estimation-service/types/structs"
	"context"
)

type ISegmentService interface {
	Push(ctx context.Context, segment string, userId uint32) *structs.CustomError
}
