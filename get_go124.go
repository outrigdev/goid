//go:build go1.24 && !go1.25 && (amd64 || arm64)

// Copyright 2025, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package goid

import (
	"sync/atomic"
	"unsafe"
)

// guintptr holds a goroutine pointer, but typed as a uintptr to bypass write barriers
type guintptr uintptr

// muintptr is a *m that is not tracked by the garbage collector
type muintptr uintptr

// waitReason explains why a goroutine has been stopped
type waitReason uint8

// stack describes a Go execution stack
type stack struct {
	lo uintptr
	hi uintptr
}

// gobuf holds the scheduling context for a goroutine
type gobuf struct {
	sp   uintptr
	pc   uintptr
	g    guintptr
	ctxt unsafe.Pointer
	ret  uintptr
	lr   uintptr
	bp   uintptr // for framepointer-enabled architectures
}

// Forward declarations for types we don't need to fully define
type _panic struct{}
type _defer struct{}
type m struct{}
type ancestorInfo struct{}
type sudog struct{}
type timer struct{}
type goroutineProfileStateHolder struct{}
type coro struct{}
type synctestGroup struct{}
type gTraceState struct{}

// g struct definition for Go 1.24 - exact copy from runtime source
type g struct {
	// Stack parameters.
	// stack describes the actual stack memory: [stack.lo, stack.hi).
	// stackguard0 is the stack pointer compared in the Go stack growth prologue.
	// It is stack.lo+StackGuard normally, but can be StackPreempt to trigger a preemption.
	// stackguard1 is the stack pointer compared in the //go:systemstack stack growth prologue.
	// It is stack.lo+StackGuard on g0 and gsignal stacks.
	// It is ~0 on other goroutine stacks, to trigger a call to morestackc (and crash).
	stack       stack   // offset known to runtime/cgo
	stackguard0 uintptr // offset known to liblink
	stackguard1 uintptr // offset known to liblink

	_panic    *_panic // innermost panic - offset known to liblink
	_defer    *_defer // innermost defer
	m         *m      // current m; offset known to arm liblink
	sched     gobuf
	syscallsp uintptr // if status==Gsyscall, syscallsp = sched.sp to use during gc
	syscallpc uintptr // if status==Gsyscall, syscallpc = sched.pc to use during gc
	syscallbp uintptr // if status==Gsyscall, syscallbp = sched.bp to use in fpTraceback
	stktopsp  uintptr // expected sp at top of stack, to check in traceback
	// param is a generic pointer parameter field used to pass
	// values in particular contexts where other storage for the
	// parameter would be difficult to find. It is currently used
	// in four ways:
	// 1. When a channel operation wakes up a blocked goroutine, it sets param to
	//    point to the sudog of the completed blocking operation.
	// 2. By gcAssistAlloc1 to signal back to its caller that the goroutine completed
	//    the GC cycle. It is unsafe to do so in any other way, because the goroutine's
	//    stack may have moved in the meantime.
	// 3. By debugCallWrap to pass parameters to a new goroutine because allocating a
	//    closure in the runtime is forbidden.
	// 4. When a panic is recovered and control returns to the respective frame,
	//    param may point to a savedOpenDeferState.
	param        unsafe.Pointer
	atomicstatus atomic.Uint32
	stackLock    uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
	goid         uint64
	schedlink    guintptr
	waitsince    int64      // approx time when the g become blocked
	waitreason   waitReason // if status==Gwaiting

	preempt       bool // preemption signal, duplicates stackguard0 = stackpreempt
	preemptStop   bool // transition to _Gpreempted on preemption; otherwise, just deschedule
	preemptShrink bool // shrink stack at synchronous safe point

	// asyncSafePoint is set if g is stopped at an asynchronous
	// safe point. This means there are frames on the stack
	// without precise pointer information.
	asyncSafePoint bool

	paniconfault bool // panic (instead of crash) on unexpected fault address
	gcscandone   bool // g has scanned stack; protected by _Gscan bit in status
	throwsplit   bool // must not split stack
	// activeStackChans indicates that there are unlocked channels
	// pointing into this goroutine's stack. If true, stack
	// copying needs to acquire channel locks to protect these
	// areas of the stack.
	activeStackChans bool
	// parkingOnChan indicates that the goroutine is about to
	// park on a chansend or chanrecv. Used to signal an unsafe point
	// for stack shrinking.
	parkingOnChan atomic.Bool
	// inMarkAssist indicates whether the goroutine is in mark assist.
	// Used by the execution tracer.
	inMarkAssist bool
	coroexit     bool // argument to coroswitch_m

	raceignore    int8  // ignore race detection events
	nocgocallback bool  // whether disable callback from C
	tracking      bool  // whether we're tracking this G for sched latency statistics
	trackingSeq   uint8 // used to decide whether to track this G
	trackingStamp int64 // timestamp of when the G last started being tracked
	runnableTime  int64 // the amount of time spent runnable, cleared when running, only used when tracking
	lockedm       muintptr
	fipsIndicator uint8
	sig           uint32
	writebuf      []byte
	sigcode0      uintptr
	sigcode1      uintptr
	sigpc         uintptr
	parentGoid    uint64          // goid of goroutine that created this goroutine
	gopc          uintptr         // pc of go statement that created this goroutine
	ancestors     *[]ancestorInfo // ancestor information goroutine(s) that created this goroutine (only used if debug.tracebackancestors)
	startpc       uintptr         // pc of goroutine function
	racectx       uintptr
	waiting       *sudog         // sudog structures this g is waiting on (that have a valid elem ptr); in lock order
	cgoCtxt       []uintptr      // cgo traceback context
	labels        unsafe.Pointer // profiler labels
	timer         *timer         // cached timer for time.Sleep
	sleepWhen     int64          // when to sleep until
	selectDone    atomic.Uint32  // are we participating in a select and did someone win the race?

	// goroutineProfiled indicates the status of this goroutine's stack for the
	// current in-progress goroutine profile
	goroutineProfiled goroutineProfileStateHolder

	coroarg   *coro // argument during coroutine transfers
	syncGroup *synctestGroup

	// Per-G tracer state.
	trace gTraceState

	// Per-G GC state

	// gcAssistBytes is this G's GC assist credit in terms of
	// bytes allocated. If this is positive, then the G has credit
	// to allocate gcAssistBytes bytes without assisting. If this
	// is negative, then the G must correct this by performing
	// scan work. We track this in bytes to make it fast to update
	// and check for debt in the malloc hot path. The assist ratio
	// determines how this corresponds to scan work debt.
	gcAssistBytes int64
}

// getg returns the current goroutine pointer (implemented in assembly)
func getg() *g

// Get returns the current goroutine ID using the optimized assembly implementation for Go 1.24 on amd64/arm64.
// This implementation directly accesses the goroutine structure in memory for maximum performance.
func Get() uint64 {
	gptr := getg()
	if gptr == nil {
		return getFromStack()
	}
	return gptr.goid
}
