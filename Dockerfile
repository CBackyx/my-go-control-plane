FROM golang:1.13.6-stretch

WORKDIR /go-control-plane

# Fetch and preserve module dependencies
ENV GOPROXY=https://goproxy.cn
# ENV TEMP_HANDLER_HOST=127.0.0.1
# ENV TEMP_HANDLER_HOST_PORT=8081
# ENV MIDDLEWARE_HOST=127.0.0.1
# ENV MIDDLEWARE_HOST_PORT=8091

COPY . ./
RUN go mod download

RUN go build -v -race -o main internal/example/main/main.go

CMD ["./main"]