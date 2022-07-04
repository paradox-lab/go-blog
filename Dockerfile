# syntax=docker/dockerfile:1

FROM golang:1.18.3-stretch

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
# 安装测试报告工具
RUN go install github.com/vakenbolt/go-test-report@latest

COPY tests ./tests
COPY report ./report

# RUN go build -o /go-blog

#CMD [ "go", "test", "-json", "./tests/...", "|", "go-test-report", "-o", 'report/html/report-$(date "+%Y%m%d%H%M%S").html']
CMD go test -json ./tests/... | go-test-report -o report/html/report-$(date "+%Y%m%d%H%M%S").html