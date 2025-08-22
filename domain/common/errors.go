package common

import (
	se "module-service/domain/service/error"
)

var (
	ErrModuleNotFound          = se.NewErr("Không tìm thấy module")
	ErrInvalidModuleName       = se.NewErr("Tên module không hợp lệ")
	ErrInvalidDescription      = se.NewErr("Mô tả module không hợp lệ")
	ErrInvalidStatus           = se.NewErr("Trạng thái module không hợp lệ")
	ErrInvalidModule           = se.NewErr("Module không hợp lệ")
	ErrModuleNameAlreadyExists = se.NewErr("Tên module đã tồn tại")
)

var (
	ErrModuleChildNotFound        = se.NewErr("Không tìm thấy module con")
	ErrInvalidModuleChildID       = se.NewErr("ID module con không hợp lệ")
	ErrInvalidModuleChildName     = se.NewErr("Tên module con không hợp lệ")
	ErrInvalidPath                = se.NewErr("Đường dẫn không hợp lệ")
	ErrInvalidMethod              = se.NewErr("Phương thức không hợp lệ")
	ErrInvalidModuleChild         = se.NewErr("Module con không hợp lệ")
	ErrPathAndMethodAlreadyExists = se.NewErr("Đường dẫn và phương thức đã tồn tại")
)
