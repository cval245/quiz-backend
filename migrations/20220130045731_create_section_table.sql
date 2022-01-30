-- +goose Up
-- +goose StatementBegin
create table section
(
    id         serial,
    name       varchar(255),
    chapter_id int
);

alter table section
    add unique (id);

ALTER table section
    add constraint fk_chapter
        foreign key (chapter_id)
            references chapter (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table section;
-- +goose StatementEnd
