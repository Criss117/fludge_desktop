package dbutils

import (
	"database/sql"
	"time"
)

func TimeToInt64(t time.Time) int64 {
	return t.UnixMilli()
}

func TimeFromInt64(t int64) time.Time {
	return time.UnixMilli(t)
}

func TimeToInt64Nullable(t *time.Time) *int64 {
	if t == nil {
		return nil
	}

	date := t.UnixMilli()

	return &date
}

func TimeFromInt64Nullable(t *int64) *time.Time {
	if t == nil {
		return nil
	}

	date := time.UnixMilli(*t)

	return &date
}

func TimeFromSQLNullable(ms sql.NullInt64) *time.Time {
	if !ms.Valid {
		return nil
	}

	date := time.UnixMilli(ms.Int64)

	return &date
}

func TimeToSQLNullable(t *time.Time) sql.NullInt64 {
	if t == nil {
		return sql.NullInt64{
			Valid: false,
		}
	}

	return sql.NullInt64{
		Int64: t.UnixMilli(),
		Valid: true,
	}
}

func IntToBool(num int64) bool {
	if num == 1 {
		return true
	}

	return false
}

func BoolToInt(value bool) int64 {
	if value {
		return 1
	}

	return 0
}

func StringToSQLNullable(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	str := sql.NullString{
		String: *s,
		Valid:  true,
	}

	return str
}

func StringFromSQLNullable(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}

	str := s.String

	return &str
}
