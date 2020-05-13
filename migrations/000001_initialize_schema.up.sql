CREATE TABLE tokens(
    "id"            VARCHAR(32) NOT NULL,
    "name"          TEXT        NOT NULL,
    "enabled"       BOOLEAN     NOT NULL DEFAULT TRUE,
    "created_at"    TIMESTAMP   NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX tokens_id_idx ON tokens("id");