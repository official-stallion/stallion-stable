# syntax=docker/dockerfile:1
## Build image
FROM golang:1.18-buster AS build

# app work directory
WORKDIR /app

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

# root work directory
WORKDIR /

# copy executable file
COPY --from=build /docker-gs /docker-gs

# executing the runnable file.
CMD ["/docker-gs"]