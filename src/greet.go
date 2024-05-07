package main

//#cgo LDFLAGS: -L${SRCDIR}/../target/aarch64-apple-darwin/debug -lrust_greet
//#include <stdlib.h>
// char* greet(const char* name);
// void free_string(char* s);
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	name := C.CString("World")
	defer C.free(unsafe.Pointer(name))
	greeting := C.greet(name)
	defer C.free_string(greeting)
	fmt.Println(C.GoString(greeting))
}
