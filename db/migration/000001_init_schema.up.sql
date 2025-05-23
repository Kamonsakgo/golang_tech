
CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL,
  "username" text NOT NULL,
  "balance" numeric(12,2) NOT NULL DEFAULT 0,
  "affiliate_id" uuid
);

CREATE TABLE "product" (
  "id" uuid PRIMARY KEY NOT NULL,
  "name" text NOT NULL,
  "quantity" integer NOT NULL DEFAULT 0,
  "price" numeric(12,2) NOT NULL
);

CREATE TABLE "commission" (
  "id" uuid PRIMARY KEY NOT NULL,
  "order_id" uuid NOT NULL,
  "affiliate_id" uuid NOT NULL,
  "amount" numeric(12,2) NOT NULL
);

CREATE TABLE "affiliate" (
  "id" uuid PRIMARY KEY NOT NULL,
  "name" text NOT NULL,
  "master_affiliate" uuid,
  "balance" numeric(12,2) NOT NULL DEFAULT 0
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("affiliate_id");

CREATE INDEX ON "product" ("name");

CREATE INDEX ON "commission" ("affiliate_id");

CREATE INDEX ON "commission" ("order_id");

CREATE INDEX ON "affiliate" ("name");

CREATE INDEX ON "affiliate" ("master_affiliate");

ALTER TABLE "users" ADD FOREIGN KEY ("affiliate_id") REFERENCES "affiliate" ("id");

ALTER TABLE "commission" ADD FOREIGN KEY ("affiliate_id") REFERENCES "affiliate" ("id");

ALTER TABLE "affiliate" ADD FOREIGN KEY ("master_affiliate") REFERENCES "affiliate" ("id");
-- เปิด extension ถ้ายังไม่เปิด
CREATE EXTENSION IF NOT EXISTS pgcrypto;

ALTER TABLE users
  ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE product
  ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE affiliate
  ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE commission
  ALTER COLUMN id SET DEFAULT gen_random_uuid();
