CREATE TYPE project_role AS ENUM ('owner', 'member');

ALTER TABLE projects_users
ADD COLUMN role project_role DEFAULT 'member';