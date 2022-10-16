package utils

import "database/sql"

func NewNullInt(number int16) sql.NullInt16 {
	invalidConditions := number < 0
	if invalidConditions {
		return sql.NullInt16{}
	}
	return sql.NullInt16{
		Int16: number,
		Valid: true,
	}
}

func NewNullString(s string) sql.NullString {
	invalidConditions := len(s) == 0 || s == "" || len(s) < 2

	if invalidConditions {
		return sql.NullString{}
	}

	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
