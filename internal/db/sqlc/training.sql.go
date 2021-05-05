// Code generated by sqlc. DO NOT EDIT.
// source: training.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const findAllExerciseLogsForWorkoutLog = `-- name: FindAllExerciseLogsForWorkoutLog :many
SELECT id, wlid, name, target_rep, num_sets, weight, rest_duration, completed, pos, created_at, updated_at FROM exercise_logs
WHERE wlid = $1
`

func (q *Queries) FindAllExerciseLogsForWorkoutLog(ctx context.Context, wlid uuid.UUID) ([]ExerciseLog, error) {
	rows, err := q.db.QueryContext(ctx, findAllExerciseLogsForWorkoutLog, wlid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ExerciseLog
	for rows.Next() {
		var i ExerciseLog
		if err := rows.Scan(
			&i.ID,
			&i.Wlid,
			&i.Name,
			&i.TargetRep,
			&i.NumSets,
			&i.Weight,
			&i.RestDuration,
			&i.Completed,
			&i.Pos,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllSetLogsForExerciseLog = `-- name: FindAllSetLogsForExerciseLog :many
SELECT id, elid, actual_rep_count, duration, created_at, updated_at FROM set_logs
WHERE elid = $1
`

func (q *Queries) FindAllSetLogsForExerciseLog(ctx context.Context, elid uuid.UUID) ([]SetLog, error) {
	rows, err := q.db.QueryContext(ctx, findAllSetLogsForExerciseLog, elid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SetLog
	for rows.Next() {
		var i SetLog
		if err := rows.Scan(
			&i.ID,
			&i.Elid,
			&i.ActualRepCount,
			&i.Duration,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllWorkoutLogsForAthlete = `-- name: FindAllWorkoutLogsForAthlete :many
SELECT id, aid, title, date, current_pos, completed, created_at, updated_at FROM workout_logs
WHERE aid = $1
`

func (q *Queries) FindAllWorkoutLogsForAthlete(ctx context.Context, aid uuid.UUID) ([]WorkoutLog, error) {
	rows, err := q.db.QueryContext(ctx, findAllWorkoutLogsForAthlete, aid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []WorkoutLog
	for rows.Next() {
		var i WorkoutLog
		if err := rows.Scan(
			&i.ID,
			&i.Aid,
			&i.Title,
			&i.Date,
			&i.CurrentPos,
			&i.Completed,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findExerciseLogByID = `-- name: FindExerciseLogByID :one
SELECT id, wlid, name, target_rep, num_sets, weight, rest_duration, completed, pos, created_at, updated_at FROM exercise_logs
WHERE id = $1 LIMIT 1
`

func (q *Queries) FindExerciseLogByID(ctx context.Context, id uuid.UUID) (ExerciseLog, error) {
	row := q.db.QueryRowContext(ctx, findExerciseLogByID, id)
	var i ExerciseLog
	err := row.Scan(
		&i.ID,
		&i.Wlid,
		&i.Name,
		&i.TargetRep,
		&i.NumSets,
		&i.Weight,
		&i.RestDuration,
		&i.Completed,
		&i.Pos,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findWorkoutLogByID = `-- name: FindWorkoutLogByID :one
SELECT id, aid, title, date, current_pos, completed, created_at, updated_at FROM workout_logs
WHERE id = $1 LIMIT 1
`

func (q *Queries) FindWorkoutLogByID(ctx context.Context, id uuid.UUID) (WorkoutLog, error) {
	row := q.db.QueryRowContext(ctx, findWorkoutLogByID, id)
	var i WorkoutLog
	err := row.Scan(
		&i.ID,
		&i.Aid,
		&i.Title,
		&i.Date,
		&i.CurrentPos,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const removeExerciseLog = `-- name: RemoveExerciseLog :exec
