-- Down
DROP TABLE IF EXISTS "users";

-- Up
CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL DEFAULT '',
    "email" VARCHAR(255) NOT NULL,
    "password" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE("email")
);