BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

ALTER TABLE tokens RENAME COLUMN "id" TO "secret";
ALTER TABLE tokens ADD COLUMN "id" uuid DEFAULT uuid_generate_v4();

ALTER TABLE tokens ADD COLUMN "updated_at" TIMESTAMP;
ALTER TABLE tokens ADD PRIMARY KEY ("id");

DROP INDEX tokens_id_idx;

ALTER TABLE tokens ADD COLUMN "user_id" uuid;
UPDATE tokens SET "user_id" = 'a8b6457e-d009-44df-86ea-d28b8b61eb42';

UPDATE tokens SET "updated_at" = "created_at";

ALTER TABLE tokens ALTER COLUMN "user_id" SET NOT NULL;
ALTER TABLE tokens ALTER COLUMN "updated_at" SET NOT NULL;

CREATE UNIQUE INDEX tokens_secret_idx ON tokens("secret");
CREATE INDEX tokens_user_id_idx ON tokens("user_id");

END;