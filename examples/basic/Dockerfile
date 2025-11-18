FROM golang:1.22-alpine3.20

WORKDIR /go/src/github.com/KhalisaGhaizani/otelchi
COPY . .

WORKDIR /go/src/github.com/KhalisaGhaizani/otelchi/examples/basic
RUN go mod download -x
RUN go build -o server

ENTRYPOINT ["./server"]
