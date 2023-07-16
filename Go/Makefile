NAME=main

.PHONY=build
build:
	go build -o ${NAME} main.go

.PHONY=run
run:
	./${NAME} ${ARGS}

.PHONY=clean
clean:
	go clean

.PHONY=start
start:build run
