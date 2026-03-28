package handlers

import (
	"desktop/internal/appstate"
	"desktop/internal/platform/iam/application"
	"desktop/internal/platform/iam/application/commands"
	"desktop/internal/platform/iam/application/responses"
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/shared/infrastructure"
)

type IamOrganizationHandler struct {
	baseHandler   infrastructure.BaseHandler
	app           *application.UseCasesContainer
	queries       *application.QueriesContainer
	onStateChange OnStateChange
}

func NewIamOrganizationHandler(
	app *application.UseCasesContainer,
	queries *application.QueriesContainer,
	onStateChange OnStateChange,
	getCtx infrastructure.GetCtxFunc,
	getSession infrastructure.GetSessionFunc,
) IamOrganizationHandler {
	return IamOrganizationHandler{
		app:           app,
		queries:       queries,
		onStateChange: onStateChange,
		baseHandler:   infrastructure.NewBaseHandler(getCtx, getSession),
	}
}

func (h *IamOrganizationHandler) FindOneOrganization(cmd *commands.FindOneOrganization) (*responses.Organization, error) {
	ctx := h.baseHandler.Context()
	currentOperator, err := h.baseHandler.CurrentOperator()

	if err != nil {
		return nil, err
	}

	organization, err := h.queries.FindOneOrganization.Execute(
		ctx,
		currentOperator.ID,
		cmd.OrganizationId,
	)

	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (h *IamOrganizationHandler) RegisterOrganization(cmd *commands.RegisterOrganization) (*responses.Organization, error) {
	ctx := h.baseHandler.Context()
	currentOperator, err := h.baseHandler.CurrentOperator()

	if err != nil {
		return nil, err
	}

	organization, err := h.app.RegisterOrganization.Execute(ctx, currentOperator.ID, cmd)

	if err != nil {
		return nil, err
	}

	res := responses.OrganizationFromDomain(organization)

	return &res, nil
}

func (h *IamOrganizationHandler) UpdateOrganization(cmd *commands.UpdateOrganization) (*responses.Organization, error) {
	ctx := h.baseHandler.Context()
	currentOperator, err := h.baseHandler.CurrentOperator()

	if err != nil {
		return nil, err
	}

	updatedOrg, errUpdated := h.app.UpdateOrganization.Execute(ctx, currentOperator.ID, cmd)

	if errUpdated != nil {
		return nil, errUpdated
	}

	h.onStateChange(appstate.StateChangeEvent{
		Type:         appstate.SwitchOrganization,
		Organization: updatedOrg,
	})

	res := responses.OrganizationFromDomain(updatedOrg)

	return &res, nil
}

func (h *IamOrganizationHandler) FindManyOrganizationsByRootOperator() ([]responses.Organization, error) {
	ctx := h.baseHandler.Context()
	currentOperator, err := h.baseHandler.CurrentOperator()

	if err != nil {
		return nil, err
	}

	if !currentOperator.IsRoot() {
		return nil, derrors.ErrOperatorMustBeRoot
	}

	organizations, err := h.queries.FindManyOrganizationsByRootOperator.Execute(ctx, currentOperator.ID)

	if err != nil {
		return nil, err
	}

	return organizations, nil
}
