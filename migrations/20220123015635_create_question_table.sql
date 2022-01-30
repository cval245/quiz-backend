-- +goose Up
-- +goose StatementBegin
create table question(
    id serial,
    question_text text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table question;
-- +goose StatementEnd
