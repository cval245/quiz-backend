-- +goose Up
-- +goose StatementBegin
ALTER table answer
    add column correct bool;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER table answer
    drop column correct;
-- +goose StatementEnd
