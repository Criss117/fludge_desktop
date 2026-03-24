package aggregates

import (
	"desktop/internal/platform/iam/domain/valueobjects"
	"desktop/internal/shared/lib"
	"time"
)

type Operator struct {
	ID           string
	Name         string
	Username     string
	Email        valueobjects.Email
	Pin          valueobjects.Pin
	OperatorType valueobjects.OperatorType
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func NewOperator(
	name, username, email, pin string,
	OperatorType string,
) (*Operator, error) {
	validEmail, errEmail := valueobjects.NewEmail(email)

	if errEmail != nil {
		return nil, errEmail
	}

	validPin, errPin := valueobjects.NewPin(pin)

	if errPin != nil {
		return nil, errPin
	}

	validOperatorType, errOperatorType := valueobjects.NewOperatorType(OperatorType)

	if errOperatorType != nil {
		return nil, errOperatorType
	}

	return &Operator{
		ID:           lib.GenerateUUID(),
		Name:         name,
		Username:     username,
		Email:        validEmail,
		Pin:          validPin,
		OperatorType: validOperatorType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}, nil
}

func ReconstituteOperator(
	id, name, username, email, pin string,
	operatorType string,
	createdAt, updatedAt time.Time,
	deletedAt *time.Time,
) *Operator {
	return &Operator{
		ID:           id,
		Name:         name,
		Username:     username,
		Pin:          valueobjects.ReconstitutePin(pin),
		Email:        valueobjects.ReconstituteEmail(email),
		OperatorType: valueobjects.ReconstituteOperatorType(operatorType),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeletedAt:    deletedAt,
	}
}

func (op *Operator) Delete() {
	now := time.Now()
	op.DeletedAt = &now
	op.UpdatedAt = now
}

func (op *Operator) IsRoot() bool {
	return op.OperatorType.IsRoot()
}

func (op *Operator) IsActive() bool {
	if op.DeletedAt != nil {
		return false
	}

	return true
}

func (op *Operator) VerifyPIN(pin string) bool {
	return op.Pin.ValidatePin(pin)
}

func (op *Operator) Equals(other *Operator) bool {
	return op.ID == other.ID
}
