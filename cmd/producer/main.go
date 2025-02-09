package main

import (
	"context"
	"log"
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/config"
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/messaging"
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize producer
	producer, err := messaging.NewKafkaProducer(cfg.Kafka)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	// Setup graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Generate logs
	go generateLogs(ctx, producer, cfg.Producer)

	<-sigChan
	log.Println("Shutting down producer...")
}

func generateLogs(ctx context.Context, producer messaging.Egester, cfg config.ProducerConfig) {
	ticker := time.NewTicker(cfg.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			messages := utils.GenerateLogBatch(cfg.BatchSize)
			if err := producer.Produce(ctx, messages); err != nil {
				log.Printf("Failed to produce messages: %v", err)
			}
		}
	}
}