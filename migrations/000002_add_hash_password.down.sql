BEGIN;

ALTER TABLE "users" DROP COLUMN "hash_password";

COMMIT;