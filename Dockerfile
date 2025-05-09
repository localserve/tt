# This stage sets up tools to build the project, clones it, and creates an executable named `tt`.
FROM golang:alpine AS build

ENV CGO_ENABLED=0

WORKDIR /app

COPY . .
RUN go build -o bin/tt src/*.go

# Copies the `tt` executable and sets it as the command to run.
FROM alpine

COPY --from=build  /app/bin/tt /bin/tt

ENTRYPOINT ["tt"]
