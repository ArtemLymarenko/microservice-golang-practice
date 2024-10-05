CREATE TYPE project_status AS ENUM ('active', 'closed', 'idle');

CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    description TEXT NULL,
    status project_status NOT NULL DEFAULT 'idle',
    production_start_at TIMESTAMPTZ NULL,
    production_end_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    archived_at TIMESTAMPTZ NULL
);

CREATE INDEX IF NOT EXISTS id_idx ON projects(id);
CREATE INDEX IF NOT EXISTS name_status_idx ON projects(name, status);