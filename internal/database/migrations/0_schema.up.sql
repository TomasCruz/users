CREATE TABLE IF NOT EXISTS users (
    user_id               UUID NOT NULL PRIMARY KEY,
    first_name            VARCHAR(30) NOT NULL,
    last_name             VARCHAR(30) NOT NULL,
    pswd_hash             VARCHAR NOT NULL,
    email                 VARCHAR(42) NOT NULL,
    country               VARCHAR(3) NOT NULL,
    created_at            TIMESTAMP NOT NULL,
    updated_at            TIMESTAMP NOT NULL,
    UNIQUE(email)
);
