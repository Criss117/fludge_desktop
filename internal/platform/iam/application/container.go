package application

import (
	"desktop/internal/platform/iam/application/queries"
	"desktop/internal/platform/iam/application/usecases"
	"desktop/internal/platform/iam/domain/ports"
	"desktop/internal/shared/db/dbutils"
)

type UseCasesContainer struct {
	SignIn               *usecases.SignIn
	RegisterRootOperator *usecases.RegisterRootOperator
	RegisterOrganization *usecases.RegisterOrganization
	UpdateOrganization   *usecases.UpdateOrganization
	UpdateTeam           *usecases.UpdateTeam
	CreateTeam           *usecases.CreateTeam
	DeleteTeam           *usecases.DeleteTeam
}

type QueriesContainer struct {
	FindOneOrganization                 *queries.FindOneOrganization
	FindManyOrganizationsByRootOperator *queries.FindManyOrganizationsByRootOperator
}

func NewUseCasesContainer(
	txManager dbutils.TxManager,
	operatorRepository ports.OperatorRepository,
	organizationRepository ports.OrganizationRepository,
	memberRepository ports.OrganizationMemberRepository,
	teamRepository ports.OrganizationTeamRepository,
) *UseCasesContainer {
	// Operator - UseCases
	signIn := usecases.NewSignIn(operatorRepository)
	registerRootOperator := usecases.NewRegisterRootOperator(operatorRepository)

	// Organization - UseCases
	registerOrganization := usecases.NewRegisterOrganization(
		txManager,
		organizationRepository,
		teamRepository,
		memberRepository,
	)
	updateOrganization := usecases.NewUpdateOrganization(
		organizationRepository,
	)

	// Team - UseCases
	updateTeam := usecases.NewUpdateTeam(
		teamRepository,
	)
	createTeam := usecases.NewCreateTeam(
		teamRepository,
	)
	deleteTeam := usecases.NewDeleteTeam(
		teamRepository,
	)

	return &UseCasesContainer{
		SignIn:               signIn,
		RegisterRootOperator: registerRootOperator,
		RegisterOrganization: registerOrganization,
		UpdateOrganization:   updateOrganization,
		UpdateTeam:           updateTeam,
		CreateTeam:           createTeam,
		DeleteTeam:           deleteTeam,
	}
}

func NewQueriesContainer(organizationRepository ports.OrganizationRepository) *QueriesContainer {
	// Organization - Queries
	findOneOrganization := queries.NewFindOneOrganization(organizationRepository)
	findManyOrganizationsByRootOperator := queries.NewFindManyOrganizationsByRootOperator(organizationRepository)

	return &QueriesContainer{
		FindOneOrganization:                 findOneOrganization,
		FindManyOrganizationsByRootOperator: findManyOrganizationsByRootOperator,
	}
}
