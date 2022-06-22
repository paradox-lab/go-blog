# syntax=docker/dockerfile:1

FROM golang:1.18.2-stretch

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod download

COPY tests ./tests

# RUN go build -o /go-blog

CMD [ "go", "test", "-v", "./tests", "|", "go-test-report", "-o", 'report/html/report-$(date "+%Y%m%d%H%M%S").html']