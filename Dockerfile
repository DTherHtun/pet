FROM golang:alpine AS builder
RUN apk add --no-cache git && mkdir -p $GOPATH/src/github.com/DTherHtun/pet
WORKDIR $GOPATH/src/github.com/DTherHtun/pet
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/pet .
FROM scratch
COPY --from=builder /go/bin/pet /go/bin/pet
ENTRYPOINT ["/go/bin/pet"]
EXPOSE 8080
