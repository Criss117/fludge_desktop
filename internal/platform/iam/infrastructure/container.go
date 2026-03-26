package infrastructure

import (
	"desktop/internal/platform/iam/domain/ports"
	"desktop/internal/platform/iam/infrastructure/repositories"
	"desktop/internal/shared/db"
)

type Container struct {
	AppStateRepository           ports.AppStateRepository
	OperatorRepository           ports.OperatorRepository
	OrganizationRepository       ports.OrganizationRepository
	OrganizationMemberRepository ports.OrganizationMemberRepository
	OrganizationTeamRepository   ports.OrganizationTeamRepository
}

func NewContainer(querires *db.Queries) *Container {
	operatorRepository := repositories.NewSqliteOperatorRepository(querires)
	organizationRepository := repositories.NewSqliteOrganizationRepository(querires)
	memberRepository := repositories.NewSqliteOrganizationMemberRepository(querires)
	teamRepository := repositories.NewSqliteOrganizationTeamRepository(querires)
	appStateRepository := repositories.NewSqliteAppRepository(
		querires,
		organizationRepository,
		operatorRepository,
	)

	return &Container{
		AppStateRepository:           appStateRepository,
		OperatorRepository:           operatorRepository,
		OrganizationRepository:       organizationRepository,
		OrganizationMemberRepository: memberRepository,
		OrganizationTeamRepository:   teamRepository,
	}
}
