package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type ElasticsearchConfig struct {
	Addresses []string
	Username  string
	Password  string
	Index     string
}

type ElasticsearchRepository struct {
	client *elasticsearch.TypedClient
	config ElasticsearchConfig
}

func NewElasticRepository(cfg ElasticsearchConfig) (*ElasticsearchRepository, error) {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: cfg.Addresses,
		Username:  cfg.Username,
		Password:  cfg.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create ES client: %w", err)
	}

	return &ElasticsearchRepository{
		client: client,
		config: cfg,
	}, nil
}

func (r *ElasticsearchRepository) Store(ctx context.Context, data []byte) error {
	_, err := r.client.Index(r.config.Index).
		Request(json.RawMessage(data)).
		Do(ctx)

	return err
}

func (r *ElasticsearchRepository) Query(ctx context.Context, q Query) (QueryResult, error) {
	res, err := r.client.Search().
		Index(r.config.Index).
		Query(&types.Query{
			Bool: &types.BoolQuery{
				Filter: []types.Query{
					{
						Range: map[string]types.RangeQuery{
							"timestamp": {
								Gte: &q.TimeRange.Start,
								Lte: &q.TimeRange.End,
							},
						},
					},
				},
			},
		}).
		From(q.Pagination.Offset).
		Size(q.Pagination.Limit).
		Do(ctx)

	if err != nil {
		return QueryResult{}, err
	}

	records := make([][]byte, 0, len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		records = append(records, hit.Source_)
	}

	return QueryResult{
		Records: records,
		Count:   int(res.Hits.Total.Value),
	}, nil
}