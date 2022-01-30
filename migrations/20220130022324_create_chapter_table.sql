-- +goose Up
-- +goose StatementBegin
create table chapter
(
    id   serial,
    name varchar(255),
    course_id int
);
alter table chapter
    add unique (id);
ALTER table chapter
    add constraint fk_course
        foreign key (course_id)
            references course (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chapter
-- +goose StatementEnd
