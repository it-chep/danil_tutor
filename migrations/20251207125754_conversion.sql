-- +goose Up
-- +goose StatementBegin
alter table students
    add column state smallint;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table students
    drop column state smallint;
-- +goose StatementEnd
