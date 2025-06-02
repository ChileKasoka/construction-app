CREATE TABLE permissions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  path VARCHAR(255) NOT NULL,
  method VARCHAR(10) NOT NULL,
  UNIQUE(path, method)
);

-- 🌱 Seed data
INSERT INTO permissions (name, path, method) VALUES 
  ('Create Project', '/projects', 'POST'),
  ('Get All Projects', '/projects', 'GET'),
  ('Get One Project', '/projects/:id', 'GET'),
  ('Update Project', '/projects/:id', 'PUT'),
  ('Delete Project', '/projects/:id', 'DELETE'),

  ('Create Task', '/tasks', 'POST'),
  ('Get All Tasks', '/tasks', 'GET'),
  ('Update Task', '/tasks/:id', 'PUT'),
  ('Delete Task', '/tasks/:id', 'DELETE'),

  ('Assign Task to User', '/user-tasks', 'POST'),
  ('Unassign Task from User', '/user-tasks/:id', 'DELETE'),

  ('Create User', '/users', 'POST'),
  ('Update User', '/users/:id', 'PUT'),
  ('Delete User', '/users/:id', 'DELETE'),

  ('Get All Users', '/users', 'GET'),
  ('Get One User', '/users/:id', 'GET');