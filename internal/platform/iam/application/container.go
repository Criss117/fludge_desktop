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

	return &UseCasesContainer{
		SignIn:               signIn,
		RegisterRootOperator: registerRootOperator,
		RegisterOrganization: registerOrganization,
		UpdateOrganization:   updateOrganization,
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
