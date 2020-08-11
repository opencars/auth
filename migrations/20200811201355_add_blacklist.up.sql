CREATE TABLE blacklist(
    "ipv4"          TEXT        NOT NULL,
    "enabled"       BOOLEAN     NOT NULL DEFAULT TRUE,
    "created_at"    TIMESTAMP   NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX blacklist_ipv4_idx ON blacklist("ipv4");