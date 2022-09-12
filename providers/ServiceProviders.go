package providers

import (
	"arman-estimation-service/infrastructures/components"
	"arman-estimation-service/services"
	serviceInterfaces "arman-estimation-service/types/interfaces/services"
)

func ServiceProvider() {

	components.Register[serviceInterfaces.ISegmentService](
		func(params ...any) serviceInterfaces.ISegmentService {
			return &services.SegmentService{
				RedisClient: components.GetRedisClient(),
			}
		},
	)

	components.Register[serviceInterfaces.IEstimationService](
		func(params ...any) serviceInterfaces.IEstimationService {
			return &services.EstimateService{
				RedisClient: components.GetRedisClient(),
			}
		},
	)
}
