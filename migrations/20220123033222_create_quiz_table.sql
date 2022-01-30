-- +goose Up
-- +goose StatementBegin
create table quiz(
    id serial,
    name varchar(255)
);
alter table quiz add unique (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
