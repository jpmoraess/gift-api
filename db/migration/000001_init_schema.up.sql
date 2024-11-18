CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- files
CREATE TABLE "files" (
    "id" uuid PRIMARY KEY,
    "name" varchar NOT NULL,
    "extension" varchar NOT NULL,
    "size" integer NOT NULL,
    "path" varchar NOT NULL
);

-- users

CREATE TABLE "users" (
    "id" uuid PRIMARY KEY,
    "username" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("email");
CREATE INDEX ON "users" ("username");

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
    "external_id" varchar NOT NULL,
    "amount" numeric(10, 2) NOT NULL,
    "date" timestamptz NOT NULL,
    "status" transaction_status NOT NULL,
    FOREIGN KEY (gift_id) REFERENCES gifts(id)
);
