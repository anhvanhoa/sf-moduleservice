-- Create resource_permissions table
CREATE TABLE resource_permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    resource_type VARCHAR(255) NOT NULL,
    resource_id UUID NOT NULL,
    action VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create trigger for resource_permissions updated_at
CREATE TRIGGER update_resource_permissions_updated_at BEFORE
UPDATE ON resource_permissions FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Create indexes for better performance
CREATE INDEX idx_resource_permissions_user_id ON resource_permissions(user_id);
CREATE INDEX idx_resource_permissions_resource_type ON resource_permissions(resource_type);
CREATE INDEX idx_resource_permissions_resource_id ON resource_permissions(resource_id);
CREATE INDEX idx_resource_permissions_action ON resource_permissions(action);
