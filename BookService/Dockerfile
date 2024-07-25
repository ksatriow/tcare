FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod init book
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o book .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/book .
RUN chmod +x ./book
EXPOSE 8082
CMD ["./book"]