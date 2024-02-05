CREATE TYPE "storagetransaction_type" AS ENUM (
  'minus',
  'plus'
);
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
CREATE TYPE "staff_type" AS ENUM (
  'shop_assistant',
  'cashier'
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

CREATE TABLE "staff" (
  "id" uuid PRIMARY KEY,
  "tarif_id" uuid,
  "branch_id" uuid,
  "name" varchar(30),
  "age" int,
  "birth_date" date,
  "login" varchar(30),
  "password" varchar(100),
  "staff_type" staff_type,
  "balance" int DEFAULT 0,
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

CREATE TABLE "category" (
  "id" uuid PRIMARY KEY,
  "parent_id" uuid,
  "name"varchar(50),
  "update_at" timestamp DEFAULT (NOW()),
  "created_at" timestamp DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "product" (
  "id" uuid PRIMARY KEY,
  "name" varchar(20),
  "price" numeric,
  "barcode" int UNIQUE,
  "category_id" uuid,
  "update_at" timestamp DEFAULT (NOW()),
  "created_at" timestamp DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "storage" (
  "id" uuid PRIMARY KEY,
  "product_id" uuid,
  "branch_id" uuid,
  "count" int,
  "update_at" timestamp DEFAULT (NOW()),  
  "created_at" timestamp DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "storagetransaction" (
  "id" uuid PRIMARY KEY,
  "branch_id" uuid,
  "staff_id" uuid,
  "product_id" uuid,
  "storagetransactiontype" storagetransaction_type,
  "price" numeric,
  "quantity" int,
  "update_at" timestamp DEFAULT (NOW()),
  "created_at" timestamp DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "basket" (
  "id" uuid PRIMARY KEY,
  "product_id" uuid,
  "quantity" int,
  "price" numeric,
  "sale_id" 
  "update_at" timestamp DEFAULT (NOW()),
  "created_at" timestamp DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT null
);

ALTER TABLE "sale" ADD FOREIGN KEY ("branch_id") REFERENCES "branch" ("id");

ALTER TABLE "staff" ADD FOREIGN KEY ("tarif_id") REFERENCES "staff_tarif" ("id");

ALTER TABLE "staff" ADD FOREIGN KEY ("branch_id") REFERENCES "branch" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("sale_id") REFERENCES "sale" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("staff_id") REFERENCES "staff" ("id");

ALTER TABLE "category" ADD FOREIGN KEY ("parent_id") REFERENCES "category" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "storage" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "storage" ADD FOREIGN KEY ("branch_id") REFERENCES "branch" ("id");

ALTER TABLE "storagetransaction" ADD FOREIGN KEY ("branch_id") REFERENCES "branch" ("id");

ALTER TABLE "storagetransaction" ADD FOREIGN KEY ("staff_id") REFERENCES "staff" ("id");

ALTER TABLE "storagetransaction" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "basket" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");
