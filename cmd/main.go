package main

import (
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/config"
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/core"
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/messaging"
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/storage"
)

func main() {
	cfg, err := config.Load("configs/production.yaml")
	if err != nil {
		panic(err)
	}

	// Initialize storage
	esRepo, _ := storage.NewElasticRepository(cfg.Elastic)
	s3Repo, _ := storage.NewS3Repository(cfg.S3)

	// Initialize Kafka
	consumer, _ := messaging.NewKafkaConsumer(cfg.Kafka)
	producer, _ := messaging.NewKafkaProducer(cfg.Kafka)

	// Create processing engine
	engine := core.NewEngine(
		consumer,
		analytics.NewDetector(),
		storage.NewTieredRepository(esRepo, s3Repo),
		producer,
	)

	// Start processing
	if err := engine.Start(context.Background()); err != nil {
		panic(err)
	}
}