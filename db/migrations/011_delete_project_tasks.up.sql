ALTER TABLE tasks
DROP CONSTRAINT IF EXISTS tasks_project_id_fkey;

ALTER TABLE tasks
ADD CONSTRAINT tasks_project_id_fkey
FOREIGN KEY (project_id)
REFERENCES projects(id)
ON DELETE CASCADE;
