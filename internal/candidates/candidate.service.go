package candidates

import (
	"context"

	"github.com/LeMinh0706/fakeapi/internal/db"
)

type CandidateService struct {
	q *db.Queries
}

func (s *CandidateService) GetCandidates(ctx context.Context) ([]db.GetCandidateRow, error) {
	return s.q.GetCandidate(ctx)
}

func NewCandidateService(q *db.Queries) *CandidateService {
	return &CandidateService{
		q: q,
	}
}
