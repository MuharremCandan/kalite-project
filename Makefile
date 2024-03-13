dev:
	docker compose up -d &&  air --build.cmd "go build -o /tmp/mfk ./main.go" --build.bin "/tmp/mfk"

hot:
	docker build -t kalite-project:latest .

.PHONY: dev