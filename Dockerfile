FROM golang:1.15.3-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io,direct
RUN mkdir /iptodns
WORKDIR /iptodns
ADD . /iptodns
RUN go build -o iptodns .
CMD [ "/iptodns/iptodns" ]