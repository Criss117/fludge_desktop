package ports

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
)

type MemberRepository interface {
	Create(ctx context.Context, member *aggregates.Member) error
	// Update(member *Member) error
}
