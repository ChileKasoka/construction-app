INSERT INTO permissions (name, path, method)
VALUES 
  ('Create Permission', '/permissions', 'POST'),
  ('Get All Permissions', '/permissions', 'GET')
ON CONFLICT (path, method) DO NOTHING;


INSERT INTO role_permissions (role_id, permission_id)
SELECT 1, id FROM permissions 
WHERE (path = '/permissions' AND method = 'POST')
   OR (path = '/permissions' AND method = 'GET')
ON CONFLICT DO NOTHING;
