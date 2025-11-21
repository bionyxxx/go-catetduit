-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions
(
    id               SERIAL PRIMARY KEY,
    user_id          INT            NOT NULL,
    amount           DECIMAL(10, 2) NOT NULL,
    type             VARCHAR(50)    NOT NULL,
    description      TEXT           NULL,
    transaction_date TIMESTAMP      NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd
