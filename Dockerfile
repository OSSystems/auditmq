FROM golang:1.10.0-alpine3.7 as builder

RUN apk --update --no-cache add git openssh make
RUN go get -u github.com/Masterminds/glide

ADD . /go/src/github.com/OSSystems/auditmq
WORKDIR /go/src/github.com/OSSystems/auditmq

RUN glide install
RUN make

FROM alpine:3.7

COPY --from=builder /go/src/github.com/OSSystems/auditmq/auditmq /usr/local/auditmq
