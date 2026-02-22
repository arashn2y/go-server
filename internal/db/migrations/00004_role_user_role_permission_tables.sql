-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE user_role (
    user_id UUID REFERENCES "user"(id) ON DELETE CASCADE,
    role_id INT REFERENCES "role"(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

CREATE TABLE role_permission (
    id SERIAL PRIMARY KEY,
    role_id INT NOT NULL REFERENCES "role"(id) ON DELETE CASCADE,
    resource TEXT NOT NULL,
    permission_id INT NOT NULL REFERENCES "permission"(id) ON DELETE CASCADE,
    UNIQUE(role_id, resource, permission_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS role_permission;
DROP TABLE IF EXISTS user_role;
-- +goose StatementEnd
