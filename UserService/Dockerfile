FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod init user
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o user .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/user .
RUN chmod +x ./user
EXPOSE 8081
CMD ["./user"]