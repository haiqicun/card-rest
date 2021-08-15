FROM golang:1.16-alpine as builder
COPY go.mod go.sum /go/src/github.com/haiqicun/card-rest/
WORKDIR /go/src/github.com/haiqicun/card-rest/
RUN go mod download
COPY . /go/src/github.com/haiqicun/card-rest/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/card-rest github.com/haiqicun/card-rest/cmd

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/haiqicun/card-rest/build/card-rest /usr/bin/card-rest
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/card-rest"]