dev:
	docker compose up -d &&  air --build.cmd "go build -o /tmp/mfk ./main.go" --build.bin "/tmp/mfk"

dev2:
	docker compose up -d && SERVICE_NAME=goGinApp INSECURE_MODE=true OTEL_EXPORTER_OTLP_ENDPOINT=localhost:8089 go run main.go

hot:
	docker build -t kalite-project:latest .

.PHONY: dev