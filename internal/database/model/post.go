package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Post struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

func CreatePostTable(db *sqlx.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS posts
(
    id         uuid        NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    title      text        NOT NULL,
    content    text        NOT NULL,
    created_at timestamptz NOT NULL        DEFAULT now(),
    updated_at timestamptz NOT NULL        DEFAULT now(),

    CONSTRAINT pk_post PRIMARY KEY (id)
);`)

	return err
}
