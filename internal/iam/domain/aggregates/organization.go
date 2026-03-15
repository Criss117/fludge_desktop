package aggregates

import (
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/domain/valueobjects"
	"desktop/internal/shared/lib"
	"time"
)

type PrimitiveOrganization struct {
	ID           string
	Name         string
	Slug         string
	LegalName    string
	Address      string
	Logo         *string
	ContactPhone *string
	ContactEmail *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Members      []*PrimitiveMember
	Teams        []*PrimitiveTeam
}

type Organization struct {
	id           string
	name         string
	slug         valueobjects.Slug
	legalName    string
	address      string
	logo         *string
	contactPhone *string
	contactEmail *valueobjects.Email
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    *time.Time
	members      []*Member
	teams        []*Team
}

func NewOrganization(
	name, slug, legalName, address string,
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

	validSlug := valueobjects.NewSlug(slug)

	organizationId := lib.GenerateUUID()

	defaultTeam := DefaultTeam(organizationId)

	return &Organization{
		id:           organizationId,
		name:         name,
		slug:         validSlug,
		legalName:    legalName,
		address:      address,
		logo:         logo,
		contactPhone: contactPhone,
		contactEmail: contactEmailVO,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
		deletedAt:    nil,
		teams:        []*Team{defaultTeam},
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
		validContactEmail, errEmail := valueobjects.NewEmail(*contactEmail)

		if errEmail != nil {
			return nil
		}

		contactEmailVO = &validContactEmail
	}

	return &Organization{
		id:           id,
		name:         name,
		slug:         valueobjects.ReconstituteSlug(slug),
		legalName:    legalName,
		address:      address,
		logo:         logo,
		contactPhone: contactPhone,
		contactEmail: contactEmailVO,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
		deletedAt:    deletedAt,
		members:      members,
		teams:        teams,
	}
}

func (o *Organization) ID() string {
	return o.id
}

func (o *Organization) Delete() {
	now := time.Now()
	o.deletedAt = &now
	o.updatedAt = now
}

func (o *Organization) IsActive() bool {
	if o.deletedAt != nil {
		return false
	}

	return true
}

func (o *Organization) FindMemberByOperatorId(operatorID string) *Member {
	for _, member := range o.members {
		if member.OperatorID() == operatorID {
			return member
		}
	}

	return nil
}

func (o *Organization) AddMember(member *Member) error {
	existingMember := o.FindMemberByOperatorId(member.OperatorID())

	if existingMember != nil {
		return derrors.ErrMemberAlreadyExists
	}

	o.members = append(o.members, member)
	o.updatedAt = time.Now()

	return nil
}

func (o *Organization) RemoveMember(member *Member) error {
	for i, m := range o.members {
		if m.Equals(member) {
			o.members = append(o.members[:i], o.members[i+1:]...)
			o.updatedAt = time.Now()
			return nil
		}
	}

	return derrors.ErrMemberNotFound
}

func (o *Organization) ToValues() PrimitiveOrganization {
	contactEmail := o.contactEmail.Value()

	primitiveMembers := make([]*PrimitiveMember, len(o.members))
	primitiveTeams := make([]*PrimitiveTeam, len(o.teams))

	for i, member := range o.members {
		memberValues := member.ToValues()

		primitiveMembers[i] = &memberValues
	}

	for i, team := range o.teams {
		teamValues := team.ToValues()

		primitiveTeams[i] = &teamValues
	}

	return PrimitiveOrganization{
		ID:           o.id,
		Name:         o.name,
		Slug:         o.slug.Value(),
		LegalName:    o.legalName,
		Address:      o.address,
		Logo:         o.logo,
		ContactPhone: o.contactPhone,
		ContactEmail: &contactEmail,
		CreatedAt:    o.createdAt,
		UpdatedAt:    o.updatedAt,
		DeletedAt:    o.deletedAt,
		Members:      primitiveMembers,
		Teams:        primitiveTeams,
	}
}
