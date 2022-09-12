# Estimation Service, "Arman solution" go challenge

This service is responsible for saving user into segments to estimate number of users for each segmenet.

## Quick start
To quickly jump into the main logic, go to following links:

### Main Endpoints:
- [Save user segments into redis](https://github.com/hosseinm1997/go-challenge/blob/main/services/SegmentService.go)
- [Estimation logic](https://github.com/hosseinm1997/go-challenge/blob/main/services/EstimateService.go)

## Overview

### Approach

For overcomming this challenge I Use Redis [HyperLogLog](https://redis.com/redis-best-practices/counting/hyperloglog) feature.
It uses a probablistic algorithm to find estimated number of unique elements. It is really much faster than normal counting.

For microservices internal communations, gRPC used for faster and cleaner response.


### Solution



### Framework
This service was made based on a simple framework made by myself (in a limited time). I'm not interested in `reinvent the wheel` myself!! My idea behind this is to dig into the Go language deeper. It has following features:

- IoC implemented using service container, created by new `generic` feature of go 1.18. [see ServiceContainer.go](https://github.com/hosseinm1997/credit-service/blob/main/infrastructures/ServiceContainer.go)
- Service pattern considered

### Packages
Direct packages used:

- Viper for managing env variables
- Go redis package
- gRPC package for golang
- Go protobuf

<br/>
<br/>
<br/>
