# glti

import "."

Package glti (Golang Tool Interface) is a programming interface for use by
tools. It provides access to the type system of the Go runtime and an API which
can be used for performance monitoring of the running Go application. GLTI
supports tools that need access to Go runtime state, including but not limited
to: monitoring, profiling and debugging.

## Usage

#### func  InitGLTI

```go
func InitGLTI()
```
InitGLTI loads and initializes the agent shared library.

#### type GLTIEnv

```go
type GLTIEnv struct {
	GltiVersion string // offset known to monitoring plugins
	GoVersion   string // offset known to monitoring plugins

	Trace    TracingEnv
	Profiler ProfilerEnv
}
```

GLTIEnv is the Golang Tool Interface root object

#### func (*GLTIEnv) FuncForName

```go
func (env *GLTIEnv) FuncForName(symbolName string) *FuncType
```
FuncForName returns the function type of the Go function with the given name, or
nil if not found.

#### func (*GLTIEnv) FuncForPC

```go
func (env *GLTIEnv) FuncForPC(fAddr uintptr) *runtime.Func
```
FuncForPC returns a *Func describing the function that contains the given
program counter address, or else nil.

#### func (*GLTIEnv) GetG

```go
func (env *GLTIEnv) GetG() interface{}
```
GetG returns the currently running goroutine (g struct pointer).

#### func (*GLTIEnv) GetGls

```go
func (env *GLTIEnv) GetGls(gp interface{}) interface{}
```
GetGls returns the given goroutine's associated context object.

#### func (*GLTIEnv) GetGlsPtr

```go
func (env *GLTIEnv) GetGlsPtr(gp interface{}) uintptr
```
GetGlsPtr returns the given goroutine's associated memory address.

#### func (*GLTIEnv) ReflectCall

```go
func (env *GLTIEnv) ReflectCall(fn, arg uintptr, argsize uint32) error
```
ReflectCall can be used to call arbitrary Go functions. It calls function fn
with argsize bytes pointed at by arg (params + return values). The error return
value is set if a recoverable panic occurred during execution of fn. go:nosplit

#### func (*GLTIEnv) RegisterEvUnhandledPanic

```go
func (env *GLTIEnv) RegisterEvUnhandledPanic(cb func(obj interface{}))
```
RegisterEvUnhandledPanic registers the given callback function to be called
before a Go application exits due to an unhandled panic.

#### func (*GLTIEnv) ResolveIface

```go
func (env *GLTIEnv) ResolveIface(iface interface{}) reflect.Type
```
ResolveIface returns type information of the given interface.

#### func (*GLTIEnv) ResolveSymbol

```go
func (env *GLTIEnv) ResolveSymbol(symbolName string) uintptr
```
ResolveSymbol returns the address of the Go function or global variable with the
given name.

#### func (*GLTIEnv) SetGls

```go
func (env *GLTIEnv) SetGls(gp interface{}, obj interface{})
```
SetGls associates the given object with the goroutine gp (essentially
implementing GLS, goroutine local storage).

#### func (*GLTIEnv) SetGlsPtr

```go
func (env *GLTIEnv) SetGlsPtr(gp interface{}, objPtr uintptr)
```
SetGlsPtr associates the given memory address with the goroutine gp (usually
points to native memory provided by non-Go agents)

#### type ProfilerEnv

```go
type ProfilerEnv struct {
}
```

ProfilerEnv is the root object for registering profiling handlers

#### func (*ProfilerEnv) RegisterPprofBlockProfileHandler

```go
func (env *ProfilerEnv) RegisterPprofBlockProfileHandler(profileRate int, cb func([]BlockProfileRecord)) error
```
RegisterPprofBlockProfileHandler enables block profiling and registers the given
callback function to be called with each snapshot taken.

Block profiling shows where goroutines block waiting on synchronization
primitives (including timer channels).

#### func (*ProfilerEnv) RegisterPprofCpuProfileHandler

```go
func (env *ProfilerEnv) RegisterPprofCpuProfileHandler(profileRate int, cb func([]StackRecord)) error
```
RegisterPprofCpuProfileHandler enables CPU profiling and registers the given
callback function to be called with each snapshot taken.

CPU profiling determines where a program spends its time while actively
consuming CPU cycles (as opposed to sleeping or waiting for I/O).

#### func (*ProfilerEnv) RegisterPprofGoroutineProfileHandler

```go
func (env *ProfilerEnv) RegisterPprofGoroutineProfileHandler(profileRate int, cb func([]StackRecord)) error
```
RegisterPprofGoroutineProfileHandler enables goroutine profiling and registers
the given callback function to be called with each snapshot taken.

