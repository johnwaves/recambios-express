FROM golang:latest AS build

RUN go install github.com/go-task/task/v3/cmd/task@latest

FROM debian:stable-slim AS final

COPY --from=build /usr/local/go/ /usr/local/go/
COPY --from=build /go/bin/task /usr/local/bin/task

ENV PATH="/usr/local/go/bin:/usr/local/bin:${PATH}"

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN adduser -u 1001 test
USER test

WORKDIR /app/test

ENTRYPOINT ["task", "test"]
