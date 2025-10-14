-- Seed core roles
INSERT INTO roles (id, name, description, variant, status)
VALUES
    ('00000000-0000-0000-0000-000000000001', 'Admin', 'System administrator with full access', 'system', 'active'),
    ('00000000-0000-0000-0000-000000000002', 'User', 'Regular user with basic access', 'system', 'active')
ON CONFLICT DO NOTHING;

-- Seed core permissions (resource, action)
INSERT INTO permissions (id, resource, action, description)
VALUES
    ('10000000-0000-0000-0000-000000000001', 'role', 'create', 'Create roles'),
    ('10000000-0000-0000-0000-000000000002', 'role', 'read', 'Read roles'),
    ('10000000-0000-0000-0000-000000000003', 'role', 'update', 'Update roles'),
    ('10000000-0000-0000-0000-000000000004', 'role', 'delete', 'Delete roles'),
    ('10000000-0000-0000-0000-000000000005', 'permission', 'create', 'Create permissions'),
    ('10000000-0000-0000-0000-000000000006', 'permission', 'read', 'Read permissions'),
    ('10000000-0000-0000-0000-000000000007', 'permission', 'update', 'Update permissions'),
    ('10000000-0000-0000-0000-000000000008', 'permission', 'delete', 'Delete permissions'),
    ('10000000-0000-0000-0000-000000000009', 'user_role', 'assign', 'Assign roles to users'),
    ('10000000-0000-0000-0000-00000000000a', 'user_role', 'revoke', 'Revoke roles from users')
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


