package benchmark

import (
	"context"

	"github.com/LeMinh0706/fakeapi/internal/db"
)

type BenchmarkService struct {
	q *db.Queries
}

func (s *BenchmarkService) GetResults(ctx context.Context) ([]string, error) {
	return s.q.GetResults(ctx)
}

func NewBenchmarkService(q *db.Queries) *BenchmarkService {
	return &BenchmarkService{
		q: q,
	}
}
