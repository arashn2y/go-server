-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE "role" (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE "permission" (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS "permission";
DROP TABLE IF EXISTS "role";
-- +goose StatementEnd
