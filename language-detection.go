package main

// #cgo CFLAGS: -I/opt/halon/include
// #cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-all
// #include <HalonMTA.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/abadojack/whatlanggo"
)

func main() {}

func GetArgumentAsString(args *C.HalonHSLArguments, pos uint64) (string, error) {
	var x = C.HalonMTA_hsl_argument_get(args, C.ulong(pos))
	if x == nil {
		return "", fmt.Errorf("missing argument at position %d", pos)
	}
	var y *C.char
	if C.HalonMTA_hsl_value_get(x, C.HALONMTA_HSL_TYPE_STRING, unsafe.Pointer(&y), nil) {
		return C.GoString(y), nil
	} else {
		return "", fmt.Errorf("invalid argument at position %d", pos)
	}
}

func SetException(hhc *C.HalonHSLContext, msg string) {
	x := C.CString(msg)
	y := unsafe.Pointer(x)
	defer C.free(y)
	exception := C.HalonMTA_hsl_throw(hhc)
	C.HalonMTA_hsl_value_set(exception, C.HALONMTA_HSL_TYPE_EXCEPTION, y, 0)
}

func SetReturnValueToString(ret *C.HalonHSLValue, val string) {
	x := C.CString(val)
	y := unsafe.Pointer(x)
	defer C.free(y)
	C.HalonMTA_hsl_value_set(ret, C.HALONMTA_HSL_TYPE_STRING, y, 0)
}

//export Halon_version
func Halon_version() C.int {
	return C.HALONMTA_PLUGIN_VERSION
}

//export Halon_init
func Halon_init(hic *C.HalonInitContext) C.bool {
	return true
}

//export detect_language
func detect_language(hhc *C.HalonHSLContext, args *C.HalonHSLArguments, ret *C.HalonHSLValue) {
	text, err := GetArgumentAsString(args, 0)
	if err != nil {
		SetException(hhc, err.Error())
		return
	}
	info := whatlanggo.Detect(text)
	language := info.Lang.String()
	SetReturnValueToString(ret, language)
}

//export Halon_hsl_register
func Halon_hsl_register(hhrc *C.HalonHSLRegisterContext) C.bool {
	detect_language_cs := C.CString("detect_language")
	C.HalonMTA_hsl_register_function(hhrc, detect_language_cs, nil)
	C.HalonMTA_hsl_module_register_function(hhrc, detect_language_cs, nil)
	return true
}
