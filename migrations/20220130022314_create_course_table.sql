-- +goose Up
-- +goose StatementBegin
create table course
(
    id   serial,
    name varchar(255)
);
alter table course
    add unique (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table course;
-- +goose StatementEnd
