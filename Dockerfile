FROM golang:1.10 as builder

RUN mkdir -p /build
ADD ./app1 /build/
WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go get -d github.com/gorilla/mux
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app1 .

# Stage 2
FROM golang:1.10 as build1

COPY --from=builder /build/ /app/
WORKDIR /app
EXPOSE 9000
CMD [ "./app1" ]