all: build

build:
	go build .

mockgen:
	mockgen -package mocks github.com/NeowayLabs/wabbit Delivery,Channel > mocks/wabbit.go
	mockgen -package mocks -source storage/storage.go Storage > mocks/storage.go
	mockgen -package mocks -source storage/service.go Service > mocks/service.go
	mockgen -package mocks -source publisher/publisher.go Publisher > mocks/publisher.go
