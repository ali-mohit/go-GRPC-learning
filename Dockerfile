FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=offj
ADD ./src /build
ADD ./go.mod /build/go.mod
RUN cd /build && go mod download
RUN go build -o ./go-GRPC-learning

CMD ["/build/go-GRPC-learning"]