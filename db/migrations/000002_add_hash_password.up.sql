BEGIN;

ALTER TABLE "users" ADD COLUMN "hash_password" varchar NOT NULL;

COMMIT;