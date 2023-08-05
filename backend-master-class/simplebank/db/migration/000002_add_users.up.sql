CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- Create a unique index with composite keys 
-- on accounts(owner, currency) to prevent duplicate accounts with the same owner and currency
-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");

-- Another way to create a unique index with composite keys it can use the unique constraint to prevent duplicate accounts with the same owner and currency
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner","currency");
