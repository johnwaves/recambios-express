FROM golang:latest AS build

RUN go install github.com/go-task/task/v3/cmd/task@latest

FROM debian:stable-slim AS final

COPY --from=build /usr/local/go/ /usr/local/go/
COPY --from=build /go/bin/task /usr/local/bin/task
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV PATH="/usr/local/go/bin:/usr/local/bin:${PATH}"

RUN mkdir /app /cache && chmod -R 777 /cache
ENV GOPATH=/home/test/go
ENV GOCACHE=/home/test/.cache/go-build
ENV PATH="$GOPATH/bin:$PATH"

RUN mkdir -p "$GOPATH" "$GOCACHE" && chmod -R 777 "$GOPATH" "$GOCACHE"

RUN adduser test
USER test

WORKDIR /app/test

ENTRYPOINT ["task", "test"]

