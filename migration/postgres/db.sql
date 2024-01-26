CREATE TYPE "payment_type" AS ENUM (
  'card',
  'cash'
);

CREATE TYPE "status_type" AS ENUM (
  'in_process',
  'success',
  'cancel'
);

CREATE TYPE "transaction_type" AS ENUM (
  'withdraw',
  'topup'
);

CREATE TYPE "source_type" AS ENUM (
  'bonus',
  'sales'
);

CREATE TYPE "tarif_type" AS ENUM (
  'percent',
  'fixed'
);

CREATE TABLE "branch" (
  "id" uuid PRIMARY KEY,
  "name" varchar(30),
  "address" varchar(50),
  "update_at" timestamp DEFAULT (NOW()),
  "created_at" timestamp DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "sale" (
  "id" uuid PRIMARY KEY,
  "branch_id" uuid,
  "shop_assistant_id" varchar(70),
  "cashier" uuid,
  "payment_type" payment_type,
  "price" numeric,
  "status" status_type,
  "client_name" varchar(30),
  "update_at" timestamp DEFAULT (NOW()),
  "created_at" timestamp DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "staff" (
  "id" uuid PRIMARY KEY,
  "tarif_id" uuid,
  "branch_id" uuid,
  "update_at" timestamp DEFAULT (NOW()),
  "created_at" timestamp DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "transaction" (
  "id" uuid PRIMARY KEY,
  "sale_id" uuid,
  "staff_id" uuid,
  "transaction_type" transaction_type,
  "source_type" source_type,
  "amount" numeric,
  "description" text,
  "update_at" timestamp DEFAULT (NOW()),
  "created_at" timestamp DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "staff_tarif" (
  "id" uuid PRIMARY KEY,
  "name" varchar(30),
  "tarif_type" tarif_type,
  "amount_for_cash" numeric,
  "amount_for_card" numeric,
  "update_at" timestamp DEFAULT (NOW()),
  "created_at" timestamp DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT null
);

ALTER TABLE "sale" ADD FOREIGN KEY ("branch_id") REFERENCES "branch" ("id");

ALTER TABLE "staff" ADD FOREIGN KEY ("tarif_id") REFERENCES "staff_tarif" ("id");

ALTER TABLE "staff" ADD FOREIGN KEY ("branch_id") REFERENCES "branch" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("sale_id") REFERENCES "sale" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("staff_id") REFERENCES "staff" ("id");