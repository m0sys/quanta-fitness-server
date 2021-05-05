CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "athletes" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "workout_plans" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "aid" uuid,
    "title" VARCHAR NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "exercises" (
    "id" uuid  PRIMARY KEY DEFAULT uuid_generate_v4(),
    "wpid" uuid,
    "aid" uuid,
    "name" VARCHAR NOT NULL,
    "target_rep" int NOT NULL,
    "num_sets" int NOT NULL,
    "weight" float8 NOT NULL,
    "rest_duration" float8 NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "workout_plans" ADD FOREIGN KEY ("aid") REFERENCES "athletes" ("id");
ALTER TABLE "exercises" ADD FOREIGN KEY ("aid") REFERENCES "athletes" ("id");
ALTER TABLE "exercises" ADD FOREIGN KEY ("wpid") REFERENCES "workout_plans" ("id");
