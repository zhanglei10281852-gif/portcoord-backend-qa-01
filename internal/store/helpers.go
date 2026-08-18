package store

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// nowStamp returns the current time in RFC3339 format for DB columns.
func nowStamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")
}

// parseTime parses an RFC3339 time string, returning zero on failure.
func parseTime(s string) time.Time {
	if s == "" {
		return time.Time{}
	}
	t, err := time.Parse("2006-01-02T15:04:05Z07:00", s)
	if err != nil {
		t, err = time.Parse(time.RFC3339, s)
		if err != nil {
			return time.Time{}
		}
	}
	return t
}

// nullString converts a string to a nullable sql.NullString.
func nullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

// nullTime converts a *time.Time to a nullable sql.NullString.
func nullTime(t *time.Time) sql.NullString {
	if t == nil || t.IsZero() {
		return sql.NullString{}
	}
	return sql.NullString{String: t.Format("2006-01-02T15:04:05Z07:00"), Valid: true}
}

// parseNullTime parses a nullable time column.
func parseNullTime(ns sql.NullString) *time.Time {
	if !ns.Valid || ns.String == "" {
		return nil
	}
	t := parseTime(ns.String)
	if t.IsZero() {
		return nil
	}
	return &t
}

// parseNullString extracts the string from a nullable column.
func parseNullString(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}

// newID generates a UUID v4 string.
func newID() string {
	return uuid.NewString()
}

// timeNow returns the current UTC time.
func timeNow() time.Time {
	return time.Now().UTC()
}
