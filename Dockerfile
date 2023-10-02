# Building
FROM golang:1.21.1-bookworm AS Builder

WORKDIR /project
COPY src .
RUN make build

# Deploying
FROM alpine:latest
RUN apk add gcompat

WORKDIR /project
COPY --from=Builder /project/main .
EXPOSE 4001
