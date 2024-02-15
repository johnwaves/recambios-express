FROM golang:latest AS build

RUN go install github.com/go-task/task/v3/cmd/task@latest

FROM debian:stable-slim AS final

COPY --from=build /usr/local/go/ /usr/local/go/
COPY --from=build /go/bin/task /usr/local/bin/task
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV PATH="/usr/local/go/bin:/usr/local/bin:${PATH}"

RUN mkdir -p /go /go/cache && \
    chmod -R 1777 /go /go/cache

RUN adduser test
USER test

ENV GOPATH=/go
ENV GOCACHE=/go/cache
ENV PATH="$GOPATH/bin:$PATH"

WORKDIR /app/test

ENTRYPOINT ["task", "test"]
