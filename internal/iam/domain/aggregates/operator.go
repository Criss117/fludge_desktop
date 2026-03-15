package aggregates

import (
	"desktop/internal/iam/domain/valueobjects"
	"desktop/internal/shared/lib"
	"time"
)

type PrimitiveOperator struct {
	ID        string
	Name      string
	Username  string
	Email     string
	Pin       string
	IsRoot    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Operator struct {
	id        string
	name      string
	username  string
	email     valueobjects.Email
	pin       valueobjects.Pin
	root      bool
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func NewOperator(name, username, email, pin string, isRoot bool) (*Operator, error) {
	validEmail, errEmail := valueobjects.NewEmail(email)

	if errEmail != nil {
		return nil, errEmail
	}

	validPin, errPin := valueobjects.NewPin(pin)

	if errPin != nil {
		return nil, errPin
	}

	return &Operator{
		id:        lib.GenerateUUID(),
		name:      name,
		username:  username,
		email:     validEmail,
		pin:       validPin,
		root:      isRoot,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		deletedAt: nil,
	}, nil
}

func ReconstituteOperator(
	id, name, username, email, pin string,
	isRoot bool,
	createdAt, updatedAt time.Time,
	deletedAt *time.Time,
) *Operator {
	return &Operator{
		id:        id,
		name:      name,
		username:  username,
		pin:       valueobjects.ReconstitutePin(pin),
		email:     valueobjects.ReconstituteEmail(email),
		root:      isRoot,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

func (op *Operator) ID() string {
	return op.id
}

func (op *Operator) Delete() {
	now := time.Now()
	op.deletedAt = &now
	op.updatedAt = now
}

func (op *Operator) IsRoot() bool {
	return op.root
}

func (op *Operator) IsActive() bool {
	if op.deletedAt != nil {
		return false
	}

	return true
}

func (op *Operator) ValidatePin(pin string) bool {
	return op.pin.ValidatePin(pin)
}

func (op *Operator) ToValues() PrimitiveOperator {
	return PrimitiveOperator{
		ID:        op.id,
		Name:      op.name,
		Username:  op.username,
		Email:     op.email.Value(),
		Pin:       op.pin.Value(),
		IsRoot:    op.root,
		CreatedAt: op.createdAt,
		UpdatedAt: op.updatedAt,
		DeletedAt: op.deletedAt,
	}
}

func (op *Operator) Equals(other *Operator) bool {
	return op.id == other.id
}
