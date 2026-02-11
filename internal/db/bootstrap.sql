SELECT format(
    'CREATE DATABASE %I OWNER %I',
    :'db_name',
    :'db_user'
)
WHERE NOT EXISTS (
    SELECT 1 FROM pg_database WHERE datname = :'db_name'
)\gexec

\connect :db_name

CREATE TABLE IF NOT EXISTS users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE users OWNER TO :db_user;
ALTER ROLE :db_user SET search_path = public;
