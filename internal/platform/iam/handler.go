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
	getCtx GetCtxFunc,
	onStateChange OnStateChange,
	app *application.Container,
) *IamHandler {
	return &IamHandler{
		getCtx:        getCtx,
		app:           app,
		onStateChange: onStateChange,
	}
}

func (h *IamHandler) getCurrentContext() context.Context {
	ctx := context.Background()
	if h.getCtx != nil {
		ctx = h.getCtx()
	}
	return ctx
}

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

func (h *IamHandler) RegisterRootOperator(cmd *commands.RegisterRootOperator) (*responses.OperatorResponse, error) {
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

func (h *IamHandler) SignIn(cmd *commands.SignIn) (*responses.OperatorResponse, error) {
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

func (h *IamHandler) SwitchOrganization(cmd *commands.SwitchOrganization) (*responses.OrganizationResponse, error) {
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

func (h *IamHandler) FindOneOrganization(cmd *commands.FindOneOrganization) (*responses.OrganizationResponse, error) {
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

func (h *IamHandler) CreateOrganization(cmd *commands.RegisterOrganization) (*responses.OrganizationResponse, error) {
	ctx := h.getCurrentContext()
	currentOperator, err := h.currentOperator()

	if err != nil {
		return nil, err
	}

	h.app.RegisterOrganization.Execute(ctx, currentOperator.ID, cmd)

	return nil, nil
}
