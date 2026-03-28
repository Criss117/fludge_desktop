package infrastructure

import (
	"context"
	"desktop/internal/appstate"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/derrors"
)

type GetCtxFunc func() context.Context
type GetSessionFunc func() *appstate.SessionState

type BaseHandler struct {
	getCtx     GetCtxFunc
	getSession GetSessionFunc
}

func NewBaseHandler(getCtx GetCtxFunc, getSession GetSessionFunc) BaseHandler {
	return BaseHandler{
		getCtx:     getCtx,
		getSession: getSession,
	}
}

func (b *BaseHandler) Context() context.Context {
	ctx := context.Background()
	if b.getCtx != nil {
		ctx = b.getCtx()
	}
	return ctx
}

func (b *BaseHandler) CurrentOrganization() (*aggregates.Organization, error) {
	sessionState := b.getSession()

	if sessionState == nil || sessionState.ActiveOrganization == nil {
		return nil, derrors.ErrNoActiveOrganization
	}

	return sessionState.ActiveOrganization, nil
}

func (b *BaseHandler) CurrentOperator() (*aggregates.Operator, error) {
	sessionState := b.getSession()

	if sessionState == nil || sessionState.ActiveOperator == nil {
		return nil, derrors.ErrNoActiveOperator
	}

	return sessionState.ActiveOperator.Operator, nil
}
