package platform

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

func FromMillisNullable(ms sql.NullInt64) *time.Time {
	if !ms.Valid {
		return nil
	}

	date := time.UnixMilli(ms.Int64)

	return &date
}

func ToMillisNullable(t *time.Time) sql.NullInt64 {
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

func FromMillis(ms int64) time.Time {
	return time.UnixMilli(ms)
}

func ToMillis(t time.Time) int64 {
	return t.UnixMilli()
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

func ToStringNullable(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{
			Valid: false,
		}
	}

	str := sql.NullString{
		String: *s,
		Valid:  true,
	}

	return str
}

func FromStringNullable(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}

	str := s.String

	return &str
}
