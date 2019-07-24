-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."books" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" text,
  "author" text,
  "description" text,
  "category_id" uuid NOT NULL ,
  CONSTRAINT "books_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "category_fkey" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id")
) WITH (oids = false);  

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.