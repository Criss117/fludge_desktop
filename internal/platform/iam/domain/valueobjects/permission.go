// valueobjects/permission.go
package valueobjects

import (
	"fmt"
	"strings"
)

type Action string
type Resource string

const (
	ActionCreate Action = "create"
	ActionRead   Action = "read"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"
)

const (
	ResourceTeam     Resource = "team"
	ResourceProduct  Resource = "product"
	ResourceTicket   Resource = "ticket"
	ResourceCustomer Resource = "customer"
	ResourceCategory Resource = "category"
	ResourceSupplier Resource = "supplier"
	ResourceEmployee Resource = "employee"
)

var validActions = map[Action]struct{}{
	ActionCreate: {},
	ActionRead:   {},
	ActionUpdate: {},
	ActionDelete: {},
}

var validResources = map[Resource]struct{}{
	ResourceTeam:     {},
	ResourceProduct:  {},
	ResourceTicket:   {},
	ResourceCustomer: {},
	ResourceCategory: {},
	ResourceSupplier: {},
	ResourceEmployee: {},
}

type Permission struct {
	action   Action
	resource Resource
}

func NewPermission(raw string) (Permission, error) {
	parts := strings.SplitN(raw, ":", 2)
	if len(parts) != 2 {
		return Permission{}, fmt.Errorf("permission format invalid: expected action:resource, got %q", raw)
	}

	action := Action(parts[0])
	resource := Resource(parts[1])

	if _, ok := validActions[action]; !ok {
		return Permission{}, fmt.Errorf("unknown action %q", action)
	}
	if _, ok := validResources[resource]; !ok {
		return Permission{}, fmt.Errorf("unknown resource %q", resource)
	}

	return Permission{action: action, resource: resource}, nil
}

func PermissionFromStorage(raw string) Permission {
	parts := strings.SplitN(raw, ":", 2)
	return Permission{
		action:   Action(parts[0]),
		resource: Resource(parts[1]),
	}
}

func (p Permission) Value() string {
	return fmt.Sprintf("%s:%s", p.action, p.resource)
}

func (p Permission) Action() Action {
	return p.action
}

func (p Permission) Resource() Resource {
	return p.resource
}

func (p Permission) Equals(other Permission) bool {
	return p.action == other.action && p.resource == other.resource
}

func AllPermissions() []Permission {
	actions := []Action{ActionCreate, ActionRead, ActionUpdate, ActionDelete}
	resources := []Resource{
		ResourceTeam, ResourceProduct, ResourceTicket,
		ResourceCustomer, ResourceCategory, ResourceSupplier, ResourceEmployee,
	}

	perms := make([]Permission, 0, len(actions)*len(resources))

	for _, a := range actions {
		for _, r := range resources {
			perms = append(perms, Permission{action: a, resource: r})
		}
	}
	return perms
}
