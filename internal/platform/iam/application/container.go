package application

import (
	"desktop/internal/platform/iam/application/usecases"
	"desktop/internal/platform/iam/domain/ports"
	"desktop/internal/shared/db/dbutils"
)

type Container struct {
	SignIn                              *usecases.SignIn
	RegisterRootOperator                *usecases.RegisterRootOperator
	FindOneOrganization                 *usecases.FindOneOrganization
	RegisterOrganization                *usecases.RegisterOrganization
	FindManyOrganizationsByRootOperator *usecases.FindManyOrganizationsByRootOperator
	UpdateOrganization                  *usecases.UpdateOrganization
}

func NewContainer(
	txManager dbutils.TxManager,
	operatorRepository ports.OperatorRepository,
	organizationRepository ports.OrganizationRepository,
	memberRepository ports.OrganizationMemberRepository,
	teamRepository ports.OrganizationTeamRepository,
) *Container {
	// Operator - UseCases
	signIn := usecases.NewSignIn(operatorRepository)
	registerRootOperator := usecases.NewRegisterRootOperator(operatorRepository)

	// Organization - UseCases
	findOneOrganization := usecases.NewFindOneOrganization(organizationRepository)
	registerOrganization := usecases.NewRegisterOrganization(
		txManager,
		organizationRepository,
		teamRepository,
		memberRepository,
	)
	findManyOrganizationsByRootOperator := usecases.NewFindManyOrganizationsByRootOperator(
		organizationRepository,
	)
	updateOrganization := usecases.NewUpdateOrganization(
		organizationRepository,
	)

	return &Container{
		SignIn:                              signIn,
		RegisterRootOperator:                registerRootOperator,
		FindOneOrganization:                 findOneOrganization,
		RegisterOrganization:                registerOrganization,
		FindManyOrganizationsByRootOperator: findManyOrganizationsByRootOperator,
		UpdateOrganization:                  updateOrganization,
	}
}
