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
  
  ('Create Permission', '/permissions', 'POST'),
  ('Get All Permissions', '/permissions', 'GET'),
  
  ('Assign Permission to Role', '/role-permissions/:id', 'POST'),
  ('List All Role Permissions', '/role-permissions', 'GET'),
  ('Get Role Permission by UserID', '/role-permissions/:id', 'GET'),
  ('List Role Permissions by RoleID', '/role-permissions/:id/permissions', 'GET'),
  ('Delete Role Permission', '/role-permissions/:id', 'DELETE'),

  ('Get All Roles', '/roles', 'GET'),
  ('Create Role', '/roles', 'POST'),
  ('Get Role By ID', '/roles/:id', 'GET'),
  ('Update Role', '/roles/:id', 'PUT'),
  ('Delete Role', '/roles/:id', 'DELETE'),
  ('Find Role By Name', '/roles/name/:name', 'GET'),

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
  ('Get One User', '/users/:id', 'GET')
ON CONFLICT (path, method) DO NOTHING;
