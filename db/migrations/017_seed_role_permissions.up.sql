INSERT INTO permissions (name, path, method) VALUES
  ('Assign Permission to Role', '/role-permissions/:id', 'POST'),
  ('List All Role Permissions', '/role-permissions', 'GET'),
  ('Get Role Permission by UserID', '/role-permissions/:id', 'GET'),
  ('List Role Permissions by RoleID', '/role-permissions/:id/permissions', 'GET'),
  ('Delete Role Permission', '/role-permissions/:id', 'DELETE'),
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
  ('Get One User', '/users/:id', 'GET'),

  ('Create Permission', '/permissions', 'POST'),
  ('Get All Permissions', '/permissions', 'GET');

INSERT INTO role_permissions (role_id, permission_id)
SELECT 1, id FROM permissions
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT 6, id FROM permissions
ON CONFLICT DO NOTHING;