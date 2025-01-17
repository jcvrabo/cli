// Code generated by counterfeiter. DO NOT EDIT.
package coreconfigfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/v9/cf/configuration/coreconfig"
)

type FakeEndpointRepository struct {
	GetCCInfoStub        func(string) (*coreconfig.CCInfo, string, error)
	getCCInfoMutex       sync.RWMutex
	getCCInfoArgsForCall []struct {
		arg1 string
	}
	getCCInfoReturns struct {
		result1 *coreconfig.CCInfo
		result2 string
		result3 error
	}
	getCCInfoReturnsOnCall map[int]struct {
		result1 *coreconfig.CCInfo
		result2 string
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEndpointRepository) GetCCInfo(arg1 string) (*coreconfig.CCInfo, string, error) {
	fake.getCCInfoMutex.Lock()
	ret, specificReturn := fake.getCCInfoReturnsOnCall[len(fake.getCCInfoArgsForCall)]
	fake.getCCInfoArgsForCall = append(fake.getCCInfoArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetCCInfoStub
	fakeReturns := fake.getCCInfoReturns
	fake.recordInvocation("GetCCInfo", []interface{}{arg1})
	fake.getCCInfoMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeEndpointRepository) GetCCInfoCallCount() int {
	fake.getCCInfoMutex.RLock()
	defer fake.getCCInfoMutex.RUnlock()
	return len(fake.getCCInfoArgsForCall)
}

func (fake *FakeEndpointRepository) GetCCInfoCalls(stub func(string) (*coreconfig.CCInfo, string, error)) {
	fake.getCCInfoMutex.Lock()
	defer fake.getCCInfoMutex.Unlock()
	fake.GetCCInfoStub = stub
}

func (fake *FakeEndpointRepository) GetCCInfoArgsForCall(i int) string {
	fake.getCCInfoMutex.RLock()
	defer fake.getCCInfoMutex.RUnlock()
	argsForCall := fake.getCCInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEndpointRepository) GetCCInfoReturns(result1 *coreconfig.CCInfo, result2 string, result3 error) {
	fake.getCCInfoMutex.Lock()
	defer fake.getCCInfoMutex.Unlock()
	fake.GetCCInfoStub = nil
	fake.getCCInfoReturns = struct {
		result1 *coreconfig.CCInfo
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeEndpointRepository) GetCCInfoReturnsOnCall(i int, result1 *coreconfig.CCInfo, result2 string, result3 error) {
	fake.getCCInfoMutex.Lock()
	defer fake.getCCInfoMutex.Unlock()
	fake.GetCCInfoStub = nil
	if fake.getCCInfoReturnsOnCall == nil {
		fake.getCCInfoReturnsOnCall = make(map[int]struct {
			result1 *coreconfig.CCInfo
			result2 string
			result3 error
		})
	}
	fake.getCCInfoReturnsOnCall[i] = struct {
		result1 *coreconfig.CCInfo
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeEndpointRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCCInfoMutex.RLock()
	defer fake.getCCInfoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEndpointRepository) recordInvocation(key string, args []interface{}) {
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

var _ coreconfig.EndpointRepository = new(FakeEndpointRepository)
