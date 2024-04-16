package main

import (
	"context"
	"fmt"
	"go-backend-test/pkg/api"
	"go-backend-test/pkg/config"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server, err := api.NewServer(&config)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	go recordMetrics()
	go collectMetrics()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := server.StartServer(ctx); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func recordMetrics() {
	go func() {
		for {
			requestTotal.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage_percentage",
			Help: "Current CPU usage percentage of the API.",
		},
		[]string{"core"},
	)

	ramUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "ram_usage_bytes",
			Help: "Current RAM usage in bytes of the API.",
		},
	)
	requestTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total number of requests to the API.",
		},
	)
)

func init() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(ramUsage)
	//prometheus.MustRegister(opsProcessed)
}

func collectMetrics() {
	for {
		// CPU kullanımını al
		percentages, err := cpu.PercentWithContext(context.Background(), time.Second, false)
		if err != nil {
			log.Println("Error getting CPU usage:", err)
		} else {
			for i, usage := range percentages {
				cpuUsage.WithLabelValues(fmt.Sprintf("cpu%d", i)).Set(usage)
			}
		}

		// RAM kullanımını al
		memInfo, err := mem.VirtualMemoryWithContext(context.Background())
		if err != nil {
			log.Println("Error getting RAM usage:", err)
		} else {
			ramUsage.Set(float64(memInfo.Used))
		}

		time.Sleep(time.Second)
	}
}
