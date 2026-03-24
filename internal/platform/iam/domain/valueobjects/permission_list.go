// valueobjects/permission_list.go
package valueobjects

import (
	"desktop/internal/platform/iam/domain/derrors"
	"encoding/json"
)

type PermissionList []Permission

// NewPermissionList — desde input externo, valida cada item
func NewPermissionList(raw []string) (PermissionList, error) {
	if len(raw) == 0 {
		return nil, derrors.ErrPermissionListEmpty
	}
	perms := make(PermissionList, 0, len(raw))
	for _, r := range raw {
		p, err := NewPermission(r)
		if err != nil {
			return nil, err
		}
		perms = append(perms, p)
	}
	return perms, nil
}

// PermissionListFromStorage — desde DB con []string, confía en los datos
func PermissionListFromStorage(raw []string) PermissionList {
	perms := make(PermissionList, 0, len(raw))
	for _, r := range raw {
		perms = append(perms, PermissionFromStorage(r))
	}
	return perms
}

// PermissionListFromJSON — desde DB con json.RawMessage, confía en los datos
func PermissionListFromJSON(raw json.RawMessage) (PermissionList, error) {
	var strings []string
	if err := json.Unmarshal(raw, &strings); err != nil {
		return nil, err
	}
	return PermissionListFromStorage(strings), nil
}

// NewPermissionListFromJSON — desde input externo con json.RawMessage, valida cada item
func NewPermissionListFromJSON(raw json.RawMessage) (PermissionList, error) {
	var strings []string
	if err := json.Unmarshal(raw, &strings); err != nil {
		return nil, err
	}
	return NewPermissionList(strings)
}

// ToStrings — para persistencia y responses
func (pl PermissionList) ToStrings() []string {
	raw := make([]string, len(pl))
	for i, p := range pl {
		raw[i] = p.Value()
	}
	return raw
}

// ToJSON — para persistencia cuando la DB espera json.RawMessage
func (pl PermissionList) ToJSON() (json.RawMessage, error) {
	return json.Marshal(pl.ToStrings())
}

// Has — verifica si un permiso específico está en la lista
func (pl PermissionList) Has(p Permission) bool {
	return pl.includes(p)
}

// HasAll — equivalente a hasAllPermissions en TS
func (pl PermissionList) HasAll(required PermissionList) bool {
	for _, r := range required {
		if !pl.includes(r) {
			return false
		}
	}
	return true
}

// HasSome — equivalente a hasSomePermissions en TS
func (pl PermissionList) HasSome(required PermissionList) bool {
	for _, r := range required {
		if pl.includes(r) {
			return true
		}
	}
	return false
}

func (pl PermissionList) includes(p Permission) bool {
	for _, item := range pl {
		if item.Equals(p) {
			return true
		}
	}
	return false
}
