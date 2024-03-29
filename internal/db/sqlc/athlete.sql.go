// Code generated by sqlc. DO NOT EDIT.
// source: athlete.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createAthlete = `-- name: CreateAthlete :one
INSERT INTO athletes (id) VALUES ($1) RETURNING id, created_at, updated_at
`

func (q *Queries) CreateAthlete(ctx context.Context, id uuid.UUID) (Athlete, error) {
	row := q.db.QueryRowContext(ctx, createAthlete, id)
	var i Athlete
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}
