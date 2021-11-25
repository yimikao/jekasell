CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()) 
);

CREATE TABLE "products" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "quantity" bigint NOT NULL,
    "description" varchar NOT NULL,
    "avatar_url" varchar NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_products" (
    "order_id" bigint NOT NULL,
    "product_id" bigint NOT NULL,
    "quantity" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);
