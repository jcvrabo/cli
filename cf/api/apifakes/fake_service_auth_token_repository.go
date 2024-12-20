// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"sync"

	"code.cloudfoundry.org/cli/v9/cf/api"
	"code.cloudfoundry.org/cli/v9/cf/models"
)

type FakeServiceAuthTokenRepository struct {
	CreateStub        func(models.ServiceAuthTokenFields) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 models.ServiceAuthTokenFields
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(models.ServiceAuthTokenFields) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 models.ServiceAuthTokenFields
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	FindAllStub        func() ([]models.ServiceAuthTokenFields, error)
	findAllMutex       sync.RWMutex
	findAllArgsForCall []struct {
	}
	findAllReturns struct {
		result1 []models.ServiceAuthTokenFields
		result2 error
	}
	findAllReturnsOnCall map[int]struct {
		result1 []models.ServiceAuthTokenFields
		result2 error
	}
	FindByLabelAndProviderStub        func(string, string) (models.ServiceAuthTokenFields, error)
	findByLabelAndProviderMutex       sync.RWMutex
	findByLabelAndProviderArgsForCall []struct {
		arg1 string
		arg2 string
	}
	findByLabelAndProviderReturns struct {
		result1 models.ServiceAuthTokenFields
		result2 error
	}
	findByLabelAndProviderReturnsOnCall map[int]struct {
		result1 models.ServiceAuthTokenFields
		result2 error
	}
	UpdateStub        func(models.ServiceAuthTokenFields) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 models.ServiceAuthTokenFields
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceAuthTokenRepository) Create(arg1 models.ServiceAuthTokenFields) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 models.ServiceAuthTokenFields
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceAuthTokenRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeServiceAuthTokenRepository) CreateCalls(stub func(models.ServiceAuthTokenFields) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeServiceAuthTokenRepository) CreateArgsForCall(i int) models.ServiceAuthTokenFields {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceAuthTokenRepository) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAuthTokenRepository) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAuthTokenRepository) Delete(arg1 models.ServiceAuthTokenFields) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 models.ServiceAuthTokenFields
	}{arg1})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceAuthTokenRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeServiceAuthTokenRepository) DeleteCalls(stub func(models.ServiceAuthTokenFields) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeServiceAuthTokenRepository) DeleteArgsForCall(i int) models.ServiceAuthTokenFields {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceAuthTokenRepository) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAuthTokenRepository) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAuthTokenRepository) FindAll() ([]models.ServiceAuthTokenFields, error) {
	fake.findAllMutex.Lock()
	ret, specificReturn := fake.findAllReturnsOnCall[len(fake.findAllArgsForCall)]
	fake.findAllArgsForCall = append(fake.findAllArgsForCall, struct {
	}{})
	stub := fake.FindAllStub
	fakeReturns := fake.findAllReturns
	fake.recordInvocation("FindAll", []interface{}{})
	fake.findAllMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceAuthTokenRepository) FindAllCallCount() int {
	fake.findAllMutex.RLock()
	defer fake.findAllMutex.RUnlock()
	return len(fake.findAllArgsForCall)
}

func (fake *FakeServiceAuthTokenRepository) FindAllCalls(stub func() ([]models.ServiceAuthTokenFields, error)) {
	fake.findAllMutex.Lock()
	defer fake.findAllMutex.Unlock()
	fake.FindAllStub = stub
}

func (fake *FakeServiceAuthTokenRepository) FindAllReturns(result1 []models.ServiceAuthTokenFields, result2 error) {
	fake.findAllMutex.Lock()
	defer fake.findAllMutex.Unlock()
	fake.FindAllStub = nil
	fake.findAllReturns = struct {
		result1 []models.ServiceAuthTokenFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAuthTokenRepository) FindAllReturnsOnCall(i int, result1 []models.ServiceAuthTokenFields, result2 error) {
	fake.findAllMutex.Lock()
	defer fake.findAllMutex.Unlock()
	fake.FindAllStub = nil
	if fake.findAllReturnsOnCall == nil {
		fake.findAllReturnsOnCall = make(map[int]struct {
			result1 []models.ServiceAuthTokenFields
			result2 error
		})
	}
	fake.findAllReturnsOnCall[i] = struct {
		result1 []models.ServiceAuthTokenFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAuthTokenRepository) FindByLabelAndProvider(arg1 string, arg2 string) (models.ServiceAuthTokenFields, error) {
	fake.findByLabelAndProviderMutex.Lock()
	ret, specificReturn := fake.findByLabelAndProviderReturnsOnCall[len(fake.findByLabelAndProviderArgsForCall)]
	fake.findByLabelAndProviderArgsForCall = append(fake.findByLabelAndProviderArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.FindByLabelAndProviderStub
	fakeReturns := fake.findByLabelAndProviderReturns
	fake.recordInvocation("FindByLabelAndProvider", []interface{}{arg1, arg2})
	fake.findByLabelAndProviderMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceAuthTokenRepository) FindByLabelAndProviderCallCount() int {
	fake.findByLabelAndProviderMutex.RLock()
	defer fake.findByLabelAndProviderMutex.RUnlock()
	return len(fake.findByLabelAndProviderArgsForCall)
}

func (fake *FakeServiceAuthTokenRepository) FindByLabelAndProviderCalls(stub func(string, string) (models.ServiceAuthTokenFields, error)) {
	fake.findByLabelAndProviderMutex.Lock()
	defer fake.findByLabelAndProviderMutex.Unlock()
	fake.FindByLabelAndProviderStub = stub
}

func (fake *FakeServiceAuthTokenRepository) FindByLabelAndProviderArgsForCall(i int) (string, string) {
	fake.findByLabelAndProviderMutex.RLock()
	defer fake.findByLabelAndProviderMutex.RUnlock()
	argsForCall := fake.findByLabelAndProviderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServiceAuthTokenRepository) FindByLabelAndProviderReturns(result1 models.ServiceAuthTokenFields, result2 error) {
	fake.findByLabelAndProviderMutex.Lock()
	defer fake.findByLabelAndProviderMutex.Unlock()
	fake.FindByLabelAndProviderStub = nil
	fake.findByLabelAndProviderReturns = struct {
		result1 models.ServiceAuthTokenFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAuthTokenRepository) FindByLabelAndProviderReturnsOnCall(i int, result1 models.ServiceAuthTokenFields, result2 error) {
	fake.findByLabelAndProviderMutex.Lock()
	defer fake.findByLabelAndProviderMutex.Unlock()
	fake.FindByLabelAndProviderStub = nil
	if fake.findByLabelAndProviderReturnsOnCall == nil {
		fake.findByLabelAndProviderReturnsOnCall = make(map[int]struct {
			result1 models.ServiceAuthTokenFields
			result2 error
		})
	}
	fake.findByLabelAndProviderReturnsOnCall[i] = struct {
		result1 models.ServiceAuthTokenFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAuthTokenRepository) Update(arg1 models.ServiceAuthTokenFields) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 models.ServiceAuthTokenFields
	}{arg1})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceAuthTokenRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeServiceAuthTokenRepository) UpdateCalls(stub func(models.ServiceAuthTokenFields) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeServiceAuthTokenRepository) UpdateArgsForCall(i int) models.ServiceAuthTokenFields {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceAuthTokenRepository) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAuthTokenRepository) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAuthTokenRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.findAllMutex.RLock()
	defer fake.findAllMutex.RUnlock()
	fake.findByLabelAndProviderMutex.RLock()
	defer fake.findByLabelAndProviderMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceAuthTokenRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.ServiceAuthTokenRepository = new(FakeServiceAuthTokenRepository)
