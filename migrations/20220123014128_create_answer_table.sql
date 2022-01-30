-- +goose Up
-- +goose StatementBegin
CREATE TABLE answer (
    id serial,
    answer_text text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table answer;
-- +goose StatementEnd
