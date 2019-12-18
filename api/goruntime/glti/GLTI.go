// Package glti (Golang Tool Interface) implements an interface for use by tools. It provides access to the type
// system of the Go runtime and an API which can be used for performance monitoring of the running Go application.
// GLTI supports tools that need access to Go runtime state, including but not limited to monitoring, profiling, and
// debugging.
package glti

import (
	"reflect"
	"runtime"
)

// GLTIEnv is the Golang Tool Interface root object
type GLTIEnv struct {
	GltiVersion string // offset known to monitoring plugins
	GoVersion   string // offset known to monitoring plugins

	Trace    TracingEnv
	Profiler ProfilerEnv
}

// Public API (called by agent code) ===================================================================================

// RegisterEvUnhandledPanic registers the given callback function to be called before a Go application exits due to an
// unhandled panic.
func (env *GLTIEnv) RegisterEvUnhandledPanic(cb func(obj interface{})) {
	panic("not implemented")
}

// ResolveSymbol returns the address of the Go function or global variable with the given name.
func (env *GLTIEnv) ResolveSymbol(symbolName string) uintptr {
	panic("not implemented")

	// Uses gosymtab and gopclntab --> no exposure of internal data structures
}

// ResolveIface returns type information of the given interface.
func (env *GLTIEnv) ResolveIface(iface interface{}) reflect.Type {
	panic("not implemented")

	// Resolves Go interfaces using the Go runtime type system --> no exposure of internal data structures
}

// FuncForPC returns a *Func describing the function that contains the given program counter address, or else nil.
func (env *GLTIEnv) FuncForPC(fAddr uintptr) *runtime.Func {
	return runtime.FuncForPC(fAddr)
}

// FuncForName returns the function type of the Go function with the given name, or nil if not found.
func (env *GLTIEnv) FuncForName(symbolName string) *FuncType {
	panic("not implemented")

	// May perform internal lookup using the Go runtime type system --> no exposure of internal data structures
}

// GetG returns the currently running goroutine (g struct pointer).
func (env *GLTIEnv) GetG() interface{} {
	panic("not implemented")

	// return runtime.getg()
}

// SetGls associates the given object with the goroutine gp (essentially implementing GLS, goroutine local storage).
func (env *GLTIEnv) SetGls(gp interface{}, obj interface{}) {
	panic("not implemented")
}

// GetGls returns the given goroutine's associated context object.
func (env *GLTIEnv) GetGls(gp interface{}) interface{} {
	panic("not implemented")
}

// SetGlsPtr associates the given memory address with the goroutine gp (usually points to native memory provided
// by non-Go agents)
func (env *GLTIEnv) SetGlsPtr(gp interface{}, objPtr uintptr) {
	panic("not implemented")
}

// GetGlsPtr returns the given goroutine's associated memory address.
func (env *GLTIEnv) GetGlsPtr(gp interface{}) uintptr {
	panic("not implemented")
}

// ReflectCall can be used to call arbitrary Go functions. It calls function fn with argsize bytes pointed at by
// arg (params + return values).
// The error return value is set if a recoverable panic occurred during execution of fn.
// go:nosplit
func (env *GLTIEnv) ReflectCall(fn, arg uintptr, argsize uint32) error {
	panic("not implemented")
}
