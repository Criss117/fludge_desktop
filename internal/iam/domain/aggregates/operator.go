package aggregates

import (
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/domain/valueobjects"
	"desktop/internal/shared/lib"
	"time"
)

type OperatorOrganization struct {
	ID   string
	Slug string
	Name string
}

type Operator struct {
	ID         string
	Name       string
	Username   string
	Email      valueobjects.Email
	Pin        valueobjects.Pin
	Root       bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	IsMemberIn []*OperatorOrganization
}

func NewOperator(
	name, username, email, pin string,
	isRoot bool,
	isMemberIn []*OperatorOrganization,
) (*Operator, error) {

	if !isRoot && len(isMemberIn) != 1 {
		return nil, derrors.ErrOperatorCanBeMemberInOrganization
	}

	validEmail, errEmail := valueobjects.NewEmail(email)

	if errEmail != nil {
		return nil, errEmail
	}

	validPin, errPin := valueobjects.NewPin(pin)

	if errPin != nil {
		return nil, errPin
	}

	return &Operator{
		ID:         lib.GenerateUUID(),
		Name:       name,
		Username:   username,
		Email:      validEmail,
		Pin:        validPin,
		Root:       isRoot,
		IsMemberIn: isMemberIn,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  nil,
	}, nil
}

func ReconstituteOperator(
	id, name, username, email, pin string,
	isRoot bool,
	createdAt, updatedAt time.Time,
	deletedAt *time.Time,
	isMemberIn []*OperatorOrganization,
) *Operator {
	return &Operator{
		ID:         id,
		Name:       name,
		Username:   username,
		Pin:        valueobjects.ReconstitutePin(pin),
		Email:      valueobjects.ReconstituteEmail(email),
		Root:       isRoot,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeletedAt:  deletedAt,
		IsMemberIn: isMemberIn,
	}
}

func (op *Operator) Delete() {
	now := time.Now()
	op.DeletedAt = &now
	op.UpdatedAt = now
}

func (op *Operator) IsRoot() bool {
	return op.Root
}

func (op *Operator) IsActive() bool {
	if op.DeletedAt != nil {
		return false
	}

	return true
}

func (op *Operator) ValidatePin(pin string) bool {
	return op.Pin.ValidatePin(pin)
}

func (op *Operator) Equals(other *Operator) bool {
	return op.ID == other.ID
}
