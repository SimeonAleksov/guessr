// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package database

import (
	"time"
)

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
}
