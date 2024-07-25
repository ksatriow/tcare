FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod init bill
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o bill .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bill .
RUN chmod +x ./bill
EXPOSE 8083
CMD ["./bill"]