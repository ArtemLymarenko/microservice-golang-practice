CREATE TABLE IF NOT EXISTS users (
    id uuid primary key,
    email text not null unique,
    password text not null
);