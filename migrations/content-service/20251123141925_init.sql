-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id          UUID        NOT NULL    PRIMARY KEY     DEFAULT uuid_generate_v4(),
    created_at  TIMESTAMP   NOT NULL                    DEFAULT (now() at time zone 'utc'),
    updated_at  TIMESTAMP   NOT NULL                    DEFAULT (now() at time zone 'utc'),
    deleted_at  TIMESTAMP,
    name        TEXT        NOT NULL,
    username    TEXT        NOT NULL,
    email       TEXT,
    phone       TEXT
);

CREATE UNIQUE INDEX IF NOT EXISTS users_username_idx on users (username) WHERE (deleted_at IS NULL);
CREATE INDEX IF NOT EXISTS users_created_at_idx on users (created_at);

CREATE TABLE IF NOT EXISTS posts (
    id          UUID        NOT NULL    PRIMARY KEY     DEFAULT uuid_generate_v4(),
    created_at  TIMESTAMP   NOT NULL                    DEFAULT (now() at time zone 'utc'),
    updated_at  TIMESTAMP   NOT NULL                    DEFAULT (now() at time zone 'utc'),
    deleted_at  TIMESTAMP,
    author_id   UUID        NOT NULL,
    title       TEXT        NOT NULL,
    body        TEXT        NOT NULL
);

CREATE INDEX IF NOT EXISTS posts_author_idx on posts (author_id) WHERE (deleted_at IS NULL);

CREATE TABLE IF NOT EXISTS comments (
    id          UUID        NOT NULL    PRIMARY KEY     DEFAULT uuid_generate_v4(),
    created_at  TIMESTAMP   NOT NULL                    DEFAULT (now() at time zone 'utc'),
    updated_at  TIMESTAMP   NOT NULL                    DEFAULT (now() at time zone 'utc'),
    deleted_at  TIMESTAMP,
    post_id     UUID        NOT NULL,
    author_id   UUID        NOT NULL,
    body        TEXT        NOT NULL
);

CREATE INDEX IF NOT EXISTS comments_author_idx on comments (author_id) WHERE (deleted_at IS NULL);
CREATE INDEX IF NOT EXISTS comments_post_idx on comments (post_id) WHERE (deleted_at IS NULL);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users CASCADE;
DROP TABLE posts CASCADE;
DROP TABLE comments CASCADE;
-- +goose StatementEnd
