CREATE TABLE "permission" (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE role_permission (
    id SERIAL PRIMARY KEY,
    role_id INT NOT NULL REFERENCES "role"(id) ON DELETE CASCADE,
    resource TEXT NOT NULL,
    permission_id INT NOT NULL REFERENCES "permission"(id) ON DELETE CASCADE,
    UNIQUE(role_id, resource, permission_id)
);