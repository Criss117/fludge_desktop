package aggregates

import (
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/platform/iam/domain/valueobjects"
	"desktop/internal/shared/lib"
	"time"
)

type Organization struct {
	ID           string
	Name         string
	Slug         valueobjects.Slug
	LegalName    string
	Address      string
	Logo         *string
	ContactPhone *string
	ContactEmail *valueobjects.Email
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Members      []*Member
	Teams        []*Team
}

func NewOrganization(
	name, legalName, address string,
	logo, contactPhone, contactEmail *string,
) (*Organization, error) {
	var contactEmailVO *valueobjects.Email = nil

	if contactEmail != nil {
		validContactEmail, errEmail := valueobjects.NewEmail(*contactEmail)

		if errEmail != nil {
			return nil, errEmail
		}

		contactEmailVO = &validContactEmail
	}

	validSlug := valueobjects.NewSlug(name)

	organizationId := lib.GenerateUUID()

	defaultTeam := DefaultTeam(organizationId)

	return &Organization{
		ID:           organizationId,
		Name:         name,
		Slug:         validSlug,
		LegalName:    legalName,
		Address:      address,
		Logo:         logo,
		ContactPhone: contactPhone,
		ContactEmail: contactEmailVO,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
		Teams:        []*Team{defaultTeam},
		Members:      []*Member{},
	}, nil
}

func ReconstituteOrganization(
	id, name, slug, legalName, address string,
	logo, contactPhone, contactEmail *string,
	createdAt, updatedAt time.Time,
	deletedAt *time.Time,
	members []*Member,
	teams []*Team,
) *Organization {
	var contactEmailVO *valueobjects.Email = nil

	if contactEmail != nil {
		cEmail := valueobjects.ReconstituteEmail(*contactEmail)

		contactEmailVO = &cEmail
	}

	return &Organization{
		ID:           id,
		Name:         name,
		Slug:         valueobjects.ReconstituteSlug(slug),
		LegalName:    legalName,
		Address:      address,
		Logo:         logo,
		ContactPhone: contactPhone,
		ContactEmail: contactEmailVO,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeletedAt:    deletedAt,
		Members:      members,
		Teams:        teams,
	}
}

func (o *Organization) Delete() {
	now := time.Now()
	o.DeletedAt = &now
	o.UpdatedAt = now
}

func (o *Organization) IsActive() bool {
	if o.DeletedAt != nil {
		return false
	}

	return true
}

func (o *Organization) FindMemberByOperatorId(operatorID string) *Member {
	for _, member := range o.Members {
		if member.OperatorID == operatorID {
			return member
		}
	}

	return nil
}

func (o *Organization) FindTeamsByOperatorId(operatorID string) []*Team {
	teams := make([]*Team, 0, len(o.Teams))

	for _, team := range o.Teams {
		if team.OperatorIsMember(operatorID) {
			teams = append(teams, team)
		}
	}

	return teams
}

func (o *Organization) AddMember(member *Member) error {
	existingMember := o.FindMemberByOperatorId(member.OperatorID)

	if existingMember != nil {
		return derrors.ErrMemberAlreadyExists
	}

	if member.Role.IsRoot() {
		return derrors.ErrRootMemberCannotBeAdded
	}

	o.Members = append(o.Members, member)
	o.UpdatedAt = time.Now()

	return nil
}

func (o *Organization) RemoveMember(member *Member) error {
	for i, m := range o.Members {
		if m.Equals(member) {
			o.Members = append(o.Members[:i], o.Members[i+1:]...)
			o.UpdatedAt = time.Now()
			return nil
		}
	}

	return derrors.ErrMemberNotFound
}

func (o *Organization) UpdateDetails(
	name, legalName, address string,
	logo, contactPhone, contactEmail *string,
) error {
	var contactEmailVO *valueobjects.Email = nil

	if contactEmail != nil {
		validContactEmail, errEmail := valueobjects.NewEmail(*contactEmail)

		if errEmail != nil {
			return errEmail
		}

		contactEmailVO = &validContactEmail
	}

	validSlug := valueobjects.NewSlug(name)

	o.Name = name
	o.Slug = validSlug
	o.LegalName = legalName
	o.Address = address
	o.Logo = logo
	o.ContactPhone = contactPhone
	o.ContactEmail = contactEmailVO
	o.UpdatedAt = time.Now()

	return nil
}
