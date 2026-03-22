package ports

import "context"

type CategoryChecker interface {
	Exists(ctx context.Context, organizationID, categoryID string) (bool, error)
}
