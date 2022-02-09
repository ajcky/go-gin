FROM golang:1.17.1-alpine3.14 as builder
WORKDIR /app
COPY . /app/
RUN GOOS=linux go build -o main .
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main  /app/

# 执行启动命令
CMD ["/app/main"]
