# basic-redis

## Start Redis instance with Docker
```console
docker pull redis # get latest redis image from docker
docker run --name redis-local-instance -p 6379:6379 -d redis # listen redis on port 6379
docker ps | grep redis-local-instance # check if docker is ready
```

## Get the redis for using in import statetment
```console
go get "github.com/go-redis/redis/v9"
```