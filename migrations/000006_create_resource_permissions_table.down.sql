-- Drop trigger
DROP TRIGGER IF EXISTS update_resource_permissions_updated_at ON resource_permissions;

-- Drop indexes
DROP INDEX IF EXISTS idx_resource_permissions_action;
DROP INDEX IF EXISTS idx_resource_permissions_resource_id;
DROP INDEX IF EXISTS idx_resource_permissions_resource_type;
DROP INDEX IF EXISTS idx_resource_permissions_user_id;

-- Drop table
DROP TABLE IF EXISTS resource_permissions;
