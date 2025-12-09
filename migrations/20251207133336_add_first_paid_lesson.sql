-- +goose Up
-- +goose StatementBegin
alter table conducted_lessons
    add column is_first_paid_lesson bool default false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table conducted_lessons
    drop column is_first_paid_lesson;
-- +goose StatementEnd
