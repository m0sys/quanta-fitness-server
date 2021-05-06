-- name: CreateAthlete :one
INSERT INTO athletes (id) VALUES ($1) RETURNING *;
