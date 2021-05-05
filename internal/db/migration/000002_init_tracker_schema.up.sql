CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "workout_logs" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "aid" uuid,
    "title" VARCHAR NOT NULL,
    "date" timestamptz NOT NULL,
    "current_pos" int NOT NULL,
    "completed" boolean NOT NULL DEFAULT FALSE,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "exercise_logs" (
    "id" uuid  PRIMARY KEY DEFAULT uuid_generate_v4(),
    "wlid" uuid,
    "name" VARCHAR NOT NULL,
    "target_rep" int NOT NULL,
    "num_sets" int NOT NULL,
    "weight" float8 NOT NULL,
    "rest_duration" float8 NOT NULL,
    "completed" boolean NOT NULL DEFAULT FALSE,
    "pos" int NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "set_logs" (
    "id" uuid  PRIMARY KEY DEFAULT uuid_generate_v4(),
    "elid" uuid,
    "actual_rep_count" int NOT NULL,
    "duration" float8 NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "workout_logs" ADD FOREIGN KEY ("aid") REFERENCES "athletes" ("id");
ALTER TABLE "exercise_logs" ADD FOREIGN KEY ("wlid") REFERENCES "workout_logs" ("id");
ALTER TABLE "set_logs" ADD FOREIGN KEY ("elid") REFERENCES "exercise_logs" ("id");
