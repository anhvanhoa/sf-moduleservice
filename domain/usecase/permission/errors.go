package permission

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrPermissionNotFound        = oops.New("Không tìm thấy quyền")
	ErrPermissionAlreadyExists   = oops.New("Quyền đã tồn tại")
	ErrCreatePermission          = oops.New("Lỗi khi tạo quyền")
	ErrListPermissions           = oops.New("Lỗi khi lấy danh sách quyền")
	ErrUpdatePermission          = oops.New("Lỗi khi cập nhật quyền")
	ErrDeletePermission          = oops.New("Lỗi khi xóa quyền")
	ErrDeleteByResourceAndAction = oops.New("Lỗi khi xóa quyền theo resource và action")
	ErrCountByResource           = oops.New("Lỗi khi đếm số lượng quyền theo resource")
	ErrCreateManyPermission      = oops.New("Lỗi khi tạo nhiều quyền")
)
