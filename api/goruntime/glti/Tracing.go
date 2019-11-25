package glti

// TracingEnv root object of trace event registration
type TracingEnv struct {
}

// ---------------------------------------------------------------------------------------------------------------------
// Trace events (see https://golang.org/pkg/internal/trace/ --> requires multi-user support):

// RegisterTraceEvProcStart registers given callback function to be called when a new thread starts running.
func (env *TracingEnv) RegisterTraceEvProcStart(func(ts int64, p interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvProcStop registers given callback function to be called when a thread stops.
func (env *TracingEnv) RegisterTraceEvProcStop(func(ts int64, p interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGCStart registers given callback function to be called when GC starts running.
func (env *TracingEnv) RegisterTraceEvGCStart(func(ts int64)) error {
	panic("not implemented")
}

// RegisterTraceEvGCDone registers given callback function to be called when GC is done.
func (env *TracingEnv) RegisterTraceEvGCDone(func(ts int64)) error {
	panic("not implemented")
}

// RegisterTraceEvGCSTWStart registers given callback function to be called when stop-the-world GC starts.
func (env *TracingEnv) RegisterTraceEvGCSTWStart(func(ts int64, gcKind int)) error {
	panic("not implemented")
}

// RegisterTraceEvGCSTWDone registers given callback function to be called when stop-the-world GC is done.
func (env *TracingEnv) RegisterTraceEvGCSTWDone(func(ts int64)) error {
	panic("not implemented")
}

// RegisterTraceEvGCSweepStart registers given callback function to be called when GC sweep phase starts.
func (env *TracingEnv) RegisterTraceEvGCSweepStart(func(ts int64)) error {
	panic("not implemented")
}

// RegisterTraceEvGCSweepDone registers given callback function to be called when GC sweep phase is done.
func (env *TracingEnv) RegisterTraceEvGCSweepDone(func(ts int64, bytesSwept, bytesReclaimed uintptr)) error {
	panic("not implemented")
}

// RegisterTraceEvGoCreate registers given callback function to be called when a new goroutine is created.
func (env *TracingEnv) RegisterTraceEvGoCreate(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoStart registers given callback function to be called when a goroutine starts running.
func (env *TracingEnv) RegisterTraceEvGoStart(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoEnd registers given callback function to be called when a goroutine ends
func (env *TracingEnv) RegisterTraceEvGoEnd(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoStop registers given callback function to be called when goroutine stops (like in select{})
func (env *TracingEnv) RegisterTraceEvGoStop(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoSched registers given callback function to be called when a goroutine calls Gosched
func (env *TracingEnv) RegisterTraceEvGoSched(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoPreempt registers given callback function to be called when a goroutine is preempted
func (env *TracingEnv) RegisterTraceEvGoPreempt(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoSleep registers given callback function to be called when a goroutine calls Sleep
func (env *TracingEnv) RegisterTraceEvGoSleep(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoBlock registers given callback function to be called when a goroutine blocks
func (env *TracingEnv) RegisterTraceEvGoBlock(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoUnblock registers given callback function to be called when a goroutine is unblocked
func (env *TracingEnv) RegisterTraceEvGoUnblock(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoBlockSend registers given callback function to be called when a goroutine blocks on chan send
func (env *TracingEnv) RegisterTraceEvGoBlockSend(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoBlockRecv registers given callback function to be called when a goroutine blocks on chan recv
func (env *TracingEnv) RegisterTraceEvGoBlockRecv(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoBlockSelect registers given callback function to be called when a goroutine blocks on select
func (env *TracingEnv) RegisterTraceEvGoBlockSelect(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoBlockSync registers given callback function to be called when a goroutine blocks on Mutex/RWMutex
func (env *TracingEnv) RegisterTraceEvGoBlockSync(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoBlockCond registers given callback function to be called when a goroutine blocks on Cond
func (env *TracingEnv) RegisterTraceEvGoBlockCond(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoBlockNet registers given callback function to be called when a goroutine blocks network
func (env *TracingEnv) RegisterTraceEvGoBlockNet(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoSysCall registers given callback function to be called when a syscall is entered
func (env *TracingEnv) RegisterTraceEvGoSysCall(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoSysExit registers given callback function to be called when a syscall exits
func (env *TracingEnv) RegisterTraceEvGoSysExit(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoSysBlock registers given callback function to be called when a syscall blocks
func (env *TracingEnv) RegisterTraceEvGoSysBlock(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoWaiting registers given callback function to be called for goroutines which are blocked when tracing starts
func (env *TracingEnv) RegisterTraceEvGoWaiting(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoInSyscall registers given callback function to be called for goroutines which are in syscall when tracing starts
func (env *TracingEnv) RegisterTraceEvGoInSyscall(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvHeapAlloc registers given callback function to be called when memstats.heap_live changes
func (env *TracingEnv) RegisterTraceEvHeapAlloc(func(ts int64, heapAlloc uint64)) error {
	panic("not implemented")
}

// RegisterTraceEvNextGC registers given callback function to be called when memstats.next_gc changes
func (env *TracingEnv) RegisterTraceEvNextGC(func(ts int64, nextGc uint64)) error {
	panic("not implemented")
}

// RegisterTraceEvTimerGoroutine registers given callback function to be called on new timer goroutines
func (env *TracingEnv) RegisterTraceEvTimerGoroutine(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvFutileWakeup registers given callback function to be called when the previous wakeup of this goroutine
// was futile
func (env *TracingEnv) RegisterTraceEvFutileWakeup(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvString registers given callback function to be called when a new string is added by the tracing framework
func (env *TracingEnv) RegisterTraceEvString(func(ts int64, val string)) error {
	panic("not implemented")
}

// RegisterTraceEvGoStartLocal registers given callback function to be called when a goroutine starts running on the same P
// as the last event
func (env *TracingEnv) RegisterTraceEvGoStartLocal(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoUnblockLocal registers given callback function to be called when a goroutine is unblocked on the same P
// as the last event
func (env *TracingEnv) RegisterTraceEvGoUnblockLocal(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGoSysExitLocal registers given callback function to be called when syscall exits on the same P as the
// last event
func (env *TracingEnv) RegisterTraceEvGoSysExitLocal(func(ts int64)) error {
	panic("not implemented")
}

// RegisterTraceEvGoStartLabel registers given callback function to be called when a goroutine starts running with label
func (env *TracingEnv) RegisterTraceEvGoStartLabel(func(ts int64, gp interface{}, label string)) error {
	panic("not implemented")
}

// RegisterTraceEvGoBlockGC registers given callback function to be called when goroutine blocks on GC assist
func (env *TracingEnv) RegisterTraceEvGoBlockGC(func(ts int64, gp interface{})) error {
	panic("not implemented")
}

// RegisterTraceEvGCMarkAssistStart registers given callback function to be called when GC mark assist starts
func (env *TracingEnv) RegisterTraceEvGCMarkAssistStart(func(ts int64)) error {
	panic("not implemented")
}

// RegisterTraceEvGCMarkAssistDone registers given callback function to be called when GC mark assist is done
func (env *TracingEnv) RegisterTraceEvGCMarkAssistDone(func(ts int64)) error {
	panic("not implemented")
}

// RegisterTraceEvUserTaskCreate registers given callback function to be called when a user task starts (trace.NewContext)
func (env *TracingEnv) RegisterTraceEvUserTaskCreate(func(ts int64, taskId, parentTaskId uint64, name string)) error {
	panic("not implemented")
}

// RegisterTraceEvUserTaskEnd registers given callback function to be called when a user task ends
func (env *TracingEnv) RegisterTraceEvUserTaskEnd(func(ts int64, taskId uint64)) error {
	panic("not implemented")
}

// RegisterTraceEvUserRegion registers given callback function to be called on user regions (trace.WithRegion)
func (env *TracingEnv) RegisterTraceEvUserRegion(func(ts int64, taskId uint64, mode uint8, name string)) error {
	panic("not implemented")
}

// RegisterTraceEvUserLog registers given callback function to be called on user logs (trace.Log)
func (env *TracingEnv) RegisterTraceEvUserLog(func(ts int64, taskId uint64, key, value string)) error {
	panic("not implemented")
}
