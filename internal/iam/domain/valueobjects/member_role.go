package valueobjects

import "desktop/internal/iam/domain/derrors"

type MemberRole struct {
	value string
}

var (
	MemberRoleRoot     = MemberRole{value: "ROOT"}
	MemberRoleEmployee = MemberRole{value: "EMPLOYEE"}
)

func NewMemberRole(role string) (MemberRole, error) {
	r := MemberRole{value: role}
	switch r {
	case MemberRoleRoot, MemberRoleEmployee:
		return r, nil
	default:
		return MemberRole{}, derrors.ErrMemberRoleInvalid
	}
}

func ReconstituteMemberRole(role string) MemberRole {
	return MemberRole{value: role}
}

func (mr MemberRole) Value() string {
	return mr.value
}

func (mr MemberRole) IsRoot() bool {
	return mr == MemberRoleRoot
}

func (mr MemberRole) IsMember() bool {
	return mr.value == MemberRoleEmployee.value
}

func (mr MemberRole) Equals(other MemberRole) bool {
	return mr == other
}
