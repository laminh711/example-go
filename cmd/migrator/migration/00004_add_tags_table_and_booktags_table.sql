-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."tags" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" text NOT NULL,
  CONSTRAINT "tags_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TABLE "public"."bookstags" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "book_id" uuid NOT NULL,
  "tag_id" uuid NOT NULL,
  CONSTRAINT "bookstags_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "bookstags_book_fkey" FOREIGN KEY ("book_id") REFERENCES "public"."books" ("id"),
  CONSTRAINT "bookstags_tag_fkey" FOREIGN KEY ("tag_id") REFERENCES "public"."tags" ("id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."bookstags";
DROP TABLE "public"."tags";