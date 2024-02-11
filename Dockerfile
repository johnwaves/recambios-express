FROM bitnami/golang:latest

RUN adduser --disabled-password -u 1001 test
USER test

WORKDIR /app

ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org

RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY go.mod go.sum Taskfile.yml ./

RUN go mod download

WORKDIR /app/test

ENTRYPOINT ["task", "test"]