Goroutine profiling reports the stack traces of all current goroutines.

#### func (*ProfilerEnv) RegisterPprofHeapProfileHandler

```go
func (env *ProfilerEnv) RegisterPprofHeapProfileHandler(profileRate int, cb func([]MemProfileRecord)) error
```
RegisterPprofHeapProfileHandler enables memory profiling and registers the given
callback function to be called with each snapshot taken.

Heap profiling reports memory allocation samples; used to monitor current and
historical memory usage, and to check for memory leaks.

#### func (*ProfilerEnv) RegisterPprofMutexProfileHandler

```go
func (env *ProfilerEnv) RegisterPprofMutexProfileHandler(profileFraction int, cb func([]BlockProfileRecord)) error
```
RegisterPprofMutexProfileHandler enables lock contentions profiling and
registers the given callback function to be called with each snapshot taken.

Mutex profile reports the lock contentions. When you think your CPU is not fully
utilized due to a mutex contention, use this profile.

#### func (*ProfilerEnv) RegisterPprofThreadCreationProfileHandler

```go
func (env *ProfilerEnv) RegisterPprofThreadCreationProfileHandler(profileRate int, cb func([]StackRecord)) error
```
RegisterPprofThreadCreationProfileHandler enables thread creation profiling and
registers the given callback function to be called each time a new OS thread is
created.

Thread creation profiling reports the sections of the program that lead to the
creation of new OS threads.

#### type TracingEnv

```go
type TracingEnv struct {
}
```

TracingEnv is the root object for trace event registration

#### func (*TracingEnv) RegisterTraceEvFutileWakeup

```go
func (env *TracingEnv) RegisterTraceEvFutileWakeup(func(ts int64, gp interface{})) error
```
RegisterTraceEvFutileWakeup registers the given callback function to be called
when the previous wakeup of this goroutine was futile.

#### func (*TracingEnv) RegisterTraceEvGCDone

```go
func (env *TracingEnv) RegisterTraceEvGCDone(func(ts int64)) error
```
RegisterTraceEvGCDone registers the given callback function to be called when GC
is done.

#### func (*TracingEnv) RegisterTraceEvGCMarkAssistDone

```go
func (env *TracingEnv) RegisterTraceEvGCMarkAssistDone(func(ts int64)) error
```
RegisterTraceEvGCMarkAssistDone registers the given callback function to be
called when GC mark assist is done.

#### func (*TracingEnv) RegisterTraceEvGCMarkAssistStart

```go
func (env *TracingEnv) RegisterTraceEvGCMarkAssistStart(func(ts int64)) error
```
RegisterTraceEvGCMarkAssistStart registers the given callback function to be
called when GC mark assist starts.

#### func (*TracingEnv) RegisterTraceEvGCSTWDone

```go
func (env *TracingEnv) RegisterTraceEvGCSTWDone(func(ts int64)) error
```
RegisterTraceEvGCSTWDone registers the given callback function to be called when
stop-the-world GC is done.

#### func (*TracingEnv) RegisterTraceEvGCSTWStart

```go
func (env *TracingEnv) RegisterTraceEvGCSTWStart(func(ts int64, gcKind int)) error
```
RegisterTraceEvGCSTWStart registers the given callback function to be called
when stop-the-world GC starts.

#### func (*TracingEnv) RegisterTraceEvGCStart

```go
func (env *TracingEnv) RegisterTraceEvGCStart(func(ts int64)) error
```
RegisterTraceEvGCStart registers the given callback function to be called when
GC starts running.

#### func (*TracingEnv) RegisterTraceEvGCSweepDone

```go
func (env *TracingEnv) RegisterTraceEvGCSweepDone(func(ts int64, bytesSwept, bytesReclaimed uintptr)) error
```
RegisterTraceEvGCSweepDone registers the given callback function to be called
when GC sweep phase is done.

#### func (*TracingEnv) RegisterTraceEvGCSweepStart

```go
func (env *TracingEnv) RegisterTraceEvGCSweepStart(func(ts int64)) error
```
RegisterTraceEvGCSweepStart registers the given callback function to be called
when GC sweep phase starts.

#### func (*TracingEnv) RegisterTraceEvGoBlock

