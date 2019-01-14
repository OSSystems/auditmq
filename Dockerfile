FROM golang:1.11-alpine3.8 as builder

RUN apk --update --no-cache add git openssh make
RUN go get -u github.com/Masterminds/glide

ADD . /go/src/github.com/OSSystems/auditmq
WORKDIR /go/src/github.com/OSSystems/auditmq

RUN glide install
RUN make

FROM alpine:3.7

COPY --from=builder /go/src/github.com/OSSystems/auditmq/entrypoint.sh /usr/local/bin/
COPY --from=builder /go/src/github.com/OSSystems/auditmq/bin/auditmq /usr/local/bin/

CMD ["/usr/local/bin/entrypoint.sh"]
