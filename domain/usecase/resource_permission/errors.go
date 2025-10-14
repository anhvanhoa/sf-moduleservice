package resource_permission

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrResourcePermissionNotFound      = oops.New("Không tìm thấy quyền của tài nguyên")
	ErrResourcePermissionAlreadyExists = oops.New("Quyền của tài nguyên đã tồn tại")
	ErrCreateResourcePermission        = oops.New("Lỗi khi tạo quyền của tài nguyên")
	ErrListResourcePermissions         = oops.New("Lỗi khi lấy danh sách quyền của tài nguyên")
	ErrDeleteResourcePermission        = oops.New("Lỗi khi xóa quyền của tài nguyên")
	ErrDeleteByUserID                  = oops.New("Lỗi khi xóa quyền của tài nguyên theo ID người dùng")
	ErrDeleteByResource                = oops.New("Lỗi khi xóa quyền của tài nguyên theo loại tài nguyên")
	ErrDeleteByUserAndResource         = oops.New("Lỗi khi xóa quyền của tài nguyên theo ID người dùng và loại tài nguyên")
	ErrCountResourcePermissions        = oops.New("Lỗi khi đếm số lượng quyền của tài nguyên")
	ErrCountByUserID                   = oops.New("Lỗi khi đếm số lượng quyền của tài nguyên theo ID người dùng")
	ErrCountByResource                 = oops.New("Lỗi khi đếm số lượng quyền của tài nguyên theo loại tài nguyên")
	ErrExistsResourcePermission        = oops.New("Lỗi khi kiểm tra sự tồn tại của quyền của tài nguyên")
	ErrUpdateResourcePermission        = oops.New("Lỗi khi cập nhật quyền của tài nguyên")
)
