package glti

import "runtime"

// ProfilerEnv is the root object for registering profiling handlers
type ProfilerEnv struct {
}

// ---------------------------------------------------------------------------------------------------------------------
// pprof events (see https://golang.org/pkg/runtime/pprof/ --> requires multi-user support)

// RegisterPprofCpuProfileHandler enables CPU profiling and registers the given callback function to be called with
// each snapshot taken.
//
// CPU profiling determines where a program spends its time while actively consuming CPU cycles (as opposed to
// sleeping or waiting for I/O).
func (env *ProfilerEnv) RegisterPprofCpuProfileHandler(profileRate int, cb func([]StackRecord)) error {
	runtime.SetCPUProfileRate(profileRate)
	panic("not implemented")
}

// RegisterPprofHeapProfileHandler enables memory profiling and registers the given callback function to be called with
// each snapshot taken.
//
// Heap profiling reports memory allocation samples; used to monitor current and historical memory usage, and to check
// for memory leaks.
func (env *ProfilerEnv) RegisterPprofHeapProfileHandler(profileRate int, cb func([]MemProfileRecord)) error {
	runtime.MemProfileRate = profileRate
	panic("not implemented")
}

// RegisterPprofThreadCreationProfileHandler enables thread creation profiling and registers the given callback
// function to be called each time a new OS thread is created.
//
// Thread creation profiling reports the sections of the program that lead to the creation of new OS threads.
func (env *ProfilerEnv) RegisterPprofThreadCreationProfileHandler(profileRate int, cb func([]StackRecord)) error {
	panic("not implemented")
}

// RegisterPprofGoroutineProfileHandler enables goroutine profiling and registers the given callback function to be
// called with each snapshot taken.
//
// Goroutine profiling reports the stack traces of all current goroutines.
func (env *ProfilerEnv) RegisterPprofGoroutineProfileHandler(profileRate int, cb func([]StackRecord)) error {
	panic("not implemented")
}

// RegisterPprofBlockProfileHandler enables block profiling and registers the given callback function to be called
// with each snapshot taken.
//
// Block profiling shows where goroutines block waiting on synchronization primitives (including timer channels).
func (env *ProfilerEnv) RegisterPprofBlockProfileHandler(profileRate int, cb func([]BlockProfileRecord)) error {
	runtime.SetBlockProfileRate(profileRate)
	panic("not implemented")
}

// RegisterPprofMutexProfileHandler enables lock contentions profiling and registers the given callback function to be
// called with each snapshot taken.
//
// Mutex profile reports the lock contentions. When you think your CPU is not fully utilized due to a mutex contention,
// use this profile.
func (env *ProfilerEnv) RegisterPprofMutexProfileHandler(profileFraction int, cb func([]BlockProfileRecord)) error {
	runtime.SetMutexProfileFraction(profileFraction)
	panic("not implemented")
}
