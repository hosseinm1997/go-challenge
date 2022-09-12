# Estimation Service, "Arman solution" go challenge

This service is responsible for saving user into segments to estimate number of users for each segmenet.

## Quick start
To quickly jump into the main logic, go to following links:

### Main Endpoints:
- [Save user segments into redis](https://github.com/hosseinm1997/go-challenge/blob/main/services/SegmentService.go)
- [Estimation logic](https://github.com/hosseinm1997/go-challenge/blob/main/services/EstimateService.go)

## Overview

### Approach

For overcomming this challenge I Use Redis [**HyperLogLog**](https://redis.com/redis-best-practices/counting/hyperloglog) feature.
It uses a probablistic algorithm to find estimated number of unique elements. It is really much faster than normal counting.

For microservices internal communations, gRPC used for faster and cleaner response.


### Solution
The solution is based on hyperloglog logic. Hyperloglog has three command for adding, merging and counting the number of elements.

`PFADD key element [element ...]`: It will add the element(s) to the key and increase its internal counter.

`PFCOUNT key [key ...]:` It will estimate the number of unique element(s) added to the key.

`PFMERGE destkey sourcekey [sourcekey]`: It will merge other HYLL elements into a new destkey uniquely.

Each time any request comes into the service to add user into segment, It will be saved into a redis key.
For example there is a request to add a user 123456 to "sport" segment, and current date is 2022-09-12. So the redis key will be `sport:20220912`.
and this command runs against redis:

`PFADD sport:20220912 123456`

If it is the first user that added to the sport segment for current today, It sets the expiration of two weeks. for other request it just add to the list, and TTL won't be updated.


The reason behind this decision is that, we can not set TTL for each user individually. So we divide users into day-to-day segmenets. Therefore our estimation has a fault about 24 hours. It can be reduced by making the redis keys more accurate, and add hours (or even minutes) to keys and set expiration based on them. But It reduce PFMERGE performance and increase its overhead.

Finally for estimation of users in a segment, we use `PFMERGE` command and aggregate last 14 days created redis keys of current segment.
for example: for sport segment we use this command:


`PFMERGE sport sport:20220912 sport:20220911 sport:20220910 sport:20220909 sport:20220908 ...`

This command will store unique users that added to the `sport` segment for last 14 days. Afterward, we use this command to estimate the users:

`PFCOUNT sport`.

To improve performance, we merge keys only every one hour. Each merged key has 1 hour time to live.

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
