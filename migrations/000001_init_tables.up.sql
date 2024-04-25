CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    phone_number VARCHAR(20),
    username VARCHAR(32),
    session_string VARCHAR NOT NULL,
);

CREATE TABLE channels (
    id SERIAL PRIMARY KEY,
    channel_id BIGINT NOT NULL,
    access_hash BIGINT NOT NULL,
    user_count INTEGER NOT NULL,
    name VARCHAR NOT NULL,
);
