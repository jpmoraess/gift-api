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

-- transactions

CREATE TYPE transaction_status AS ENUM('PENDING', 'PAID', 'FAILED', 'CANCELLED');

CREATE TABLE "transactions" (
    "id" uuid PRIMARY KEY,
    "external_id" varchar NOT NULL,
    "amount" numeric(10, 2) NOT NULL,
    "date" timestamptz NOT NULL,
    "status" transaction_status NOT NULL
);

-- sessions

CREATE TABLE "sessions" (
    "id" uuid PRIMARY KEY,
    "username" varchar NOT NULL,
    "refresh_token" varchar NOT NULL,
    "user_agent" varchar NOT NULL,
    "client_ip" varchar NOT NULL,
    "is_blocked" boolean NOT NULL DEFAULT false,
    "expires_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");