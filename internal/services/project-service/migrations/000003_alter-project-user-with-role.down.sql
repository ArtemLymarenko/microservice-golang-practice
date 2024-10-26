ALTER TABLE projects_users DROP COLUMN IF EXISTS role;

DROP TYPE IF EXISTS project_role;