# Etimation service part of "Arman solution" go challenge

This service is responsible for saving user into segment for future estimations.

## Quick start
To quickly jump into the main logic, go to following links:

### Main Endpoints:
- [Save user segments](https://github.com/hosseinm1997/go-challenge/blob/main/services/SegmentService.go)
- [Estimation logic](https://github.com/hosseinm1997/go-challenge/blob/main/services/EstimateService.go)

## Overview

### Approach

For overcomming this challenge I Use [Redis HyperLogLog feature](https://redis.com/redis-best-practices/counting/hyperloglog). 
It uses a probablistic algorithm to find estimated number of unique elements. It is really much faster than normal counting.

For microservices internal communations, gRPC used for faster and cleaner response.


### Solution



### Framework
This service was made based on a simple framework made by myself (in a limited time). I'm not interested in `reinvent the wheel` myself!! My idea behind this is to dig into the Go language deeper. It has following features:

- IoC implemented using service container, created by new `generic` feature of go 1.18. [see ServiceContainer.go](https://github.com/hosseinm1997/credit-service/blob/main/infrastructures/ServiceContainer.go)
- Routing system using middlewares and contextes. [see RoutingSystem.go](https://github.com/hosseinm1997/credit-service/blob/main/infrastructures/RoutingSystem.go)
- Easy exception handling with `Respond()` helper function. [see ResponseFormatter.go](https://github.com/hosseinm1997/credit-service/blob/main/http/middlewares/ResponseFormatter.go), [see an example](https://github.com/hosseinm1997/credit-service/blob/ab1eda279aa9e2a4d02b4d752e09de0e0f3da42f/http/endpoints/SpendCodeEndpoint.go#L71)
- Handling env variables
- Service and repository pattern considered

### Packages
Direct packages used:

- Viper for managing env variables

<br/>
<br/>
<br/>
