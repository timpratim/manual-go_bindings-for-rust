package main

//#cgo LDFLAGS: -L${SRCDIR}/../target/aarch64-apple-darwin/debug -lrust_greet
//#include <stdlib.h>
// char* sendGreeting(const char* name);
// void free_string(char* s);
import "C"
import (
	"fmt"
	"unsafe"
)

func sendGreeting(recipient string) string {
	cstring := C.CString(recipient)
	defer C.free(unsafe.Pointer(cstring))
	greeting := C.sendGreeting(cstring)
	defer C.free_string(greeting)
	return C.GoString(greeting)
}

func main() {
	name := "Golangia"
	fmt.Println(sendGreeting(name))
}
