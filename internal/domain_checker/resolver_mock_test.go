package domain_checker

// Code generated by http://github.com/gojuno/minimock (3.0.6). DO NOT EDIT.

//go:generate minimock -i github.com/rekby/lets-proxy2/internal/domain_checker.Resolver -o ./resolver_mock_test.go

import (
	"context"
	"net"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ResolverMock implements Resolver
type ResolverMock struct {
	t minimock.Tester

	funcLookupIPAddr          func(ctx context.Context, host string) (ia1 []net.IPAddr, err error)
	inspectFuncLookupIPAddr   func(ctx context.Context, host string)
	afterLookupIPAddrCounter  uint64
	beforeLookupIPAddrCounter uint64
	LookupIPAddrMock          mResolverMockLookupIPAddr
}

// NewResolverMock returns a mock for Resolver
func NewResolverMock(t minimock.Tester) *ResolverMock {
	m := &ResolverMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.LookupIPAddrMock = mResolverMockLookupIPAddr{mock: m}
	m.LookupIPAddrMock.callArgs = []*ResolverMockLookupIPAddrParams{}

	return m
}

type mResolverMockLookupIPAddr struct {
	mock               *ResolverMock
	defaultExpectation *ResolverMockLookupIPAddrExpectation
	expectations       []*ResolverMockLookupIPAddrExpectation

	callArgs []*ResolverMockLookupIPAddrParams
	mutex    sync.RWMutex
}

// ResolverMockLookupIPAddrExpectation specifies expectation struct of the Resolver.LookupIPAddr
type ResolverMockLookupIPAddrExpectation struct {
	mock    *ResolverMock
	params  *ResolverMockLookupIPAddrParams
	results *ResolverMockLookupIPAddrResults
	Counter uint64
}

// ResolverMockLookupIPAddrParams contains parameters of the Resolver.LookupIPAddr
type ResolverMockLookupIPAddrParams struct {
	ctx  context.Context
	host string
}

// ResolverMockLookupIPAddrResults contains results of the Resolver.LookupIPAddr
type ResolverMockLookupIPAddrResults struct {
	ia1 []net.IPAddr
	err error
}

// Expect sets up expected params for Resolver.LookupIPAddr
func (mmLookupIPAddr *mResolverMockLookupIPAddr) Expect(ctx context.Context, host string) *mResolverMockLookupIPAddr {
	if mmLookupIPAddr.mock.funcLookupIPAddr != nil {
		mmLookupIPAddr.mock.t.Fatalf("ResolverMock.LookupIPAddr mock is already set by Set")
	}

	if mmLookupIPAddr.defaultExpectation == nil {
		mmLookupIPAddr.defaultExpectation = &ResolverMockLookupIPAddrExpectation{}
	}

	mmLookupIPAddr.defaultExpectation.params = &ResolverMockLookupIPAddrParams{ctx, host}
	for _, e := range mmLookupIPAddr.expectations {
		if minimock.Equal(e.params, mmLookupIPAddr.defaultExpectation.params) {
			mmLookupIPAddr.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLookupIPAddr.defaultExpectation.params)
		}
	}

	return mmLookupIPAddr
}

// Inspect accepts an inspector function that has same arguments as the Resolver.LookupIPAddr
func (mmLookupIPAddr *mResolverMockLookupIPAddr) Inspect(f func(ctx context.Context, host string)) *mResolverMockLookupIPAddr {
	if mmLookupIPAddr.mock.inspectFuncLookupIPAddr != nil {
		mmLookupIPAddr.mock.t.Fatalf("Inspect function is already set for ResolverMock.LookupIPAddr")
	}

	mmLookupIPAddr.mock.inspectFuncLookupIPAddr = f

	return mmLookupIPAddr
}

// Return sets up results that will be returned by Resolver.LookupIPAddr
func (mmLookupIPAddr *mResolverMockLookupIPAddr) Return(ia1 []net.IPAddr, err error) *ResolverMock {
	if mmLookupIPAddr.mock.funcLookupIPAddr != nil {
		mmLookupIPAddr.mock.t.Fatalf("ResolverMock.LookupIPAddr mock is already set by Set")
	}

	if mmLookupIPAddr.defaultExpectation == nil {
		mmLookupIPAddr.defaultExpectation = &ResolverMockLookupIPAddrExpectation{mock: mmLookupIPAddr.mock}
	}
	mmLookupIPAddr.defaultExpectation.results = &ResolverMockLookupIPAddrResults{ia1, err}
	return mmLookupIPAddr.mock
}

//Set uses given function f to mock the Resolver.LookupIPAddr method
func (mmLookupIPAddr *mResolverMockLookupIPAddr) Set(f func(ctx context.Context, host string) (ia1 []net.IPAddr, err error)) *ResolverMock {
	if mmLookupIPAddr.defaultExpectation != nil {
		mmLookupIPAddr.mock.t.Fatalf("Default expectation is already set for the Resolver.LookupIPAddr method")
	}

	if len(mmLookupIPAddr.expectations) > 0 {
		mmLookupIPAddr.mock.t.Fatalf("Some expectations are already set for the Resolver.LookupIPAddr method")
	}

	mmLookupIPAddr.mock.funcLookupIPAddr = f
	return mmLookupIPAddr.mock
}

