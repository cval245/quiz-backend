-- +goose Up
-- +goose StatementBegin
ALTER table answer
    add column question_id int;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER table answer
    drop column question_id;
-- +goose StatementEnd
