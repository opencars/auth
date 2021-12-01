BEGIN;

DROP INDEX tokens_user_id_idx;
DROP INDEX tokens_secret_idx;

ALTER TABLE tokens DROP COLUMN "user_id";
ALTER TABLE tokens DROP COLUMN "id";
ALTER TABLE tokens DROP COLUMN "updated_at";

ALTER TABLE tokens RENAME COLUMN "secret" TO "id";

CREATE UNIQUE INDEX tokens_id_idx ON tokens("id");

END;