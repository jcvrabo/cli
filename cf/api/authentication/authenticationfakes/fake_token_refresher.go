// Code generated by counterfeiter. DO NOT EDIT.
package authenticationfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/v9/cf/api/authentication"
)

type FakeTokenRefresher struct {
	RefreshAuthTokenStub        func() (string, error)
	refreshAuthTokenMutex       sync.RWMutex
	refreshAuthTokenArgsForCall []struct {
	}
	refreshAuthTokenReturns struct {
		result1 string
		result2 error
	}
	refreshAuthTokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenRefresher) RefreshAuthToken() (string, error) {
	fake.refreshAuthTokenMutex.Lock()
	ret, specificReturn := fake.refreshAuthTokenReturnsOnCall[len(fake.refreshAuthTokenArgsForCall)]
	fake.refreshAuthTokenArgsForCall = append(fake.refreshAuthTokenArgsForCall, struct {
	}{})
	stub := fake.RefreshAuthTokenStub
	fakeReturns := fake.refreshAuthTokenReturns
	fake.recordInvocation("RefreshAuthToken", []interface{}{})
	fake.refreshAuthTokenMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTokenRefresher) RefreshAuthTokenCallCount() int {
	fake.refreshAuthTokenMutex.RLock()
	defer fake.refreshAuthTokenMutex.RUnlock()
	return len(fake.refreshAuthTokenArgsForCall)
}

func (fake *FakeTokenRefresher) RefreshAuthTokenCalls(stub func() (string, error)) {
	fake.refreshAuthTokenMutex.Lock()
	defer fake.refreshAuthTokenMutex.Unlock()
	fake.RefreshAuthTokenStub = stub
}

func (fake *FakeTokenRefresher) RefreshAuthTokenReturns(result1 string, result2 error) {
	fake.refreshAuthTokenMutex.Lock()
	defer fake.refreshAuthTokenMutex.Unlock()
	fake.RefreshAuthTokenStub = nil
	fake.refreshAuthTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenRefresher) RefreshAuthTokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.refreshAuthTokenMutex.Lock()
	defer fake.refreshAuthTokenMutex.Unlock()
	fake.RefreshAuthTokenStub = nil
	if fake.refreshAuthTokenReturnsOnCall == nil {
		fake.refreshAuthTokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.refreshAuthTokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenRefresher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.refreshAuthTokenMutex.RLock()
	defer fake.refreshAuthTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTokenRefresher) recordInvocation(key string, args []interface{}) {
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

var _ authentication.TokenRefresher = new(FakeTokenRefresher)
