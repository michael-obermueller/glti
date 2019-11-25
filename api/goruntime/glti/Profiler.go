package glti

import "runtime"

// ProfilerEnv root object to register profile handlers
type ProfilerEnv struct {
}

// ---------------------------------------------------------------------------------------------------------------------
// pprof events (see https://golang.org/pkg/runtime/pprof/ --> requires multi-user support)

// RegisterPprofCpuProfileHandler enables CPU profiling and registers callback function to be called on each snapshot
// taken.
//
// CPU profile determines where a program spends its time while actively consuming CPU cycles (as opposed to while
// sleeping or waiting for I/O).
func (env *ProfilerEnv) RegisterPprofCpuProfileHandler(profileRate int, cb func([]StackRecord)) error {
	runtime.SetCPUProfileRate(profileRate)
	panic("not implemented")
}

// RegisterPprofHeapProfileHandler enables memory profiling and registers callback function to be called on each snapshot
// taken.
//
// Heap profile reports memory allocation samples; used to monitor current and historical memory usage, and to check
// for memory leaks.
func (env *ProfilerEnv) RegisterPprofHeapProfileHandler(profileRate int, cb func([]MemProfileRecord)) error {
	runtime.MemProfileRate = profileRate
	panic("not implemented")
}

// RegisterPprofThreadCreationProfileHandler enables thread creation profiling and registers callback function to be
// called on each snapshot taken.
//
// Thread creation profile reports the sections of the program that lead the creation of new OS threads.
func (env *ProfilerEnv) RegisterPprofThreadCreationProfileHandler(profileRate int, cb func([]StackRecord)) error {
	panic("not implemented")
}

// RegisterPprofGoroutineProfileHandler enables goroutine profiling and registers callback function to be called on each
// snapshot taken.
//
// Goroutine profile reports the stack traces of all current goroutines.
func (env *ProfilerEnv) RegisterPprofGoroutineProfileHandler(profileRate int, cb func([]StackRecord)) error {
	panic("not implemented")
}

// RegisterPprofBlockProfileHandler enables block profiling and registers callback function to be called on each
// snapshot taken.
//
// Block profile shows where goroutines block waiting on synchronization primitives (including timer channels).
func (env *ProfilerEnv) RegisterPprofBlockProfileHandler(profileRate int, cb func([]BlockProfileRecord)) error {
	runtime.SetBlockProfileRate(profileRate)
	panic("not implemented")
}

// RegisterPprofMutexProfileHandler enables lock contentions profiling and registers callback function to be called on
// each snapshot taken.
//
// Mutex profile reports the lock contentions. When you think your CPU is not fully utilized due to a mutex contention,
// use this profile.
func (env *ProfilerEnv) RegisterPprofMutexProfileHandler(profileFraction int, cb func([]BlockProfileRecord)) error {
	runtime.SetMutexProfileFraction(profileFraction)
	panic("not implemented")
}