// When sets expectation for the Resolver.LookupIPAddr which will trigger the result defined by the following
// Then helper
func (mmLookupIPAddr *mResolverMockLookupIPAddr) When(ctx context.Context, host string) *ResolverMockLookupIPAddrExpectation {
	if mmLookupIPAddr.mock.funcLookupIPAddr != nil {
		mmLookupIPAddr.mock.t.Fatalf("ResolverMock.LookupIPAddr mock is already set by Set")
	}

	expectation := &ResolverMockLookupIPAddrExpectation{
		mock:   mmLookupIPAddr.mock,
		params: &ResolverMockLookupIPAddrParams{ctx, host},
	}
	mmLookupIPAddr.expectations = append(mmLookupIPAddr.expectations, expectation)
	return expectation
}

// Then sets up Resolver.LookupIPAddr return parameters for the expectation previously defined by the When method
func (e *ResolverMockLookupIPAddrExpectation) Then(ia1 []net.IPAddr, err error) *ResolverMock {
	e.results = &ResolverMockLookupIPAddrResults{ia1, err}
	return e.mock
}

// LookupIPAddr implements Resolver
func (mmLookupIPAddr *ResolverMock) LookupIPAddr(ctx context.Context, host string) (ia1 []net.IPAddr, err error) {
	mm_atomic.AddUint64(&mmLookupIPAddr.beforeLookupIPAddrCounter, 1)
	defer mm_atomic.AddUint64(&mmLookupIPAddr.afterLookupIPAddrCounter, 1)

	if mmLookupIPAddr.inspectFuncLookupIPAddr != nil {
		mmLookupIPAddr.inspectFuncLookupIPAddr(ctx, host)
	}

	mm_params := &ResolverMockLookupIPAddrParams{ctx, host}

	// Record call args
	mmLookupIPAddr.LookupIPAddrMock.mutex.Lock()
	mmLookupIPAddr.LookupIPAddrMock.callArgs = append(mmLookupIPAddr.LookupIPAddrMock.callArgs, mm_params)
	mmLookupIPAddr.LookupIPAddrMock.mutex.Unlock()

	for _, e := range mmLookupIPAddr.LookupIPAddrMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ia1, e.results.err
		}
	}

	if mmLookupIPAddr.LookupIPAddrMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLookupIPAddr.LookupIPAddrMock.defaultExpectation.Counter, 1)
		mm_want := mmLookupIPAddr.LookupIPAddrMock.defaultExpectation.params
		mm_got := ResolverMockLookupIPAddrParams{ctx, host}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmLookupIPAddr.t.Errorf("ResolverMock.LookupIPAddr got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmLookupIPAddr.LookupIPAddrMock.defaultExpectation.results
		if mm_results == nil {
			mmLookupIPAddr.t.Fatal("No results are set for the ResolverMock.LookupIPAddr")
		}
		return (*mm_results).ia1, (*mm_results).err
	}
	if mmLookupIPAddr.funcLookupIPAddr != nil {
		return mmLookupIPAddr.funcLookupIPAddr(ctx, host)
	}
	mmLookupIPAddr.t.Fatalf("Unexpected call to ResolverMock.LookupIPAddr. %v %v", ctx, host)
	return
}

// LookupIPAddrAfterCounter returns a count of finished ResolverMock.LookupIPAddr invocations
func (mmLookupIPAddr *ResolverMock) LookupIPAddrAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLookupIPAddr.afterLookupIPAddrCounter)
}

// LookupIPAddrBeforeCounter returns a count of ResolverMock.LookupIPAddr invocations
func (mmLookupIPAddr *ResolverMock) LookupIPAddrBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLookupIPAddr.beforeLookupIPAddrCounter)
}

// Calls returns a list of arguments used in each call to ResolverMock.LookupIPAddr.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLookupIPAddr *mResolverMockLookupIPAddr) Calls() []*ResolverMockLookupIPAddrParams {
	mmLookupIPAddr.mutex.RLock()

	argCopy := make([]*ResolverMockLookupIPAddrParams, len(mmLookupIPAddr.callArgs))
	copy(argCopy, mmLookupIPAddr.callArgs)

	mmLookupIPAddr.mutex.RUnlock()

	return argCopy
}

// MinimockLookupIPAddrDone returns true if the count of the LookupIPAddr invocations corresponds
// the number of defined expectations
func (m *ResolverMock) MinimockLookupIPAddrDone() bool {
	for _, e := range m.LookupIPAddrMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LookupIPAddrMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLookupIPAddrCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLookupIPAddr != nil && mm_atomic.LoadUint64(&m.afterLookupIPAddrCounter) < 1 {
		return false
	}
	return true
}

// MinimockLookupIPAddrInspect logs each unmet expectation
func (m *ResolverMock) MinimockLookupIPAddrInspect() {
	for _, e := range m.LookupIPAddrMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ResolverMock.LookupIPAddr with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LookupIPAddrMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLookupIPAddrCounter) < 1 {
		if m.LookupIPAddrMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ResolverMock.LookupIPAddr")
		} else {
			m.t.Errorf("Expected call to ResolverMock.LookupIPAddr with params: %#v", *m.LookupIPAddrMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLookupIPAddr != nil && mm_atomic.LoadUint64(&m.afterLookupIPAddrCounter) < 1 {
		m.t.Error("Expected call to ResolverMock.LookupIPAddr")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ResolverMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockLookupIPAddrInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ResolverMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ResolverMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockLookupIPAddrDone()
}
