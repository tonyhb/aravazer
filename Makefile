all: bin

bin: 
	docker run --rm \
		-v $(shell pwd -P):/go/src/seneca \
		-v $(shell pwd -P)/bin:/go/bin \
		golang:1.7.1-alpine \
		go install seneca
