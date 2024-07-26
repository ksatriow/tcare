FROM golang:latest AS builder
WORKDIR /app
COPY . .
# RUN go mod init apigateway
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o apigateway .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/apigateway .
RUN chmod +x ./apigateway
EXPOSE 8080
CMD ["./apigateway"]