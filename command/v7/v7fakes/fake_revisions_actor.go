// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/v9/actor/v7action"
	v7 "code.cloudfoundry.org/cli/v9/command/v7"
	"code.cloudfoundry.org/cli/v9/resources"
)

type FakeRevisionsActor struct {
	GetRevisionsByApplicationNameAndSpaceStub        func(string, string) ([]resources.Revision, v7action.Warnings, error)
	getRevisionsByApplicationNameAndSpaceMutex       sync.RWMutex
	getRevisionsByApplicationNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getRevisionsByApplicationNameAndSpaceReturns struct {
		result1 []resources.Revision
		result2 v7action.Warnings
		result3 error
	}
	getRevisionsByApplicationNameAndSpaceReturnsOnCall map[int]struct {
		result1 []resources.Revision
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRevisionsActor) GetRevisionsByApplicationNameAndSpace(arg1 string, arg2 string) ([]resources.Revision, v7action.Warnings, error) {
	fake.getRevisionsByApplicationNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getRevisionsByApplicationNameAndSpaceReturnsOnCall[len(fake.getRevisionsByApplicationNameAndSpaceArgsForCall)]
	fake.getRevisionsByApplicationNameAndSpaceArgsForCall = append(fake.getRevisionsByApplicationNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetRevisionsByApplicationNameAndSpaceStub
	fakeReturns := fake.getRevisionsByApplicationNameAndSpaceReturns
	fake.recordInvocation("GetRevisionsByApplicationNameAndSpace", []interface{}{arg1, arg2})
	fake.getRevisionsByApplicationNameAndSpaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeRevisionsActor) GetRevisionsByApplicationNameAndSpaceCallCount() int {
	fake.getRevisionsByApplicationNameAndSpaceMutex.RLock()
	defer fake.getRevisionsByApplicationNameAndSpaceMutex.RUnlock()
	return len(fake.getRevisionsByApplicationNameAndSpaceArgsForCall)
}

func (fake *FakeRevisionsActor) GetRevisionsByApplicationNameAndSpaceCalls(stub func(string, string) ([]resources.Revision, v7action.Warnings, error)) {
	fake.getRevisionsByApplicationNameAndSpaceMutex.Lock()
	defer fake.getRevisionsByApplicationNameAndSpaceMutex.Unlock()
	fake.GetRevisionsByApplicationNameAndSpaceStub = stub
}

func (fake *FakeRevisionsActor) GetRevisionsByApplicationNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getRevisionsByApplicationNameAndSpaceMutex.RLock()
	defer fake.getRevisionsByApplicationNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getRevisionsByApplicationNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRevisionsActor) GetRevisionsByApplicationNameAndSpaceReturns(result1 []resources.Revision, result2 v7action.Warnings, result3 error) {
	fake.getRevisionsByApplicationNameAndSpaceMutex.Lock()
	defer fake.getRevisionsByApplicationNameAndSpaceMutex.Unlock()
	fake.GetRevisionsByApplicationNameAndSpaceStub = nil
	fake.getRevisionsByApplicationNameAndSpaceReturns = struct {
		result1 []resources.Revision
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRevisionsActor) GetRevisionsByApplicationNameAndSpaceReturnsOnCall(i int, result1 []resources.Revision, result2 v7action.Warnings, result3 error) {
	fake.getRevisionsByApplicationNameAndSpaceMutex.Lock()
	defer fake.getRevisionsByApplicationNameAndSpaceMutex.Unlock()
	fake.GetRevisionsByApplicationNameAndSpaceStub = nil
	if fake.getRevisionsByApplicationNameAndSpaceReturnsOnCall == nil {
		fake.getRevisionsByApplicationNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 []resources.Revision
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getRevisionsByApplicationNameAndSpaceReturnsOnCall[i] = struct {
		result1 []resources.Revision
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRevisionsActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getRevisionsByApplicationNameAndSpaceMutex.RLock()
	defer fake.getRevisionsByApplicationNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRevisionsActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.RevisionsActor = new(FakeRevisionsActor)
