-- Create a table to store purchase transactions
CREATE TABLE IF NOT EXISTS purchase_transactions (
    id UUID PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    purchase_amount_usd NUMERIC(10, 2) NOT NULL,
    exchange_rate NUMERIC(10, 4) NOT NULL,
    purchase_amount NUMERIC(10, 2) NOT NULL
);
