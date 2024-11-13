CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- gifts

CREATE TYPE gift_status AS ENUM('PENDING', 'PAID', 'APPROVED', 'CANCELLING', 'CANCELLED');

CREATE TABLE "gifts" (
    "id" uuid PRIMARY KEY,
    "gifter" varchar NOT NULL,
    "recipient" varchar NOT NULL,
    "message" varchar NOT NULL,
    "status" gift_status NOT NULL
);

-- transactions

CREATE TYPE transaction_status AS ENUM('PENDING', 'PAID', 'FAILED', 'CANCELLED');

CREATE TABLE "transactions" (
    "id" uuid PRIMARY KEY,
    "gift_id" uuid NOT NULL,
    "amount" numeric(10, 2) NOT NULL,
    "date" timestamptz NOT NULL,
    "status" transaction_status NOT NULL,
    FOREIGN KEY (gift_id) REFERENCES gifts(id)
);
