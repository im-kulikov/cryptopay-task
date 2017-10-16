# Build image
FROM golang:latest as builder

ADD . /go/src/task

RUN set -x \
    && cd /go/src/task \
    && go build -o /go/bin/task \
        -ldflags "-extldflags \"-static\"" \
        -gcflags '-m' \
    && chmod 1777 /go/bin/task

# Executable image
FROM scratch

COPY --from=builder /go/bin/task /task
ADD ./data /data

CMD ["/task"]
