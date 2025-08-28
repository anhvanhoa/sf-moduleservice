package common

import "errors"

var (
	ErrModuleNotFound          = errors.New("Không tìm thấy module")
	ErrInvalidModuleName       = errors.New("Tên module không hợp lệ")
	ErrInvalidDescription      = errors.New("Mô tả module không hợp lệ")
	ErrInvalidStatus           = errors.New("Trạng thái module không hợp lệ")
	ErrInvalidModule           = errors.New("Module không hợp lệ")
	ErrModuleNameAlreadyExists = errors.New("Tên module đã tồn tại")
)

var (
	ErrModuleChildNotFound        = errors.New("Không tìm thấy module con")
	ErrInvalidModuleChildID       = errors.New("ID module con không hợp lệ")
	ErrInvalidModuleChildName     = errors.New("Tên module con không hợp lệ")
	ErrInvalidPath                = errors.New("Đường dẫn không hợp lệ")
	ErrInvalidMethod              = errors.New("Phương thức không hợp lệ")
	ErrInvalidModuleChild         = errors.New("Module con không hợp lệ")
	ErrPathAndMethodAlreadyExists = errors.New("Đường dẫn và phương thức đã tồn tại")
)
