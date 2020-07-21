FROM golang:alpine

WORKDIR /go/src/github.com/santosh/eldiario
COPY . .
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o eldiario .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /go/src/github.com/santosh/eldiario
COPY --from=0 /go/src/github.com/santosh/eldiario .

EXPOSE 8080
CMD ["./eldiario"]
