// Package glti (Golang Tool Interface) is a programming interface for use by tools. It provides access to the type
// system of the Go runtime and an API which can be used for performance monitoring of the running Go application.
// GLTI supports tools that need access to Go runtime state, including but not limited to: monitoring, profiling and
// debugging.
package glti

import (
	"reflect"
	"runtime"
	"unsafe"
)

// GLTIEnv Golang Tool Interface root object
type GLTIEnv struct {
	GltiVersion string // offset known to monitoring plugins
	GoVersion   string // offset known to monitoring plugins

	Trace    TracingEnv
	Profiler ProfilerEnv
}

// Public API (called by agent code) ===================================================================================

// RegisterEvUnhandledPanic registers given callback function to be called before a Go application exits due to an
// unhandled panic.
func (env *GLTIEnv) RegisterEvUnhandledPanic(cb func(panicMsg string)) {
	panic("not implemented")
}

// ResolveSymbol returns address of Go functions and global variables
func (env *GLTIEnv) ResolveSymbol(symbolName string) uintptr {
	panic("not implemented")

	// Uses gosymtab and gopclntab --> no exposure of internal data structures
}

// ResolveIface returns type information of given interface (address of runtime.iface object)
func (env *GLTIEnv) ResolveIface(iface interface{}) reflect.Type {
	panic("not implemented")

	// Resolves Go interfaces using the Go runtime type system --> no exposure of internal data structures
}

// FuncForPC returns a *Func describing the function that contains the given program counter address, or else nil.
func (env *GLTIEnv) FuncForPC(fAddr uintptr) *runtime.Func {
	return runtime.FuncForPC(fAddr)
}

// FuncForName returns function type of Go function with given symbol name or nil if not found
func (env *GLTIEnv) FuncForName(symbolName string) *FuncType {
	panic("not implemented")

	// May perform internal lookup using the Go runtime type system --> no exposure of internal data structures
}

// GetG returns currently running goroutine (g struct pointer)
func (env *GLTIEnv) GetG() interface{} {
	panic("not implemented")

	// return runtime.getg()
}

// SetGls associates pointer gls with the goroutine gp. The pointer can be used to implement GLS (goroutine local
// storage).
func (env *GLTIEnv) SetGls(gp interface{}, gls unsafe.Pointer) {
	panic("not implemented")
}

// GetGls returns a goroutine's associated context pointer.
func (env *GLTIEnv) GetGls(gp interface{}) unsafe.Pointer {
	panic("not implemented")
}

// ReflectCall can be used to call arbitrary Go functions. It calls function fn with argsize bytes pointed at by
// arg (params + return values).
// The error return value is set if a recoverable panic occurred during execution of fn.
// go:nosplit
func (env *GLTIEnv) ReflectCall(fn, arg uintptr, argsize uint32) error {
	panic("not implemented")
}
