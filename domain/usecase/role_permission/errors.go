package role_permission

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrRolePermissionNotFound      = oops.New("Không tìm thấy quyền của vai trò")
	ErrRolePermissionAlreadyExists = oops.New("Quyền của vai trò đã tồn tại")
	ErrCreateRolePermission        = oops.New("Lỗi khi tạo quyền của vai trò")
	ErrListRolePermissions         = oops.New("Lỗi khi lấy danh sách quyền của vai trò")
	ErrDeleteRolePermission        = oops.New("Lỗi khi xóa quyền của vai trò")
	ErrDeleteByPermissionID        = oops.New("Lỗi khi xóa quyền của vai trò theo ID quyền")
	ErrCountRolePermissions        = oops.New("Lỗi khi đếm số lượng quyền của vai trò")
	ErrCountByRoleID               = oops.New("Lỗi khi đếm số lượng quyền của vai trò theo ID vai trò")
	ErrCountByPermissionID         = oops.New("Lỗi khi đếm số lượng quyền của vai trò theo ID quyền")
	ErrExistsRolePermission        = oops.New("Lỗi khi kiểm tra sự tồn tại của quyền của vai trò")
)
