package phpbbencrypt

//#include <stdlib.h>
//#include "dencrypt.h"
import "C"

import (
	"unsafe"
)

func Encrypt(password string) (hash string) {
	arr := make([]C.char, 1024)
	cpassword := C.CString(password)
	defer C.free(unsafe.Pointer(cpassword))
	C.encrypt((&arr[0]), cpassword)
	return C.GoString((&arr[0]))
}

func Verify(password string, hash string) bool {
	cpassword := C.CString(password)
	defer C.free(unsafe.Pointer(cpassword))
	chash := C.CString(hash)
	defer C.free(unsafe.Pointer(chash))
	ret := C.verify(cpassword, chash)
	return 1 == ret
}
