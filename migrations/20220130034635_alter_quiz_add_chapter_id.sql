-- +goose Up
-- +goose StatementBegin
alter table quiz
    add column chapter_id int;

ALTER table quiz
    add constraint fk_quiz_chapter
        foreign key (chapter_id)
            references chapter (id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table quiz drop column chapter_id;
-- +goose StatementEnd
