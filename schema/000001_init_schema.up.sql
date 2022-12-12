CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS account(
    id UUID PRIMARY KEY  DEFAULT uuid_generate_v4 (),
    username VARCHAR(255) NOT NULL UNIQUE,
    password  VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

INSERT INTO account( username,password) VALUES('username','password');
