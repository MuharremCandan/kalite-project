# Geliştirme aşaması
FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o kaliteproject

# Üretim aşaması
FROM alpine:latest
COPY --from=builder /app/kaliteproject /usr/local/bin/kaliteproject
CMD ["kaliteproject"]
