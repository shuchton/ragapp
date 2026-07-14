package vector

import "context"

type Document struct {
	ID        string
	Content   string
	Metadata  map[string]string
	Embedding []float32
}

type Results struct {
	Document
	Score float32
}

type Store interface {
	Upsert(ctx context.Context, docs []Document) error
	Query(ctx context.Context, embedding []float32, topK int) ([]Results, error)
	Delete(ctx context.Context, ids []string) error
	DeleteBySource(ctx context.Context, source string) error
	Close() error
}
