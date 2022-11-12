-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS phones (
    id int primary key,
    data VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS phones;
-- +goose StatementEnd
