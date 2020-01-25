FROM golang:alpine as builder

RUN apk update && apk add --no-cache git
RUN mkdir -p /build
ADD . /build
WORKDIR /build
RUN go get -d -v ./...
RUN go install -v -/..
ENV GO111MODULE=on
ENV PORT=9000
RUN go build app1.go
RUN go build app2.go

# Stage 2

FROM alpine as build1
WORKDIR /build
COPY --from=builder /build/app1 /app1/
ENTRYPOINT ./app1
 
 # Stage 3

FROM alpine as build2
WORKDIR /build
COPY --from=builder /build/app2 /app2/
ENTRYPOINT ./app2
