# This stage sets up tools to build the project, clones it, and creates an executable named `tt`.
FROM golang:alpine AS build

# RUN apk update && apk add --no-cache ncurses-terminfo-base ncurses-libs ncurses

ENV CGO_ENABLED=0

ARG MONGODB_URI=""
RUN echo "Received MONGODB_URI: ${MONGODB_URI}"
ENV MONGODB_URI=${MONGODB_URI}

WORKDIR /app

#RUN apk add --no-cache git \
#    && git clone --depth=1 https://github.com/lemnos/tt . \
#    && go build -o bin/tt src/*.go

COPY . .
RUN go build -o bin/tt src/*.go

# Copies the `tt` executable and sets it as the command to run.
FROM alpine

# RUN apk update && apk add --no-cache ncurses-terminfo-base ncurses-libs ncurses

COPY --from=build  /app/bin/tt /bin/tt

ENTRYPOINT ["tt"]

