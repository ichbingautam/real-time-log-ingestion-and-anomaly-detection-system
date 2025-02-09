package core

import (
	"context"
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/analytics"
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/messaging"
	"github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system/storage"
	"time"
)

type Engine struct {
	ingester    messaging.Ingester
	detector    *analytics.DetectorChain
	storage     storage.Repository
	alertWriter messaging.Egester
}

func NewEngine(
	ingester messaging.Ingester,
	detector *analytics.DetectorChain,
	storage storage.Repository,
	alertWriter messaging.Egester,
) *Engine {
	return &Engine{
		ingester:    ingester,
		detector:    detector,
		storage:     storage,
		alertWriter: alertWriter,
	}
}

func (e *Engine) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err := e.ingester.Consume(ctx, e.processMessage)
			if err != nil {
				return err
			}
		}
	}
}

func (e *Engine) processMessage(msg []byte) error {
	start := time.Now()

	// Detect anomalies
	result := e.detector.Analyze(msg)

	// Store all messages
	if err := e.storage.Store(context.Background(), msg); err != nil {
		return err
	}

	// Write alerts
	if result.IsAnomaly {
		if err := e.alertWriter.Produce(context.Background(), [][]byte{msg}); err != nil {
			return err
		}
	}

	metrics.ProcessingLatency.Observe(time.Since(start).Seconds())
	return nil
}