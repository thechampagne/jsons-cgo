package main

/*
typedef struct
{
	int is_err;
	char* buffer;
} jsons_t;

typedef enum
{
	JSONS_FORMAT_AUTO,
	JSONS_FORMAT_JSON
} jsons_format_t;
*/
import "C"
import (
	"unsafe"
	jsons "github.com/qjebbs/go-jsons"
)

func jsons_formatToFormat(format C.jsons_format_t) jsons.Format {
	if format == C.JSONS_FORMAT_AUTO {
		return "auto"
	} else if format == C.JSONS_FORMAT_JSON {
		return "json"
	} else {
		return "auto"
	}
}

//export jsons_merge
func jsons_merge(inputs **C.char, inputs_len C.size_t) C.jsons_t {
	var self C.jsons_t
	inputs_go := (*[1<<30 - 1]*C.char)(unsafe.Pointer(inputs))[:inputs_len:inputs_len]
	res, err := jsons.Merge(inputs_go)
	if err != nil {
		self.buffer = C.CString(err.Error())
		self.is_err = 1
		return self
	}
	self.buffer = C.CString(string(res))
	self.is_err = 0
	return self
}

//export jsons_merge_as
func jsons_merge_as(format C.jsons_format_t, inputs **C.char, inputs_len C.size_t) C.jsons_t {
	var self C.jsons_t
	inputs_go := (*[1<<30 - 1]*C.char)(unsafe.Pointer(inputs))[:inputs_len:inputs_len]
	res, err := jsons.MergeAs(jsons_formatToFormat(format), inputs_go)
	if err != nil {
		self.buffer = C.CString(err.Error())
		self.is_err = 1
		return self
	}
	self.buffer = C.CString(string(res))
	self.is_err = 0
	return self
}

func main() {}
