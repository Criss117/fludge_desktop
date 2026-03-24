package valueobjects

import "desktop/internal/platform/iam/domain/derrors"

type OperatorType struct {
	value string
}

var (
	RootOperator     = OperatorType{value: "ROOT"}
	EmployeeOperator = OperatorType{value: "EMPLOYEE"}
)

func NewOperatorType(operatorType string) (OperatorType, error) {
	ot := OperatorType{value: operatorType}
	switch ot {
	case RootOperator, EmployeeOperator:
		return ot, nil
	default:
		return OperatorType{}, derrors.ErrOperatorTypeInvalid
	}
}

func ReconstituteOperatorType(operatorType string) OperatorType {
	return OperatorType{value: operatorType}
}

func (ot OperatorType) ToValue() string {
	return ot.value
}

func (ot OperatorType) IsRoot() bool {
	return ot == RootOperator
}

func (ot OperatorType) IsMember() bool {
	return ot.value == EmployeeOperator.value
}
