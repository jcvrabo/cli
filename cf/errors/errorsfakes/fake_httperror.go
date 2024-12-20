// Code generated by counterfeiter. DO NOT EDIT.
package errorsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/v9/cf/errors"
)

type FakeHTTPError struct {
	ErrorStub        func() string
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
	}
	errorReturns struct {
		result1 string
	}
	errorReturnsOnCall map[int]struct {
		result1 string
	}
	ErrorCodeStub        func() string
	errorCodeMutex       sync.RWMutex
	errorCodeArgsForCall []struct {
	}
	errorCodeReturns struct {
		result1 string
	}
	errorCodeReturnsOnCall map[int]struct {
		result1 string
	}
	StatusCodeStub        func() int
	statusCodeMutex       sync.RWMutex
	statusCodeArgsForCall []struct {
	}
	statusCodeReturns struct {
		result1 int
	}
	statusCodeReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHTTPError) Error() string {
	fake.errorMutex.Lock()
	ret, specificReturn := fake.errorReturnsOnCall[len(fake.errorArgsForCall)]
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
	}{})
	stub := fake.ErrorStub
	fakeReturns := fake.errorReturns
	fake.recordInvocation("Error", []interface{}{})
	fake.errorMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHTTPError) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeHTTPError) ErrorCalls(stub func() string) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = stub
}

func (fake *FakeHTTPError) ErrorReturns(result1 string) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = nil
	fake.errorReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeHTTPError) ErrorReturnsOnCall(i int, result1 string) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = nil
	if fake.errorReturnsOnCall == nil {
		fake.errorReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.errorReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeHTTPError) ErrorCode() string {
	fake.errorCodeMutex.Lock()
	ret, specificReturn := fake.errorCodeReturnsOnCall[len(fake.errorCodeArgsForCall)]
	fake.errorCodeArgsForCall = append(fake.errorCodeArgsForCall, struct {
	}{})
	stub := fake.ErrorCodeStub
	fakeReturns := fake.errorCodeReturns
	fake.recordInvocation("ErrorCode", []interface{}{})
	fake.errorCodeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHTTPError) ErrorCodeCallCount() int {
	fake.errorCodeMutex.RLock()
	defer fake.errorCodeMutex.RUnlock()
	return len(fake.errorCodeArgsForCall)
}

func (fake *FakeHTTPError) ErrorCodeCalls(stub func() string) {
	fake.errorCodeMutex.Lock()
	defer fake.errorCodeMutex.Unlock()
	fake.ErrorCodeStub = stub
}

func (fake *FakeHTTPError) ErrorCodeReturns(result1 string) {
	fake.errorCodeMutex.Lock()
	defer fake.errorCodeMutex.Unlock()
	fake.ErrorCodeStub = nil
	fake.errorCodeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeHTTPError) ErrorCodeReturnsOnCall(i int, result1 string) {
	fake.errorCodeMutex.Lock()
	defer fake.errorCodeMutex.Unlock()
	fake.ErrorCodeStub = nil
	if fake.errorCodeReturnsOnCall == nil {
		fake.errorCodeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.errorCodeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeHTTPError) StatusCode() int {
	fake.statusCodeMutex.Lock()
	ret, specificReturn := fake.statusCodeReturnsOnCall[len(fake.statusCodeArgsForCall)]
	fake.statusCodeArgsForCall = append(fake.statusCodeArgsForCall, struct {
	}{})
	stub := fake.StatusCodeStub
	fakeReturns := fake.statusCodeReturns
	fake.recordInvocation("StatusCode", []interface{}{})
	fake.statusCodeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHTTPError) StatusCodeCallCount() int {
	fake.statusCodeMutex.RLock()
	defer fake.statusCodeMutex.RUnlock()
	return len(fake.statusCodeArgsForCall)
}

func (fake *FakeHTTPError) StatusCodeCalls(stub func() int) {
	fake.statusCodeMutex.Lock()
	defer fake.statusCodeMutex.Unlock()
	fake.StatusCodeStub = stub
}

func (fake *FakeHTTPError) StatusCodeReturns(result1 int) {
	fake.statusCodeMutex.Lock()
	defer fake.statusCodeMutex.Unlock()
	fake.StatusCodeStub = nil
	fake.statusCodeReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeHTTPError) StatusCodeReturnsOnCall(i int, result1 int) {
	fake.statusCodeMutex.Lock()
	defer fake.statusCodeMutex.Unlock()
	fake.StatusCodeStub = nil
	if fake.statusCodeReturnsOnCall == nil {
		fake.statusCodeReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.statusCodeReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeHTTPError) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	fake.errorCodeMutex.RLock()
	defer fake.errorCodeMutex.RUnlock()
	fake.statusCodeMutex.RLock()
	defer fake.statusCodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHTTPError) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ errors.HTTPError = new(FakeHTTPError)
