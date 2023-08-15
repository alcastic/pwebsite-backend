
CREATE TABLE "messages" (
  "id" serial PRIMARY KEY,
  "remote_addr" varchar NOT NULL,
  "content" varchar NOT NULL,
  "author_name" varchar NOT NULL,
  "author_email" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);