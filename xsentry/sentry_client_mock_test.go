package xsentry

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/galecore/xslog/xsentry.SentryClient -o ./sentry_client_mock_test.go -n SentryClientMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/getsentry/sentry-go"
	"github.com/gojuno/minimock/v3"
)

// SentryClientMock implements SentryClient
type SentryClientMock struct {
	t minimock.Tester

	funcCaptureEvent          func(event *sentry.Event, hint *sentry.EventHint, scope sentry.EventModifier) (ep1 *sentry.EventID)
	inspectFuncCaptureEvent   func(event *sentry.Event, hint *sentry.EventHint, scope sentry.EventModifier)
	afterCaptureEventCounter  uint64
	beforeCaptureEventCounter uint64
	CaptureEventMock          mSentryClientMockCaptureEvent
}

// NewSentryClientMock returns a mock for SentryClient
func NewSentryClientMock(t minimock.Tester) *SentryClientMock {
	m := &SentryClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CaptureEventMock = mSentryClientMockCaptureEvent{mock: m}
	m.CaptureEventMock.callArgs = []*SentryClientMockCaptureEventParams{}

	return m
}

type mSentryClientMockCaptureEvent struct {
	mock               *SentryClientMock
	defaultExpectation *SentryClientMockCaptureEventExpectation
	expectations       []*SentryClientMockCaptureEventExpectation

	callArgs []*SentryClientMockCaptureEventParams
	mutex    sync.RWMutex
}

// SentryClientMockCaptureEventExpectation specifies expectation struct of the SentryClient.CaptureEvent
type SentryClientMockCaptureEventExpectation struct {
	mock    *SentryClientMock
	params  *SentryClientMockCaptureEventParams
	results *SentryClientMockCaptureEventResults
	Counter uint64
}

// SentryClientMockCaptureEventParams contains parameters of the SentryClient.CaptureEvent
type SentryClientMockCaptureEventParams struct {
	event *sentry.Event
	hint  *sentry.EventHint
	scope sentry.EventModifier
}

// SentryClientMockCaptureEventResults contains results of the SentryClient.CaptureEvent
type SentryClientMockCaptureEventResults struct {
	ep1 *sentry.EventID
}

// Expect sets up expected params for SentryClient.CaptureEvent
func (mmCaptureEvent *mSentryClientMockCaptureEvent) Expect(event *sentry.Event, hint *sentry.EventHint, scope sentry.EventModifier) *mSentryClientMockCaptureEvent {
	if mmCaptureEvent.mock.funcCaptureEvent != nil {
		mmCaptureEvent.mock.t.Fatalf("SentryClientMock.CaptureEvent mock is already set by Set")
	}

	if mmCaptureEvent.defaultExpectation == nil {
		mmCaptureEvent.defaultExpectation = &SentryClientMockCaptureEventExpectation{}
	}

	mmCaptureEvent.defaultExpectation.params = &SentryClientMockCaptureEventParams{event, hint, scope}
	for _, e := range mmCaptureEvent.expectations {
		if minimock.Equal(e.params, mmCaptureEvent.defaultExpectation.params) {
			mmCaptureEvent.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCaptureEvent.defaultExpectation.params)
		}
	}

	return mmCaptureEvent
}

// Inspect accepts an inspector function that has same arguments as the SentryClient.CaptureEvent
func (mmCaptureEvent *mSentryClientMockCaptureEvent) Inspect(f func(event *sentry.Event, hint *sentry.EventHint, scope sentry.EventModifier)) *mSentryClientMockCaptureEvent {
	if mmCaptureEvent.mock.inspectFuncCaptureEvent != nil {
		mmCaptureEvent.mock.t.Fatalf("Inspect function is already set for SentryClientMock.CaptureEvent")
	}

	mmCaptureEvent.mock.inspectFuncCaptureEvent = f

	return mmCaptureEvent
}

// Return sets up results that will be returned by SentryClient.CaptureEvent
func (mmCaptureEvent *mSentryClientMockCaptureEvent) Return(ep1 *sentry.EventID) *SentryClientMock {
	if mmCaptureEvent.mock.funcCaptureEvent != nil {
		mmCaptureEvent.mock.t.Fatalf("SentryClientMock.CaptureEvent mock is already set by Set")
	}

	if mmCaptureEvent.defaultExpectation == nil {
		mmCaptureEvent.defaultExpectation = &SentryClientMockCaptureEventExpectation{mock: mmCaptureEvent.mock}
	}
	mmCaptureEvent.defaultExpectation.results = &SentryClientMockCaptureEventResults{ep1}
	return mmCaptureEvent.mock
}

// Set uses given function f to mock the SentryClient.CaptureEvent method
func (mmCaptureEvent *mSentryClientMockCaptureEvent) Set(f func(event *sentry.Event, hint *sentry.EventHint, scope sentry.EventModifier) (ep1 *sentry.EventID)) *SentryClientMock {
	if mmCaptureEvent.defaultExpectation != nil {
		mmCaptureEvent.mock.t.Fatalf("Default expectation is already set for the SentryClient.CaptureEvent method")
	}

	if len(mmCaptureEvent.expectations) > 0 {
		mmCaptureEvent.mock.t.Fatalf("Some expectations are already set for the SentryClient.CaptureEvent method")
	}

	mmCaptureEvent.mock.funcCaptureEvent = f
	return mmCaptureEvent.mock
}

