package main

import (
	"log-system/internal/core"
	"log-system/internal/analytics"
	"log-system/internal/messaging"
	"log-system/internal/storage"
)

func main() {
	cfg := loadConfig()

	// Initialize components
	kafkaClient := messaging.NewKafkaClient(cfg.Kafka)
	detector := analytics.NewDetectorChain()
	repo := storage.NewTieredRepository(
		storage.NewElasticRepository(cfg.Elastic),
		storage.NewS3Repository(cfg.S3),
	)

	// Setup processing pipeline
	pipeline := core.NewProcessingPipeline(kafkaClient, detector, repo)

	// Add detection rules
	detector.AddRule(analytics.NewRateLimitRule(1000))
	detector.AddRule(analytics.NewPatternMatchRule(`ERROR`))

	// Start processing
	if err := pipeline.Run(context.Background()); err != nil {
		panic(err)
	}
}