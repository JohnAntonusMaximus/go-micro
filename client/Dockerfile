FROM golang:1.8.1-alpine
RUN apk update && apk upgrade && apk add --no-cache bash git openssh

ENV SOURCES /go/src/github.com/johnantonusmaximus/cloud-go/kubernetes/client/
ADD . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build 

WORKDIR ${SOURCES}
CMD ${SOURCES}client
EXPOSE 8080