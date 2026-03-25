package iam

import (
	"context"
	"desktop/internal/appstate"
	"desktop/internal/platform/iam/application"
	"desktop/internal/platform/iam/application/commands"
	"desktop/internal/platform/iam/application/responses"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/derrors"
)

type GetCtxFunc func() context.Context
type GetSessionFunc func() *appstate.SessionState
type OnStateChange func(e appstate.StateChangeEvent)

type IamHandler struct {
	app           *application.Container
	getCtx        GetCtxFunc
	getSession    GetSessionFunc
	onStateChange OnStateChange
}

func NewIamHandler(
	app *application.Container,
	onStateChange OnStateChange,
	getCtx GetCtxFunc,
	getSession GetSessionFunc,
) *IamHandler {
	return &IamHandler{
		getCtx:        getCtx,
		app:           app,
		onStateChange: onStateChange,
		getSession:    getSession,
	}
}

func (h *IamHandler) getCurrentContext() context.Context {
	ctx := context.Background()
	if h.getCtx != nil {
		ctx = h.getCtx()
	}
	return ctx
}

// Operator

func (h *IamHandler) currentOperator() (*aggregates.Operator, error) {
	sessionState := h.getSession()

	if sessionState == nil {
		return nil, derrors.ErrNoActiveOperator
	}

	if sessionState.ActiveOperator == nil {
		return nil, derrors.ErrNoActiveOperator
	}

	return sessionState.ActiveOperator.Operator, nil
}

func (h *IamHandler) RegisterRootOperator(cmd *commands.RegisterRootOperator) (*responses.Operator, error) {
	ctx := h.getCurrentContext()

	newOperator, err := h.app.RegisterRootOperator.Execute(ctx, cmd)

	if err != nil {
		return nil, err
	}

	h.onStateChange(appstate.StateChangeEvent{
		Type:     appstate.SignUp,
		Operator: newOperator,
	})

	return responses.OperatorResponseFromDomain(newOperator), nil
}

func (h *IamHandler) SignIn(cmd *commands.SignIn) (*responses.Operator, error) {
	ctx := h.getCurrentContext()

	newOperator, err := h.app.SignIn.Execute(ctx, cmd)

	if err != nil {
		return nil, err
	}

	h.onStateChange(appstate.StateChangeEvent{
		Type:     appstate.SignIn,
		Operator: newOperator,
	})

	return responses.OperatorResponseFromDomain(newOperator), nil
}

func (h *IamHandler) SignOut() {
	h.onStateChange(appstate.StateChangeEvent{
		Type: appstate.SignOut,
	})
}

// Organization

func (h *IamHandler) SwitchOrganization(cmd *commands.SwitchOrganization) (*responses.Organization, error) {
	ctx := h.getCurrentContext()
	currentOperator, err := h.currentOperator()

	if err != nil {
		return nil, err
	}

	organization, err := h.app.FindOneOrganization.Execute(
		ctx,
		currentOperator.ID,
		cmd.OrganizationId,
	)

	if err != nil {
		return nil, err
	}

	h.onStateChange(appstate.StateChangeEvent{
		Type:         appstate.SwitchOrganization,
		Organization: organization,
	})

	return responses.OrganizationResponseFromDomain(organization), nil
}

func (h *IamHandler) FindOneOrganization(cmd *commands.FindOneOrganization) (*responses.Organization, error) {
	ctx := h.getCurrentContext()
	currentOperator, err := h.currentOperator()

	if err != nil {
		return nil, err
	}

	organization, err := h.app.FindOneOrganization.Execute(
		ctx,
		currentOperator.ID,
		cmd.OrganizationId,
	)

	if err != nil {
		return nil, err
	}

	h.onStateChange(appstate.StateChangeEvent{
		Type:         appstate.SwitchOrganization,
		Organization: organization,
	})

	return responses.OrganizationResponseFromDomain(organization), nil
}

func (h *IamHandler) CreateOrganization(cmd *commands.RegisterOrganization) (*responses.Organization, error) {
	ctx := h.getCurrentContext()
	currentOperator, err := h.currentOperator()

	if err != nil {
		return nil, err
	}

	h.app.RegisterOrganization.Execute(ctx, currentOperator.ID, cmd)

	return nil, nil
}

func (h *IamHandler) UpdateOrganization(cmd *commands.UpdateOrganization) (*responses.Organization, error) {
	ctx := h.getCurrentContext()
	currentOperator, err := h.currentOperator()

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

	return responses.OrganizationResponseFromDomain(updatedOrg), nil
}

func (h *IamHandler) FindManyOrganizationsByRootOperator() ([]*responses.Organization, error) {
	ctx := h.getCurrentContext()
	currentOperator, err := h.currentOperator()

	if err != nil {
		return nil, err
	}

	if !currentOperator.IsRoot() {
		return nil, derrors.ErrOperatorMustBeRoot
	}

	organizations, err := h.app.FindManyOrganizationsByRootOperator.Execute(ctx, currentOperator.ID)

	if err != nil {
		return nil, err
	}

	orgRes := make([]*responses.Organization, len(organizations))

	for i, org := range organizations {
		orgRes[i] = responses.OrganizationResponseFromDomain(org)
	}

	return orgRes, nil
}
