CREATE TYPE project_role AS ENUM ('owner', 'dev', 'qa');

ALTER TABLE projects_users
ADD COLUMN role project_role;