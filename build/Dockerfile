# syntax=docker/dockerfile:1
# releaseing stallion cluster on docker images
# build for containers platforms
# ubuntu operating system
## Build image
FROM golang:1.20-buster AS build

# app work directory
WORKDIR /app

LABEL app=broker
LABEL name=stallion
LABEL type=cluster

# copy mod and sum files
COPY go.mod go.sum ./

# download deps
RUN go mod download

# copy all go files
COPY *.go ./

# creating go application
RUN go build -o /docker-gs

## Deploy
# using debian image
FROM debian:unstable-slim

# maintainer
MAINTAINER official-stallion

LABEL app=broker
LABEL name=stallion
LABEL type=cluster

ENV ST_SERVER_PORT
ENV ST_METRICS_PORT
ENV ST_USER
ENV ST_PASSWORD

# root work directory
WORKDIR /

# copy executable file
COPY --from=build /docker-gs /docker-gs

# executing the runnable file.
CMD ["/docker-gs"]
