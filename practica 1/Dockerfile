FROM golang:1.21.2-bullseye as builder

COPY . $GOPATH/src/TQS
WORKDIR $GOPATH/src/TQS
RUN ./build.sh

FROM debian:10.13
RUN mkdir -p /home/memoryServer
WORKDIR /home/memoryServer
COPY --from=builder /go/src/TQS/build/server /home/memoryServer
CMD [ "./server" ]
