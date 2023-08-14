// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"database/sql"
	"time"
)

type Message struct {
	ID          int32          `json:"id"`
	RemoteAddr  string         `json:"remoteAddr"`
	Content     sql.NullString `json:"content"`
	AuthorName  sql.NullString `json:"authorName"`
	AuthorEmail sql.NullString `json:"authorEmail"`
	CreatedAt   time.Time      `json:"createdAt"`
}
