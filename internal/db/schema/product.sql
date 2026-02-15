CREATE TABLE if NOT EXISTS product (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    description TEXT,
    price_in_cents BIGINT NOT NULL CHECK (price_in_cents >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);