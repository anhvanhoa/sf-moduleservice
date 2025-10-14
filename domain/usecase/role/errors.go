package role

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrRoleNotFound      = oops.New("Không tìm thấy vai trò")
	ErrRoleAlreadyExists = oops.New("Vai trò đã tồn tại")
	ErrCreateRole        = oops.New("Lỗi khi tạo vai trò")
	ErrListRoles         = oops.New("Lỗi khi lấy danh sách vai trò")
	ErrUpdateRole        = oops.New("Lỗi khi cập nhật vai trò")
	ErrDeleteRole        = oops.New("Lỗi khi xóa vai trò")
	ErrCheckRole         = oops.New("Lỗi khi kiểm tra sự tồn tại của vai trò")
)