// When sets expectation for the SentryClient.CaptureEvent which will trigger the result defined by the following
// Then helper
func (mmCaptureEvent *mSentryClientMockCaptureEvent) When(event *sentry.Event, hint *sentry.EventHint, scope sentry.EventModifier) *SentryClientMockCaptureEventExpectation {
	if mmCaptureEvent.mock.funcCaptureEvent != nil {
		mmCaptureEvent.mock.t.Fatalf("SentryClientMock.CaptureEvent mock is already set by Set")
	}

	expectation := &SentryClientMockCaptureEventExpectation{
		mock:   mmCaptureEvent.mock,
		params: &SentryClientMockCaptureEventParams{event, hint, scope},
	}
	mmCaptureEvent.expectations = append(mmCaptureEvent.expectations, expectation)
	return expectation
}

// Then sets up SentryClient.CaptureEvent return parameters for the expectation previously defined by the When method
func (e *SentryClientMockCaptureEventExpectation) Then(ep1 *sentry.EventID) *SentryClientMock {
	e.results = &SentryClientMockCaptureEventResults{ep1}
	return e.mock
}

// CaptureEvent implements SentryClient
func (mmCaptureEvent *SentryClientMock) CaptureEvent(event *sentry.Event, hint *sentry.EventHint, scope sentry.EventModifier) (ep1 *sentry.EventID) {
	mm_atomic.AddUint64(&mmCaptureEvent.beforeCaptureEventCounter, 1)
	defer mm_atomic.AddUint64(&mmCaptureEvent.afterCaptureEventCounter, 1)

	if mmCaptureEvent.inspectFuncCaptureEvent != nil {
		mmCaptureEvent.inspectFuncCaptureEvent(event, hint, scope)
	}

	mm_params := &SentryClientMockCaptureEventParams{event, hint, scope}

	// Record call args
	mmCaptureEvent.CaptureEventMock.mutex.Lock()
	mmCaptureEvent.CaptureEventMock.callArgs = append(mmCaptureEvent.CaptureEventMock.callArgs, mm_params)
	mmCaptureEvent.CaptureEventMock.mutex.Unlock()

	for _, e := range mmCaptureEvent.CaptureEventMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ep1
		}
	}

	if mmCaptureEvent.CaptureEventMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCaptureEvent.CaptureEventMock.defaultExpectation.Counter, 1)
		mm_want := mmCaptureEvent.CaptureEventMock.defaultExpectation.params
		mm_got := SentryClientMockCaptureEventParams{event, hint, scope}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCaptureEvent.t.Errorf("SentryClientMock.CaptureEvent got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCaptureEvent.CaptureEventMock.defaultExpectation.results
		if mm_results == nil {
			mmCaptureEvent.t.Fatal("No results are set for the SentryClientMock.CaptureEvent")
		}
		return (*mm_results).ep1
	}
	if mmCaptureEvent.funcCaptureEvent != nil {
		return mmCaptureEvent.funcCaptureEvent(event, hint, scope)
	}
	mmCaptureEvent.t.Fatalf("Unexpected call to SentryClientMock.CaptureEvent. %v %v %v", event, hint, scope)
	return
}

// CaptureEventAfterCounter returns a count of finished SentryClientMock.CaptureEvent invocations
func (mmCaptureEvent *SentryClientMock) CaptureEventAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCaptureEvent.afterCaptureEventCounter)
}

// CaptureEventBeforeCounter returns a count of SentryClientMock.CaptureEvent invocations
func (mmCaptureEvent *SentryClientMock) CaptureEventBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCaptureEvent.beforeCaptureEventCounter)
}

// Calls returns a list of arguments used in each call to SentryClientMock.CaptureEvent.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCaptureEvent *mSentryClientMockCaptureEvent) Calls() []*SentryClientMockCaptureEventParams {
	mmCaptureEvent.mutex.RLock()

	argCopy := make([]*SentryClientMockCaptureEventParams, len(mmCaptureEvent.callArgs))
	copy(argCopy, mmCaptureEvent.callArgs)

	mmCaptureEvent.mutex.RUnlock()

	return argCopy
}

// MinimockCaptureEventDone returns true if the count of the CaptureEvent invocations corresponds
// the number of defined expectations
func (m *SentryClientMock) MinimockCaptureEventDone() bool {
	for _, e := range m.CaptureEventMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CaptureEventMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCaptureEventCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCaptureEvent != nil && mm_atomic.LoadUint64(&m.afterCaptureEventCounter) < 1 {
		return false
	}
	return true
}

// MinimockCaptureEventInspect logs each unmet expectation
func (m *SentryClientMock) MinimockCaptureEventInspect() {
	for _, e := range m.CaptureEventMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SentryClientMock.CaptureEvent with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CaptureEventMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCaptureEventCounter) < 1 {
		if m.CaptureEventMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SentryClientMock.CaptureEvent")
		} else {
			m.t.Errorf("Expected call to SentryClientMock.CaptureEvent with params: %#v", *m.CaptureEventMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCaptureEvent != nil && mm_atomic.LoadUint64(&m.afterCaptureEventCounter) < 1 {
		m.t.Error("Expected call to SentryClientMock.CaptureEvent")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SentryClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCaptureEventInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SentryClientMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *SentryClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCaptureEventDone()
}
