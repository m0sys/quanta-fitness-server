-- name: StoreWorkoutLog :exec
INSERT INTO workout_logs (
    id,
    aid,
    title,
    date,
    current_pos,
    completed
    ) VALUES (
    $1, $2, $3, $4, $5, $6);

-- name: StoreExerciseLog :exec
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
    $1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: StoreSetLog :exec
INSERT INTO set_logs (
    id,
    elid,
    actual_rep_count,
    duration
    ) VALUES (
    $1, $2, $3, $4);

-- name: RemoveWorkoutLog :exec
DELETE FROM workout_logs
WHERE id = $1;

-- name: RemoveExerciseLog :exec
DELETE FROM exercise_logs
WHERE id = $1;

-- name: RemoveSetLog :exec
DELETE FROM set_logs
WHERE id = $1;

-- name: FindAllWorkoutLogsForAthlete :many
SELECT * FROM workout_logs
WHERE aid = $1;

-- name: FindAllExerciseLogsForWorkoutLog :many
SELECT * FROM exercise_logs
WHERE wlid = $1;

-- name: FindAllSetLogsForExerciseLog :many
SELECT * FROM set_logs
WHERE elid = $1;

-- name: FindWorkoutLogByID :one
SELECT * FROM workout_logs
WHERE id = $1 LIMIT 1;

-- name: FindExerciseLogByID :one
SELECT * FROM exercise_logs
WHERE id = $1 LIMIT 1;

-- name: UpdateWorkoutLog :exec
UPDATE workout_logs
SET current_pos = $2,
    completed = $3,
    updated_at = now()
WHERE id = $1;
