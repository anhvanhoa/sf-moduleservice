-- Seed core roles
INSERT INTO roles (id, name, description, variant, status)
VALUES
    ('00000000-0000-0000-0000-000000000001', 'Admin', 'System administrator with full access', 'system', 'active'),
    ('00000000-0000-0000-0000-000000000002', 'User', 'Regular user with basic access', 'system', 'active')
ON CONFLICT DO NOTHING;

-- Seed core permissions (resource, action, method)
INSERT INTO permissions (id, resource, action, method, description)
VALUES
    ('10000000-0000-0000-0000-000000000001', 'role', 'create', 'POST', 'Create roles'),
    ('10000000-0000-0000-0000-000000000002', 'role', 'read', 'GET', 'Read roles'),
    ('10000000-0000-0000-0000-000000000003', 'role', 'update', 'PUT', 'Update roles'),
    ('10000000-0000-0000-0000-000000000004', 'role', 'delete', 'DELETE', 'Delete roles'),
    ('10000000-0000-0000-0000-000000000005', 'permission', 'create', 'POST', 'Create permissions'),
    ('10000000-0000-0000-0000-000000000006', 'permission', 'read', 'GET', 'Read permissions'),
    ('10000000-0000-0000-0000-000000000007', 'permission', 'update', 'PUT', 'Update permissions'),
    ('10000000-0000-0000-0000-000000000008', 'permission', 'delete', 'DELETE', 'Delete permissions'),
    ('10000000-0000-0000-0000-000000000009', 'user_role', 'assign', 'POST', 'Assign roles to users'),
    ('10000000-0000-0000-0000-00000000000a', 'user_role', 'revoke', 'DELETE', 'Revoke roles from users')
ON CONFLICT DO NOTHING;

-- Map all permissions to Admin role
INSERT INTO role_permissions (role_id, permission_id)
SELECT '00000000-0000-0000-0000-000000000001' AS role_id, p.id AS permission_id
FROM permissions p
WHERE p.id IN (
    '10000000-0000-0000-0000-000000000001',
    '10000000-0000-0000-0000-000000000002',
    '10000000-0000-0000-0000-000000000003',
    '10000000-0000-0000-0000-000000000004',
    '10000000-0000-0000-0000-000000000005',
    '10000000-0000-0000-0000-000000000006',
    '10000000-0000-0000-0000-000000000007',
    '10000000-0000-0000-0000-000000000008',
    '10000000-0000-0000-0000-000000000009',
    '10000000-0000-0000-0000-00000000000a'
)
ON CONFLICT DO NOTHING;

-- Seed sample user roles (assigning Admin role to a sample user)
INSERT INTO user_roles (user_id, role_id)
VALUES
    ('20000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000001') -- Admin user
ON CONFLICT DO NOTHING;

-- Seed sample resource permissions
INSERT INTO resource_permissions (id, user_id, resource_type, resource_data, method, action)
VALUES
    ('30000000-0000-0000-0000-000000000001', '20000000-0000-0000-0000-000000000001', 'document', '{"id": "doc-001", "title": "Sample Document"}', 'GET', 'read'),
    ('30000000-0000-0000-0000-000000000002', '20000000-0000-0000-0000-000000000001', 'document', '{"id": "doc-001", "title": "Sample Document"}', 'PUT', 'update'),
    ('30000000-0000-0000-0000-000000000003', '20000000-0000-0000-0000-000000000001', 'document', '{"id": "doc-001", "title": "Sample Document"}', 'DELETE', 'delete')
ON CONFLICT DO NOTHING;


