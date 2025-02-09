package core

type ProcessingPipeline struct {
	ingester   messaging.Ingester
	detector   analytics.Detector
	repository storage.Repository
}

func NewPipeline(
	ingester messaging.Ingester,
	detector analytics.Detector,
	repo storage.Repository,
) *ProcessingPipeline {
	return &ProcessingPipeline{
		ingester:   ingester,
		detector:   detector,
		repository: repo,
	}
}

func (p *ProcessingPipeline) Run(ctx context.Context) error {
	return p.ingester.Consume(ctx, func(data []byte) error {
		result := p.detector.Analyze(data)

		if result.IsAnomaly {
			if err := p.repository.Store(ctx, data); err != nil {
				return err
			}
		}
		return nil
	})
}