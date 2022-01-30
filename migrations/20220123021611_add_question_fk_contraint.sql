-- +goose Up
-- +goose StatementBegin
ALTER table question
    add unique (id);

ALTER table answer
    add constraint fk_question
        foreign key (question_id)
            references question (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table answer
    drop constraint fk_question;
-- +goose StatementEnd
