-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    id              SERIAL PRIMARY KEY,
    email           VARCHAR(100) UNIQUE NOT NULL,
    username        VARCHAR(50) UNIQUE  NOT NULL,
    password        VARCHAR(255)        NOT NULL,
    level           INTEGER     NOT NULL DEFAULT 1
);
