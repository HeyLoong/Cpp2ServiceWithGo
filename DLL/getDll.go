package DLL

import (
	"syscall"
	"unsafe"
)

// 获取lib
var Lib *syscall.LazyDLL

func init() {
	Lib = syscall.NewLazyDLL("DLL/c_gcm.dll")
}

// Go调用C++时  常见变量类型需转为指针
func IntPtr(n int) uintptr {
	return uintptr(n)
}

func FloatPtr(f float64) uintptr {
	bits := *(*uint64)(unsafe.Pointer(&f))
	return uintptr(bits)
}

func StrPtr(s string) uintptr {
	ptr, _ := syscall.BytePtrFromString(s)
	return uintptr(unsafe.Pointer(ptr))
}

// 其他公有函数...
