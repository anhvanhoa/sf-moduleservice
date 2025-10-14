package entity

import (
	"time"
)

type UserRole struct {
	tableName struct{} `pg:"user_roles"`
	UserID    string
	RoleID    string
	CreatedAt time.Time
}

func (u *UserRole) TableName() any {
	return u.tableName
}

type UserRoleFilter struct {
	UserID string
	RoleID string
}
