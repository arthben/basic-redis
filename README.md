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

## Run the example
### Basic
```console
cd basic
go run basic.go
```

### Locking
```console
cd locking
# open terminal window and run ( this terminal will get message success lock):
go run locking.go

# open New terminal window and run ( this terminal will get message failed lock):
go run locking.go
```

### PubSub
```console
cd pubsub
# open terminal window and run ( Subscriber terminal )
go run pubsub.go --mode=subscriber

# open NEW terminal window and run ( Publisher terminal )
go run pubsub.go --mode=publisher

# when the publisher success, subscriber terminal will receive message that published
```