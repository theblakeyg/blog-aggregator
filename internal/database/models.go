// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.NullUUID
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Name      sql.NullString
	Url       sql.NullString
	UserID    uuid.NullUUID
}

type User struct {
	ID        uuid.NullUUID
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Name      sql.NullString
}
