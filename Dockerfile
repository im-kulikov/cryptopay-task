# Build image
FROM golang:alpine as builder

COPY . /go/src/github.com/im-kulikov/cryptopay-task

WORKDIR /go/src/github.com/im-kulikov/cryptopay-task

RUN set -x \
    && export CGO_ENABLED="0" \
    && export CGO_CFLAGS="-g -O9" \
    && export CGO_CXXFLAGS="-g -O9" \
    && export CGO_FFLAGS="-g -O9" \
    && export CGO_LDFLAGS="-g -O9" \
    && go build \
        -ldflags "-w -s -extldflags \"-static\"" \
        -gcflags '-m' \
        -gccgoflags '-O9' \
        -v \
        -o /go/bin/task

# Executable image
FROM scratch

ENV GOGC=off

COPY --from=builder /go/bin/task /task
COPY --from=builder /go/src/github.com/im-kulikov/cryptopay-task/data /data

CMD ["/task"]
