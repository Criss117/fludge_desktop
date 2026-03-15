package valueobjects

import "desktop/internal/iam/domain/derrors"

type MemberRole struct {
	value string
}

var (
	MemberRoleRoot   = MemberRole{value: "ROOT"}
	MemberRoleMember = MemberRole{value: "MEMBER"}
)

func NewMemberRole(role string) (MemberRole, error) {
	r := MemberRole{value: role}
	switch r {
	case MemberRoleRoot, MemberRoleMember:
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
	return mr == MemberRoleMember
}

func (mr MemberRole) Equals(other MemberRole) bool {
	return mr == other
}
