package application

import (
	"desktop/internal/auth/application/dtos"
	"desktop/internal/auth/domain"
)

type FindAllOperatorsUseCase struct {
	operatorRepository domain.OperatorRepository
}

func NewFindAllOperatorsUseCase(operatorRepository domain.OperatorRepository) *FindAllOperatorsUseCase {
	return &FindAllOperatorsUseCase{
		operatorRepository: operatorRepository,
	}
}

func (uc *FindAllOperatorsUseCase) Execute() ([]*dtos.SummaryOperatorDTO, error) {
	allDbOperators, err := uc.operatorRepository.FinAll()

	if err != nil {
		return nil, err
	}

	allOperators := make([]*dtos.SummaryOperatorDTO, len(allDbOperators))

	for i, operator := range allDbOperators {
		allOperators[i] = dtos.DomainToSummaryOperatorDTO(operator)
	}

	return allOperators, nil
}
