package main

// #cgo CFLAGS: -I/opt/halon/include
// #cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-all
// #include <HalonMTA.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"

	"github.com/abadojack/whatlanggo"
)

func main() {}

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
	var text string
	var text_cs *C.char

	var args_0 = C.HalonMTA_hsl_argument_get(args, 0)
	if args_0 != nil {
		if !C.HalonMTA_hsl_value_get(args_0, C.HALONMTA_HSL_TYPE_STRING, unsafe.Pointer(&text_cs), nil) {
			exception := C.HalonMTA_hsl_throw(hhc)
			value_cs := C.CString("Invalid type of \"text\" argument")
			value_cs_up := unsafe.Pointer(value_cs)
			defer C.free(value_cs_up)
			C.HalonMTA_hsl_value_set(exception, C.HALONMTA_HSL_TYPE_EXCEPTION, value_cs_up, 0)
			return
		}
		text = C.GoString(text_cs)
	} else {
		exception := C.HalonMTA_hsl_throw(hhc)
		value_cs := C.CString("Missing required \"text\" argument")
		value_cs_up := unsafe.Pointer(value_cs)
		defer C.free(value_cs_up)
		C.HalonMTA_hsl_value_set(exception, C.HALONMTA_HSL_TYPE_EXCEPTION, value_cs_up, 0)
		return
	}

	info := whatlanggo.Detect(text)
	language := info.Lang.String()
	language_cs := C.CString(language)
	language_cs_up := unsafe.Pointer(language_cs)
	defer C.free(language_cs_up)

	C.HalonMTA_hsl_value_set(ret, C.HALONMTA_HSL_TYPE_STRING, language_cs_up, 0)
}

//export Halon_hsl_register
func Halon_hsl_register(hhrc *C.HalonHSLRegisterContext) C.bool {
	detect_language_cs := C.CString("detect_language")
	C.HalonMTA_hsl_register_function(hhrc, detect_language_cs, nil)
	C.HalonMTA_hsl_module_register_function(hhrc, detect_language_cs, nil)
	return true
}
