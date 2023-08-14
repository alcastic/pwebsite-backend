
CREATE TABLE "messages" (
  "id" serial PRIMARY KEY,
  "remote_addr" varchar NOT NULL,
  "content" varchar,
  "author_name" varchar,
  "author_email" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);