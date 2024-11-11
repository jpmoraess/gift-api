CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- gifts

CREATE TYPE gift_status AS ENUM('GiftPending', 'GiftPaid', 'GiftApproved', 'GiftCancelling', 'GiftCancelled');

CREATE TABLE "gifts" (
    "id" uuid PRIMARY KEY,
    "gifter" varchar NOT NULL,
    "recipient" varchar NOT NULL,
    "message" varchar NOT NULL,
    "status" gift_status NOT NULL
);

-- payments

CREATE TYPE payment_status AS ENUM('PaymentCompleted', 'PaymentCancelled', 'PaymentFailed');

CREATE TABLE "payments" (
    "id" uuid PRIMARY KEY,
    "gift_id" uuid NOT NULL,
    "status" payment_status NOT NULL,
    FOREIGN KEY (gift_id) REFERENCES gifts(id)
);
