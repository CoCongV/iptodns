FROM golang:1.15.3-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io,direct
RUN mkdir /build /dist
WORKDIR /build
ADD . /build
RUN go build -o iptodns .
WORKDIR /dist
RUN cp /build/iptodns .
RUN rm -rf /build
CMD [ "/dist/iptodns" ]