package mappers

import (
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

func MapAppStateFromDomain(appState *aggregates.AppState) db.AppState {
	if appState == nil {
		return db.AppState{}
	}

	return db.AppState{
		ActiveOrganizationID: dbutils.StringToSQLNullable(appState.ActiveOrganizationID),
		ActiveOperatorID:     dbutils.StringToSQLNullable(appState.ActiveOperatorID),
		UpdatedAt:            dbutils.TimeToInt64(appState.UpdatedAt),
	}
}

func MapAppStateToDomain(dbAppState db.AppState) *aggregates.AppState {
	return aggregates.NewAppState(
		dbutils.StringFromSQLNullable(dbAppState.ActiveOrganizationID),
		dbutils.StringFromSQLNullable(dbAppState.ActiveOperatorID),
		dbutils.TimeFromInt64(dbAppState.UpdatedAt),
	)
}
