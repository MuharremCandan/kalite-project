# Geliştirme aşaması
FROM golang:latest as builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./kaliteapi

# Üretim aşaması
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/kaliteapi ./kaliteapi
COPY config.yaml /app/
CMD ["/app/kaliteapi"]

