-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."booklends" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "book_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "from" timestamptz NOT NULL,
  "to" timestamptz NOT NULL,
  CONSTRAINT "booklends_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "booklends_book_fkey" FOREIGN KEY ("book_id") REFERENCES "public"."books" ("id"),
  CONSTRAINT "booklends_user_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id")
) WITH (oids = false);  

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."booklends"