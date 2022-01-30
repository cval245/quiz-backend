-- +goose Up
-- +goose StatementBegin
ALTER table question
    add column quiz_id int;

ALTER table question
    add constraint fk_question
        foreign key (quiz_id)
            references quiz (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
