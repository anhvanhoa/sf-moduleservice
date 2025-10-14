package user_role

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrUserRoleNotFound      = oops.New("Không tìm thấy vai trò của người dùng")
	ErrUserRoleAlreadyExists = oops.New("Vai trò của người dùng đã tồn tại")
	ErrCreateUserRole        = oops.New("Lỗi khi tạo vai trò của người dùng")
	ErrListUserRoles         = oops.New("Lỗi khi lấy danh sách vai trò của người dùng")
	ErrUpdateUserRole        = oops.New("Lỗi khi cập nhật vai trò của người dùng")
	ErrDeleteUserRole        = oops.New("Lỗi khi xóa vai trò của người dùng")
	ErrDeleteByUserID        = oops.New("Lỗi khi xóa vai trò của người dùng theo ID người dùng")
	ErrDeleteByRoleID        = oops.New("Lỗi khi xóa vai trò của người dùng theo ID vai trò")
	ErrCountUserRoles        = oops.New("Lỗi khi đếm số lượng vai trò của người dùng")
	ErrCountByUserID         = oops.New("Lỗi khi đếm số lượng vai trò của người dùng theo ID người dùng")
	ErrCountByRoleID         = oops.New("Lỗi khi đếm số lượng vai trò của người dùng theo ID vai trò")
	ErrExistsUserRole        = oops.New("Lỗi khi kiểm tra sự tồn tại của vai trò của người dùng")
)
