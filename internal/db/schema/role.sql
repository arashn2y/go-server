CREATE TABLE "role" (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE user_role (
    user_id UUID REFERENCES "user"(id) ON DELETE CASCADE,
    role_id INT REFERENCES "role"(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);