```go
func (env *TracingEnv) RegisterTraceEvGoBlock(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoBlock registers the given callback function to be called when a
goroutine blocks.

#### func (*TracingEnv) RegisterTraceEvGoBlockCond

```go
func (env *TracingEnv) RegisterTraceEvGoBlockCond(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoBlockCond registers the given callback function to be called
when a goroutine blocks on Cond.

#### func (*TracingEnv) RegisterTraceEvGoBlockGC

```go
func (env *TracingEnv) RegisterTraceEvGoBlockGC(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoBlockGC registers the given callback function to be called when
goroutine blocks on GC assist.

#### func (*TracingEnv) RegisterTraceEvGoBlockNet

```go
func (env *TracingEnv) RegisterTraceEvGoBlockNet(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoBlockNet registers the given callback function to be called
when a goroutine blocks on network.

#### func (*TracingEnv) RegisterTraceEvGoBlockRecv

```go
func (env *TracingEnv) RegisterTraceEvGoBlockRecv(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoBlockRecv registers the given callback function to be called
when a goroutine blocks on chan recv.

#### func (*TracingEnv) RegisterTraceEvGoBlockSelect

```go
func (env *TracingEnv) RegisterTraceEvGoBlockSelect(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoBlockSelect registers the given callback function to be called
when a goroutine blocks on select.

#### func (*TracingEnv) RegisterTraceEvGoBlockSend

```go
func (env *TracingEnv) RegisterTraceEvGoBlockSend(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoBlockSend registers the given callback function to be called
when a goroutine blocks on chan send.

#### func (*TracingEnv) RegisterTraceEvGoBlockSync

```go
func (env *TracingEnv) RegisterTraceEvGoBlockSync(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoBlockSync registers the given callback function to be called
when a goroutine blocks on Mutex/RWMutex.

#### func (*TracingEnv) RegisterTraceEvGoCreate

```go
func (env *TracingEnv) RegisterTraceEvGoCreate(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoCreate registers the given callback function to be called when
a new goroutine is created.

#### func (*TracingEnv) RegisterTraceEvGoEnd

```go
func (env *TracingEnv) RegisterTraceEvGoEnd(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoEnd registers the given callback function to be called when a
goroutine ends.

#### func (*TracingEnv) RegisterTraceEvGoInSyscall

```go
func (env *TracingEnv) RegisterTraceEvGoInSyscall(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoInSyscall registers the given callback function to be called
for goroutines which are in syscall when tracing starts.

#### func (*TracingEnv) RegisterTraceEvGoPreempt

```go
func (env *TracingEnv) RegisterTraceEvGoPreempt(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoPreempt registers the given callback function to be called when
a goroutine is preempted.

#### func (*TracingEnv) RegisterTraceEvGoSched

```go
func (env *TracingEnv) RegisterTraceEvGoSched(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoSched registers the given callback function to be called when a
goroutine calls Gosched.

#### func (*TracingEnv) RegisterTraceEvGoSleep

```go
func (env *TracingEnv) RegisterTraceEvGoSleep(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoSleep registers the given callback function to be called when a
goroutine calls Sleep.

#### func (*TracingEnv) RegisterTraceEvGoStart

```go
func (env *TracingEnv) RegisterTraceEvGoStart(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoStart registers the given callback function to be called when a
goroutine starts running.

#### func (*TracingEnv) RegisterTraceEvGoStartLabel

```go
func (env *TracingEnv) RegisterTraceEvGoStartLabel(func(ts int64, gp interface{}, label string)) error
```
RegisterTraceEvGoStartLabel registers the given callback function to be called
when a goroutine starts running with label.

#### func (*TracingEnv) RegisterTraceEvGoStartLocal

```go
func (env *TracingEnv) RegisterTraceEvGoStartLocal(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoStartLocal registers the given callback function to be called
when a goroutine starts running on the same P as the last event.

#### func (*TracingEnv) RegisterTraceEvGoStop

```go
func (env *TracingEnv) RegisterTraceEvGoStop(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoStop registers the given callback function to be called when
goroutine stops (like in select{}).

#### func (*TracingEnv) RegisterTraceEvGoSysBlock

```go
func (env *TracingEnv) RegisterTraceEvGoSysBlock(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoSysBlock registers the given callback function to be called
when a syscall blocks.

#### func (*TracingEnv) RegisterTraceEvGoSysCall

```go
func (env *TracingEnv) RegisterTraceEvGoSysCall(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoSysCall registers the given callback function to be called when
a syscall is entered.

#### func (*TracingEnv) RegisterTraceEvGoSysExit

```go
func (env *TracingEnv) RegisterTraceEvGoSysExit(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoSysExit registers the given callback function to be called when
a syscall exits.

#### func (*TracingEnv) RegisterTraceEvGoSysExitLocal

```go
func (env *TracingEnv) RegisterTraceEvGoSysExitLocal(func(ts int64)) error
```
RegisterTraceEvGoSysExitLocal registers the given callback function to be called
when syscall exits on the same P as the last event.

#### func (*TracingEnv) RegisterTraceEvGoUnblock

```go
func (env *TracingEnv) RegisterTraceEvGoUnblock(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoUnblock registers the given callback function to be called when
a goroutine is unblocked.

#### func (*TracingEnv) RegisterTraceEvGoUnblockLocal

```go
func (env *TracingEnv) RegisterTraceEvGoUnblockLocal(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoUnblockLocal registers the given callback function to be called
when a goroutine is unblocked on the same P as the last event.

#### func (*TracingEnv) RegisterTraceEvGoWaiting

```go
func (env *TracingEnv) RegisterTraceEvGoWaiting(func(ts int64, gp interface{})) error
```
RegisterTraceEvGoWaiting registers the given callback function to be called for
goroutines which are blocked when tracing starts.

#### func (*TracingEnv) RegisterTraceEvHeapAlloc

```go
func (env *TracingEnv) RegisterTraceEvHeapAlloc(func(ts int64, heapAlloc uint64)) error
```
RegisterTraceEvHeapAlloc registers the given callback function to be called when
memstats.heap_live changes.

#### func (*TracingEnv) RegisterTraceEvNextGC

```go
func (env *TracingEnv) RegisterTraceEvNextGC(func(ts int64, nextGc uint64)) error
```
RegisterTraceEvNextGC registers the given callback function to be called when
memstats.next_gc changes.

#### func (*TracingEnv) RegisterTraceEvProcStart

```go
func (env *TracingEnv) RegisterTraceEvProcStart(func(ts int64, p interface{})) error
```
RegisterTraceEvProcStart registers the given callback function to be called when
a new thread starts running.

#### func (*TracingEnv) RegisterTraceEvProcStop

```go
func (env *TracingEnv) RegisterTraceEvProcStop(func(ts int64, p interface{})) error
```
RegisterTraceEvProcStop registers the given callback function to be called when
a thread stops.

#### func (*TracingEnv) RegisterTraceEvString

```go
func (env *TracingEnv) RegisterTraceEvString(func(ts int64, val string)) error
```
RegisterTraceEvString registers the given callback function to be called when a
new string is added by the tracing framework.

#### func (*TracingEnv) RegisterTraceEvTimerGoroutine

```go
func (env *TracingEnv) RegisterTraceEvTimerGoroutine(func(ts int64, gp interface{})) error
```
RegisterTraceEvTimerGoroutine registers the given callback function to be called
on new timer goroutines.

#### func (*TracingEnv) RegisterTraceEvUserLog

```go
func (env *TracingEnv) RegisterTraceEvUserLog(func(ts int64, taskId uint64, key, value string)) error
```
RegisterTraceEvUserLog registers the given callback function to be called on
user logs (trace.Log).

#### func (*TracingEnv) RegisterTraceEvUserRegion

```go
func (env *TracingEnv) RegisterTraceEvUserRegion(func(ts int64, taskId uint64, mode uint8, name string)) error
```
RegisterTraceEvUserRegion registers the given callback function to be called on
user regions (trace.WithRegion).

#### func (*TracingEnv) RegisterTraceEvUserTaskCreate

```go
func (env *TracingEnv) RegisterTraceEvUserTaskCreate(func(ts int64, taskId, parentTaskId uint64, name string)) error
```
RegisterTraceEvUserTaskCreate registers the given callback function to be called
when a user task starts (trace.NewContext).

#### func (*TracingEnv) RegisterTraceEvUserTaskEnd

```go
func (env *TracingEnv) RegisterTraceEvUserTaskEnd(func(ts int64, taskId uint64)) error
```
RegisterTraceEvUserTaskEnd registers the given callback function to be called
when a user task ends.

#### type BlockProfileRecord

```go
type BlockProfileRecord struct {
	runtime.BlockProfileRecord
	GP interface{}
}
```

BlockProfileRecord is an extension of runtime.BlockProfileRecord which
additionally contains the goroutine pointer.

#### type FuncType

```go
type FuncType struct {
	reflect.Type
	ABI int32
}
```

FuncType is an extension of reflect.Type which additionally contains an ABI
version (when future Go versions support multiple ABIs).

#### type MemProfileRecord

```go
type MemProfileRecord struct {
	runtime.MemProfileRecord
	GP interface{}
}
```

MemProfileRecord is an extension of runtime.MemProfileRecord which additionally
contains the goroutine pointer.

#### type StackRecord

```go
type StackRecord struct {
	runtime.StackRecord
	GP interface{}
}
```

StackRecord is an extension of runtime.StackRecord which additionally contains
the goroutine pointer.