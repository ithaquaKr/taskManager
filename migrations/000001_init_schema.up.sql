BEGIN;

DROP TABLE IF EXISTS "users";

DROP TABLE IF EXISTS "user_lists";

DROP TABLE IF EXISTS "lists";

DROP TABLE IF EXISTS "tasks";

DROP TABLE IF EXISTS "lists";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE
    "users" (
        "id"            UUID PRIMARY KEY                        DEFAULT (UUID_generate_v4()),
        "username"      VARCHAR(50) UNIQUE          NOT NULL    CHECK (username <> ''),
        "email"         VARCHAR(255) UNIQUE         NOT NULL    CHECK (email <> ''),
        "password"      VARCHAR(255)                NOT NULL    CHECK (octet_length(password) <> 0),
        "first_name"    VARCHAR(255)                NOT NULL    CHECK (first_name <> ''),
        "last_name"     VARCHAR(255)                NOT NULL    CHECK (last_name <> ''),
        "created_at"    TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT (NOW()),
        "updated_at"    TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    "user_lists" (
        "id"            UUID PRIMARY KEY                        DEFAULT (UUID_generate_v4()),
        "user_id"       UUID                        NOT NULL,
        "list_id"       UUID                        NOT NULL,
        "created_at"    TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT (NOW()),
        "updated_at"    TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    "lists" (
        "id"            UUID PRIMARY KEY                        DEFAULT (UUID_generate_v4()),
        "name"          VARCHAR(50)                 NOT NULL    CHECK (name <> ''),
        "type"          VARCHAR(50)                 NOT NULL    DEFAULT 'task',
        "created_at"    TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT (NOW()),
        "updated_at"    TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    "tasks" (
        "id"            UUID PRIMARY KEY                        DEFAULT (UUID_generate_v4 ()),
        "list_id"       UUID                        NOT NULL,
        "title"         VARCHAR(100)                NOT NULL    CHECK (title <> ''),
        "description"   VARCHAR(200)                            CHECK (description <> ''),
        "status"        VARCHAR(50)                 NOT NULL    DEFAULT 'doing',
        "tag"           VARCHAR(50),
        "priority"      VARCHAR(50)                 NOT NULL    DEFAULT 'no_priority',
        "due_date"      TIMESTAMP WITH TIME ZONE,
        "created_at"    TIMESTAMP WITH TIME ZONE                DEFAULT (NOW()),
        "updated_at"    TIMESTAMP WITH TIME ZONE                DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    "notes" (
        "id"            UUID PRIMARY KEY                        DEFAULT (UUID_generate_v4 ()),
        "list_id"       UUID                        NOT NULL,
        "title"         VARCHAR(100)                NOT NULL,
        "content"       TEXT,
        "reminder"      TIMESTAMP WITH TIME ZONE,
        "created_at"    TIMESTAMP WITH TIME ZONE                DEFAULT (NOW()),
        "updated_at"    TIMESTAMP WITH TIME ZONE                DEFAULT CURRENT_TIMESTAMP
    );

CREATE INDEX ON "user_lists" ("user_id", "list_id");

ALTER TABLE "user_lists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_lists" ADD FOREIGN KEY ("list_id") REFERENCES "lists" ("id");

ALTER TABLE "tasks" ADD FOREIGN KEY ("list_id") REFERENCES "lists" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("list_id") REFERENCES "lists" ("id");

COMMIT;