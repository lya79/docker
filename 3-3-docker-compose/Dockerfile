FROM golang:1.14 AS builder

COPY . /go/src/p/3-3-docker-compose 
WORKDIR /go/src/p/3-3-docker-compose
RUN go build -mod vendor -o server .

FROM golang:1.14
COPY --from=builder /go/src/p/3-3-docker-compose/server /go/src/p/3-3-docker-compose/server

WORKDIR /go/src/p/3-3-docker-compose
ENTRYPOINT [ "./server" ]

# WORKDIR /go/src/app

# COPY go.mod .
# COPY go.sum .

# RUN go mod download

# COPY . .

# RUN go build -o server

# CMD ["app"]