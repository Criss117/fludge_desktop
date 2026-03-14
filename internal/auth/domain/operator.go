package domain

import (
	"desktop/internal/shared/db/platform"
	"time"
)

type Operator struct {
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

func NewOperator(name, username, email, pin string, isRoot bool) Operator {
	return Operator{
		ID:        platform.GenerateUUID(),
		Name:      name,
		Username:  username,
		Email:     email,
		Pin:       pin,
		IsRoot:    isRoot,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}

func ReconstituteOperator(id, name, username, email, pin string, isRoot bool, createdAt, updatedAt time.Time, deletedAt *time.Time) Operator {
	return Operator{
		ID:        id,
		Name:      name,
		Username:  username,
		Pin:       pin,
		Email:     email,
		IsRoot:    isRoot,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

func (op *Operator) Delete() {
	now := time.Now()
	op.DeletedAt = &now
	op.UpdatedAt = now
}
