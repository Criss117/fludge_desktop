package usecases

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/ports"
	"time"
)

type SignOutUseCase struct {
	appStateRepo ports.AppStateRepository
}

func NewSignOutUseCase(
	appStateRepo ports.AppStateRepository,
) *SignOutUseCase {
	return &SignOutUseCase{
		appStateRepo: appStateRepo,
	}
}

func (uc *SignOutUseCase) Execute(ctx context.Context) error {
	newAppState := aggregates.NewAppState(
		nil,
		nil,
		time.Now(),
	)

	err := uc.appStateRepo.Update(ctx, &newAppState)

	if err != nil {
		return err
	}

	return nil
}
