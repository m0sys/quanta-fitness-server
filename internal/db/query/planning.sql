-- name: StoreWorkoutPlan :exec
INSERT INTO workout_plans (
    id,
    aid,
    title
    ) VALUES (
    $1, $2, $3);

-- name: FindWorkoutPlanByTitleAndAthleteID :one
SELECT * FROM workout_plans
WHERE title = $1 AND aid = $2 LIMIT 1;

-- name: FindWorkoutPlanByID :one
SELECT * FROM workout_plans
WHERE id = $1 LIMIT 1;

-- name: FindWorkoutPlanByIDAndAthleteID :one
SELECT * FROM workout_plans
WHERE id = $1 AND aid = $2 LIMIT 1;

-- name: StoreExercise :exec
INSERT INTO exercises (
    id,
    aid,
    wpid,
    target_rep,
    num_sets,
    weight,
    rest_duration
    ) VALUES (
    $1, $2, $3, $4, $5, $6, $7);

-- name: FindExerciseByID :one
SELECT * FROM exercises
WHERE id = $1 LIMIT 1;

-- name: FindExerciseByNameAndWorkoutPlanID :one
SELECT * FROM exercises
WHERE name = $1 AND wpid = $2 LIMIT 1;


-- name: RemoveExercise :exec
DELETE FROM exercises
WHERE id = $1;

-- name: UpdateWorkoutPlan :exec
UPDATE workout_plans
SET title = $2,
    updated_at = now()
WHERE id = $1;

-- name: FindAllWorkoutPlansForAthlete :many
SELECT * FROM workout_plans
WHERE aid = $1;

-- name: FindAllExercisesForWorkoutPlan :many
SELECT * FROM exercises
WHERE wpid = $1;

-- name: UpdateExercise :exec
UPDATE exercises
SET name = $2,
    target_rep = $3,
    num_sets = $4,
    weight = $5,
    rest_duration = $6
WHERE id = $1;

-- name: RemoveWorkoutPlan :exec
DELETE FROM workout_plans
WHERE id = $1;
