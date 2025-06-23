CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS currencies
(
    currency_code VARCHAR(10) PRIMARY KEY,
    currency_name VARCHAR(100) NOT NULL
);

DROP TABLE IF EXISTS currencies;
