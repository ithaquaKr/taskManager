CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar(50) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "first_name" varchar(255) NOT NULL,
  "last_name" varchar(255) NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "list_users" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "list_id" uuid,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "list" (
  "id" uuid PRIMARY KEY,
  "name" varchar(50),
  "type" varchar(50) NOT NULL DEFAULT 'task',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "task" (
  "id" uuid PRIMARY KEY,
  "list_id" uuid NOT NULL,
  "title" varchar(100) NOT NULL,
  "description" varchar(200),
  "status" varchar(50) NOT NULL DEFAULT 'doing',
  "tag" varchar(50),
  "priority" varchar(50) NOT NULL DEFAULT 'no_priority',
  "due_date" timestamptz,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "note" (
  "id" uuid PRIMARY KEY,
  "list_id" uuid NOT NULL,
  "title" varchar(100) NOT NULL,
  "content" text,
  "reminder" timestamptz,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE UNIQUE INDEX ON "users" ("username");

CREATE UNIQUE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("created_at");

CREATE INDEX ON "list_users" ("created_at");

CREATE INDEX ON "list_users" ("user_id", "list_id");

CREATE INDEX ON "list" ("created_at");

CREATE INDEX ON "task" ("status");

CREATE INDEX ON "task" ("tag");

CREATE INDEX ON "task" ("created_at");

CREATE INDEX ON "note" ("created_at");

ALTER TABLE "list_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "list_users" ADD FOREIGN KEY ("list_id") REFERENCES "list" ("id");

ALTER TABLE "task" ADD FOREIGN KEY ("list_id") REFERENCES "list" ("id");

ALTER TABLE "note" ADD FOREIGN KEY ("list_id") REFERENCES "list" ("id");
