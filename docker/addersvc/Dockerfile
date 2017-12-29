FROM golang:1.9.2 as builder
COPY . /go/src/github.com/dandeliondeathray/goosedev
RUN go get github.com/dandeliondeathray/goosedev/...
RUN go install github.com/dandeliondeathray/goosedev/cmd/addersvc

FROM golang:1.9.2
COPY --from=builder /go/bin/addersvc /go/bin/addersvc
CMD ["/go/bin/addersvc"]