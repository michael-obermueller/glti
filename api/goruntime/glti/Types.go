package glti

import (
	"reflect"
	"runtime"
)

// FuncType is an extension of reflect.Type which additionally contains
// ABI version (when future Go versions support multiple ABIs)
type FuncType struct {
	reflect.Type
	ABI int32
}

// StackRecord is an extension of runtime.StackRecord which additionally
// contains the goroutine pointer
type StackRecord struct {
	runtime.StackRecord
	GP interface{}
}

// MemProfileRecord is an extension of runtime.MemProfileRecord which
// additionally contains the goroutine pointer
type MemProfileRecord struct {
	runtime.MemProfileRecord
	GP interface{}
}

// BlockProfileRecord is an extension of runtime.BlockProfileRecord which
// additionally contains the goroutine pointer
type BlockProfileRecord struct {
	runtime.BlockProfileRecord
	GP interface{}
}
