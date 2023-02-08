CREATE TABLE "sales" (
  "id" bigserial PRIMARY KEY,
  "point_of_sale" varchar NOT NULL,
  "product" varchar NOT NULL,
  "date" timestamptz NOT NULL,
  "stock" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);