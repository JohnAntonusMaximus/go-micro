FROM golang:1.8.1-alpine
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN go get github.com/johnantonusmaximus/go-micro/proto
RUN go get github.com/micro/cli
RUN go get github.com/micro/mdns

ENV SOURCES /go/src/github.com/johnantonusmaximus/go-micro/client/
ADD . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build 

WORKDIR ${SOURCES}
CMD ${SOURCES}client
EXPOSE 8080