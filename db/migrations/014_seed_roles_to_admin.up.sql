INSERT INTO permissions (name, path, method) VALUES
  ('Get All Roles', '/roles', 'GET'),
  ('Create Role', '/roles', 'POST'),
  ('Get Role By ID', '/roles/{id}', 'GET'),
  ('Update Role', '/roles/{id}', 'PUT'),
  ('Delete Role', '/roles/{id}', 'DELETE'),
  ('Find Role By Name', '/roles/name/{name}', 'GET')
ON CONFLICT (path, method) DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT 1, id FROM permissions
WHERE (path = '/roles' AND method = 'GET')
   OR (path = '/roles' AND method = 'POST')
   OR (path = '/roles/{id}' AND method = 'GET')
   OR (path = '/roles/{id}' AND method = 'PUT')
   OR (path = '/roles/{id}' AND method = 'DELETE')
   OR (path = '/roles/name/{name}' AND method = 'GET')
ON CONFLICT DO NOTHING;
