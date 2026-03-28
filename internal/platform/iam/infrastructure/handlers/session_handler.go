package handlers

import (
	"desktop/internal/appstate"
	"desktop/internal/platform/iam/application"
	"desktop/internal/platform/iam/application/commands"
	"desktop/internal/platform/iam/application/responses"
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/shared/infrastructure"
)

type IamSessionHandler struct {
	baseHandler   infrastructure.BaseHandler
	app           *application.UseCasesContainer
	queries       *application.QueriesContainer
	onStateChange OnStateChange
}

func NewIamSessionHandler(
	app *application.UseCasesContainer,
	queries *application.QueriesContainer,
	onStateChange OnStateChange,
	getCtx infrastructure.GetCtxFunc,
	getSession infrastructure.GetSessionFunc,
) IamSessionHandler {
	return IamSessionHandler{
		app:           app,
		onStateChange: onStateChange,
		queries:       queries,
		baseHandler:   infrastructure.NewBaseHandler(getCtx, getSession),
	}
}

func (h *IamSessionHandler) RegisterRootOperator(cmd *commands.RegisterRootOperator) (*responses.Operator, error) {
	ctx := h.baseHandler.Context()

	newOperator, err := h.app.RegisterRootOperator.Execute(ctx, cmd)

	if err != nil {
		return nil, err
	}

	h.onStateChange(appstate.StateChangeEvent{
		Type:     appstate.SignUp,
		Operator: newOperator,
	})

	res := responses.OperatorFromDomain(newOperator)
	return &res, nil
}

func (h *IamSessionHandler) SignIn(cmd *commands.SignIn) (*responses.Operator, error) {
	ctx := h.baseHandler.Context()

	newOperator, err := h.app.SignIn.Execute(ctx, cmd)

	if err != nil {
		return nil, err
	}

	h.onStateChange(appstate.StateChangeEvent{
		Type:     appstate.SignIn,
		Operator: newOperator,
	})

	res := responses.OperatorFromDomain(newOperator)

	return &res, nil
}

func (h *IamSessionHandler) SignOut() {
	h.onStateChange(appstate.StateChangeEvent{
		Type: appstate.SignOut,
	})
}

func (h *IamSessionHandler) SwitchOrganization(cmd *commands.SwitchOrganization) (*responses.Organization, error) {
	ctx := h.baseHandler.Context()
	currentOperator, err := h.baseHandler.CurrentOperator()

	if err != nil {
		return nil, err
	}

	if !currentOperator.IsRoot() {
		return nil, derrors.ErrOperatorMustBeRoot
	}

	organization, err := h.queries.FindOneOrganization.ExecuteAggregate(ctx, currentOperator.ID, cmd.OrganizationId)

	if err != nil {
		return nil, err
	}

	h.onStateChange(appstate.StateChangeEvent{
		Type:         appstate.SwitchOrganization,
		Organization: organization,
	})

	res := responses.OrganizationFromDomain(organization)

	return &res, nil
}
