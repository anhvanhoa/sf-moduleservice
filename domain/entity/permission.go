package entity

import (
	"time"
)

type Permission struct {
	tableName   struct{} `pg:"permissions"`
	ID          string
	Resource    string
	Action      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func (p *Permission) TableName() any {
	return p.tableName
}

type PermissionFilter struct {
	Resource string
	Action   string
}
