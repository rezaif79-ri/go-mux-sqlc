CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "salt" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "role" varchar NOT NULL,
  "profile_id" int
);

CREATE TABLE "profiles" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_bio" varchar,
  "email" varchar,
  "first_name" varchar,
  "last_name" varchar,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "items" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "weight" float NOT NULL,
  "volume" float NOT NULL,
  "stock" int NOT NULL,
  "price" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "category_id" int
);

CREATE TABLE "categories" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar,
  "description" varchar
);

CREATE TABLE "orders" (
  "id" BIGSERIAL PRIMARY KEY,
  "buyer_id" int,
  "total_price" int,
  "total_weight" float,
  "shipment_cost" int
);

CREATE TABLE "item_orders" (
  "id" BIGSERIAL PRIMARY KEY,
  "order_id" int NOT NULL,
  "item_id" int NOT NULL,
  "quantity" int NOT NULL,
  "sub_total" int NOT NULL,
  "sub_weight" float NOT NULL
);

ALTER TABLE "users" ADD FOREIGN KEY ("profile_id") REFERENCES "profiles" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("buyer_id") REFERENCES "users" ("id");

ALTER TABLE "item_orders" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "item_orders" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

CREATE INDEX ON "users" ("profile_id");

CREATE INDEX ON "items" ("category_id");

CREATE INDEX ON "orders" ("buyer_id");

CREATE INDEX ON "item_orders" ("item_id");
