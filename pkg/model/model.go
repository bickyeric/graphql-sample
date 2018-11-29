package model

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

// NewNullString ...
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// NewNullInt64 ...
func NewNullInt64(i int) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: int64(i),
		Valid: true,
	}
}

// NewNullTime ...
func NewNullTime(t time.Time) mysql.NullTime {
	zeroTime := time.Time{}
	if t == zeroTime {
		return mysql.NullTime{}
	}
	return mysql.NullTime{
		Time:  t,
		Valid: true,
	}
}
