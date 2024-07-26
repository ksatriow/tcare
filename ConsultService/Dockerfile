FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod init consult
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o consult .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/consult .
RUN chmod +x ./consult
EXPOSE 8084
CMD ["./consult"]