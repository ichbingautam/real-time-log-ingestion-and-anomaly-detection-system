package storage

import (
	"context"
	"time"
)

type QueryResult struct {
	Records [][]byte
	Count   int
}

type MultiRepository struct {
	providers []Repository
}

func NewMultiRepository(providers ...Repository) *MultiRepository {
	return &MultiRepository{providers: providers}
}

func (mr *MultiRepository) Store(ctx context.Context, data []byte) error {
	for _, repo := range mr.providers {
		if err := repo.Store(ctx, data); err != nil {
			return err
		}
	}
	return nil
}

type Pagination struct {
	Limit  int
	Offset int
}

type TimeRange struct {
	Start time.Time
	End   time.Time
}

type Query struct {
	Filters    map[string]interface{}
	TimeRange  TimeRange
	Pagination Pagination
}


type Repository interface {
	Store(ctx context.Context, data []byte) error
	BatchStore(ctx context.Context, data [][]byte) error
	Query(ctx context.Context, q Query) (QueryResult, error)
}