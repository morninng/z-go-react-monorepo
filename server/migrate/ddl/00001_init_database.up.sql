
CREATE TABLE "tasks" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "content" varchar NOT NULL,
  "created_at" timestamp(6) NOT NULL DEFAULT current_timestamp,
  "updated_at" timestamp(6) NOT NULL DEFAULT current_timestamp
);
