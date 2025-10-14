-- Drop trigger
DROP TRIGGER IF EXISTS update_permissions_updated_at ON permissions;

-- Drop indexes
DROP INDEX IF EXISTS idx_permissions_action;
DROP INDEX IF EXISTS idx_permissions_resource;

-- Drop table
DROP TABLE IF EXISTS permissions;