DELETE FROM exercise_logs
WHERE id = $1
`

func (q *Queries) RemoveExerciseLog(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, removeExerciseLog, id)
	return err
}

const removeSetLog = `-- name: RemoveSetLog :exec
DELETE FROM set_logs
WHERE id = $1
`

func (q *Queries) RemoveSetLog(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, removeSetLog, id)
	return err
}

const removeWorkoutLog = `-- name: RemoveWorkoutLog :exec
DELETE FROM workout_logs
WHERE id = $1
`

func (q *Queries) RemoveWorkoutLog(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, removeWorkoutLog, id)
	return err
}

const storeExerciseLog = `-- name: StoreExerciseLog :exec
INSERT INTO exercise_logs (
    id,
    wlid,
    name,
    target_rep,
    num_sets,
    weight,
    rest_duration,
    completed ,
    pos
    ) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9)
`

type StoreExerciseLogParams struct {
	ID           uuid.UUID `json:"id"`
	Wlid         uuid.UUID `json:"wlid"`
	Name         string    `json:"name"`
	TargetRep    int32     `json:"target_rep"`
	NumSets      int32     `json:"num_sets"`
	Weight       float64   `json:"weight"`
	RestDuration float64   `json:"rest_duration"`
	Completed    bool      `json:"completed"`
	Pos          int32     `json:"pos"`
}

func (q *Queries) StoreExerciseLog(ctx context.Context, arg StoreExerciseLogParams) error {
	_, err := q.db.ExecContext(ctx, storeExerciseLog,
		arg.ID,
		arg.Wlid,
		arg.Name,
		arg.TargetRep,
		arg.NumSets,
		arg.Weight,
		arg.RestDuration,
		arg.Completed,
		arg.Pos,
	)
	return err
}

const storeSetLog = `-- name: StoreSetLog :exec
INSERT INTO set_logs (
    id,
    elid,
    actual_rep_count,
    duration
    ) VALUES (
    $1, $2, $3, $4)
`

type StoreSetLogParams struct {
	ID             uuid.UUID `json:"id"`
	Elid           uuid.UUID `json:"elid"`
	ActualRepCount int32     `json:"actual_rep_count"`
	Duration       float64   `json:"duration"`
}

func (q *Queries) StoreSetLog(ctx context.Context, arg StoreSetLogParams) error {
	_, err := q.db.ExecContext(ctx, storeSetLog,
		arg.ID,
		arg.Elid,
		arg.ActualRepCount,
		arg.Duration,
	)
	return err
}

const storeWorkoutLog = `-- name: StoreWorkoutLog :exec
INSERT INTO workout_logs (
    id,
    aid,
    title,
    current_pos,
    completed
    ) VALUES (
    $1, $2, $3, $4, $5)
`

type StoreWorkoutLogParams struct {
	ID         uuid.UUID `json:"id"`
	Aid        uuid.UUID `json:"aid"`
	Title      string    `json:"title"`
	CurrentPos int32     `json:"current_pos"`
	Completed  bool      `json:"completed"`
}

func (q *Queries) StoreWorkoutLog(ctx context.Context, arg StoreWorkoutLogParams) error {
	_, err := q.db.ExecContext(ctx, storeWorkoutLog,
		arg.ID,
		arg.Aid,
		arg.Title,
		arg.CurrentPos,
		arg.Completed,
	)
	return err
}

const updateWorkoutLog = `-- name: UpdateWorkoutLog :exec
UPDATE workout_logs
SET current_pos = $2,
    completed = $3,
    updated_at = now()
WHERE id = $1
`

type UpdateWorkoutLogParams struct {
	ID         uuid.UUID `json:"id"`
	CurrentPos int32     `json:"current_pos"`
	Completed  bool      `json:"completed"`
}

func (q *Queries) UpdateWorkoutLog(ctx context.Context, arg UpdateWorkoutLogParams) error {
	_, err := q.db.ExecContext(ctx, updateWorkoutLog, arg.ID, arg.CurrentPos, arg.Completed)
	return err
}