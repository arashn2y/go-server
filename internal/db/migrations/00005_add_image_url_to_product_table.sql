-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE product
ADD COLUMN IF NOT EXISTS image_url TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE product
DROP COLUMN IF EXISTS image_url;
-- +goose StatementEnd
