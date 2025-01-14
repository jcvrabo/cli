// Code generated by counterfeiter. DO NOT EDIT.
package organizationsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/v9/cf/api/organizations"
	"code.cloudfoundry.org/cli/v9/cf/models"
)

type FakeOrganizationRepository struct {
	CreateStub        func(models.Organization) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 models.Organization
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	FindByNameStub        func(string) (models.Organization, error)
	findByNameMutex       sync.RWMutex
	findByNameArgsForCall []struct {
		arg1 string
	}
	findByNameReturns struct {
		result1 models.Organization
		result2 error
	}
	findByNameReturnsOnCall map[int]struct {
		result1 models.Organization
		result2 error
	}
	GetManyOrgsByGUIDStub        func([]string) ([]models.Organization, error)
	getManyOrgsByGUIDMutex       sync.RWMutex
	getManyOrgsByGUIDArgsForCall []struct {
		arg1 []string
	}
	getManyOrgsByGUIDReturns struct {
		result1 []models.Organization
		result2 error
	}
	getManyOrgsByGUIDReturnsOnCall map[int]struct {
		result1 []models.Organization
		result2 error
	}
	ListOrgsStub        func(int) ([]models.Organization, error)
	listOrgsMutex       sync.RWMutex
	listOrgsArgsForCall []struct {
		arg1 int
	}
	listOrgsReturns struct {
		result1 []models.Organization
		result2 error
	}
	listOrgsReturnsOnCall map[int]struct {
		result1 []models.Organization
		result2 error
	}
	RenameStub        func(string, string) error
	renameMutex       sync.RWMutex
	renameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	renameReturns struct {
		result1 error
	}
	renameReturnsOnCall map[int]struct {
		result1 error
	}
	SharePrivateDomainStub        func(string, string) error
	sharePrivateDomainMutex       sync.RWMutex
	sharePrivateDomainArgsForCall []struct {
		arg1 string
		arg2 string
	}
	sharePrivateDomainReturns struct {
		result1 error
	}
	sharePrivateDomainReturnsOnCall map[int]struct {
		result1 error
	}
	UnsharePrivateDomainStub        func(string, string) error
	unsharePrivateDomainMutex       sync.RWMutex
	unsharePrivateDomainArgsForCall []struct {
		arg1 string
		arg2 string
	}
	unsharePrivateDomainReturns struct {
		result1 error
	}
	unsharePrivateDomainReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrganizationRepository) Create(arg1 models.Organization) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 models.Organization
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

func (fake *FakeOrganizationRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeOrganizationRepository) CreateCalls(stub func(models.Organization) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeOrganizationRepository) CreateArgsForCall(i int) models.Organization {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOrganizationRepository) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) CreateReturnsOnCall(i int, result1 error) {
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

func (fake *FakeOrganizationRepository) Delete(arg1 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
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

func (fake *FakeOrganizationRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeOrganizationRepository) DeleteCalls(stub func(string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeOrganizationRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOrganizationRepository) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *FakeOrganizationRepository) FindByName(arg1 string) (models.Organization, error) {
	fake.findByNameMutex.Lock()
	ret, specificReturn := fake.findByNameReturnsOnCall[len(fake.findByNameArgsForCall)]
	fake.findByNameArgsForCall = append(fake.findByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.FindByNameStub
	fakeReturns := fake.findByNameReturns
	fake.recordInvocation("FindByName", []interface{}{arg1})
	fake.findByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOrganizationRepository) FindByNameCallCount() int {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return len(fake.findByNameArgsForCall)
}

func (fake *FakeOrganizationRepository) FindByNameCalls(stub func(string) (models.Organization, error)) {
	fake.findByNameMutex.Lock()
	defer fake.findByNameMutex.Unlock()
	fake.FindByNameStub = stub
}

func (fake *FakeOrganizationRepository) FindByNameArgsForCall(i int) string {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	argsForCall := fake.findByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOrganizationRepository) FindByNameReturns(result1 models.Organization, result2 error) {
	fake.findByNameMutex.Lock()
	defer fake.findByNameMutex.Unlock()
	fake.FindByNameStub = nil
	fake.findByNameReturns = struct {
		result1 models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationRepository) FindByNameReturnsOnCall(i int, result1 models.Organization, result2 error) {
	fake.findByNameMutex.Lock()
	defer fake.findByNameMutex.Unlock()
	fake.FindByNameStub = nil
	if fake.findByNameReturnsOnCall == nil {
		fake.findByNameReturnsOnCall = make(map[int]struct {
			result1 models.Organization
			result2 error
		})
	}
	fake.findByNameReturnsOnCall[i] = struct {
		result1 models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationRepository) GetManyOrgsByGUID(arg1 []string) ([]models.Organization, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getManyOrgsByGUIDMutex.Lock()
	ret, specificReturn := fake.getManyOrgsByGUIDReturnsOnCall[len(fake.getManyOrgsByGUIDArgsForCall)]
	fake.getManyOrgsByGUIDArgsForCall = append(fake.getManyOrgsByGUIDArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.GetManyOrgsByGUIDStub
	fakeReturns := fake.getManyOrgsByGUIDReturns
	fake.recordInvocation("GetManyOrgsByGUID", []interface{}{arg1Copy})
	fake.getManyOrgsByGUIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOrganizationRepository) GetManyOrgsByGUIDCallCount() int {
	fake.getManyOrgsByGUIDMutex.RLock()
	defer fake.getManyOrgsByGUIDMutex.RUnlock()
	return len(fake.getManyOrgsByGUIDArgsForCall)
}

func (fake *FakeOrganizationRepository) GetManyOrgsByGUIDCalls(stub func([]string) ([]models.Organization, error)) {
	fake.getManyOrgsByGUIDMutex.Lock()
	defer fake.getManyOrgsByGUIDMutex.Unlock()
	fake.GetManyOrgsByGUIDStub = stub
}

func (fake *FakeOrganizationRepository) GetManyOrgsByGUIDArgsForCall(i int) []string {
	fake.getManyOrgsByGUIDMutex.RLock()
	defer fake.getManyOrgsByGUIDMutex.RUnlock()
	argsForCall := fake.getManyOrgsByGUIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOrganizationRepository) GetManyOrgsByGUIDReturns(result1 []models.Organization, result2 error) {
	fake.getManyOrgsByGUIDMutex.Lock()
	defer fake.getManyOrgsByGUIDMutex.Unlock()
	fake.GetManyOrgsByGUIDStub = nil
	fake.getManyOrgsByGUIDReturns = struct {
		result1 []models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationRepository) GetManyOrgsByGUIDReturnsOnCall(i int, result1 []models.Organization, result2 error) {
	fake.getManyOrgsByGUIDMutex.Lock()
	defer fake.getManyOrgsByGUIDMutex.Unlock()
	fake.GetManyOrgsByGUIDStub = nil
	if fake.getManyOrgsByGUIDReturnsOnCall == nil {
		fake.getManyOrgsByGUIDReturnsOnCall = make(map[int]struct {
			result1 []models.Organization
			result2 error
		})
	}
	fake.getManyOrgsByGUIDReturnsOnCall[i] = struct {
		result1 []models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationRepository) ListOrgs(arg1 int) ([]models.Organization, error) {
	fake.listOrgsMutex.Lock()
	ret, specificReturn := fake.listOrgsReturnsOnCall[len(fake.listOrgsArgsForCall)]
	fake.listOrgsArgsForCall = append(fake.listOrgsArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.ListOrgsStub
	fakeReturns := fake.listOrgsReturns
	fake.recordInvocation("ListOrgs", []interface{}{arg1})
	fake.listOrgsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOrganizationRepository) ListOrgsCallCount() int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	return len(fake.listOrgsArgsForCall)
}

func (fake *FakeOrganizationRepository) ListOrgsCalls(stub func(int) ([]models.Organization, error)) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = stub
}

func (fake *FakeOrganizationRepository) ListOrgsArgsForCall(i int) int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	argsForCall := fake.listOrgsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOrganizationRepository) ListOrgsReturns(result1 []models.Organization, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	fake.listOrgsReturns = struct {
		result1 []models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationRepository) ListOrgsReturnsOnCall(i int, result1 []models.Organization, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	if fake.listOrgsReturnsOnCall == nil {
		fake.listOrgsReturnsOnCall = make(map[int]struct {
			result1 []models.Organization
			result2 error
		})
	}
	fake.listOrgsReturnsOnCall[i] = struct {
		result1 []models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeOrganizationRepository) Rename(arg1 string, arg2 string) error {
	fake.renameMutex.Lock()
	ret, specificReturn := fake.renameReturnsOnCall[len(fake.renameArgsForCall)]
	fake.renameArgsForCall = append(fake.renameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.RenameStub
	fakeReturns := fake.renameReturns
	fake.recordInvocation("Rename", []interface{}{arg1, arg2})
	fake.renameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOrganizationRepository) RenameCallCount() int {
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	return len(fake.renameArgsForCall)
}

func (fake *FakeOrganizationRepository) RenameCalls(stub func(string, string) error) {
	fake.renameMutex.Lock()
	defer fake.renameMutex.Unlock()
	fake.RenameStub = stub
}

func (fake *FakeOrganizationRepository) RenameArgsForCall(i int) (string, string) {
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	argsForCall := fake.renameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOrganizationRepository) RenameReturns(result1 error) {
	fake.renameMutex.Lock()
	defer fake.renameMutex.Unlock()
	fake.RenameStub = nil
	fake.renameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) RenameReturnsOnCall(i int, result1 error) {
	fake.renameMutex.Lock()
	defer fake.renameMutex.Unlock()
	fake.RenameStub = nil
	if fake.renameReturnsOnCall == nil {
		fake.renameReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.renameReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) SharePrivateDomain(arg1 string, arg2 string) error {
	fake.sharePrivateDomainMutex.Lock()
	ret, specificReturn := fake.sharePrivateDomainReturnsOnCall[len(fake.sharePrivateDomainArgsForCall)]
	fake.sharePrivateDomainArgsForCall = append(fake.sharePrivateDomainArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.SharePrivateDomainStub
	fakeReturns := fake.sharePrivateDomainReturns
	fake.recordInvocation("SharePrivateDomain", []interface{}{arg1, arg2})
	fake.sharePrivateDomainMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOrganizationRepository) SharePrivateDomainCallCount() int {
	fake.sharePrivateDomainMutex.RLock()
	defer fake.sharePrivateDomainMutex.RUnlock()
	return len(fake.sharePrivateDomainArgsForCall)
}

func (fake *FakeOrganizationRepository) SharePrivateDomainCalls(stub func(string, string) error) {
	fake.sharePrivateDomainMutex.Lock()
	defer fake.sharePrivateDomainMutex.Unlock()
	fake.SharePrivateDomainStub = stub
}

func (fake *FakeOrganizationRepository) SharePrivateDomainArgsForCall(i int) (string, string) {
	fake.sharePrivateDomainMutex.RLock()
	defer fake.sharePrivateDomainMutex.RUnlock()
	argsForCall := fake.sharePrivateDomainArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOrganizationRepository) SharePrivateDomainReturns(result1 error) {
	fake.sharePrivateDomainMutex.Lock()
	defer fake.sharePrivateDomainMutex.Unlock()
	fake.SharePrivateDomainStub = nil
	fake.sharePrivateDomainReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) SharePrivateDomainReturnsOnCall(i int, result1 error) {
	fake.sharePrivateDomainMutex.Lock()
	defer fake.sharePrivateDomainMutex.Unlock()
	fake.SharePrivateDomainStub = nil
	if fake.sharePrivateDomainReturnsOnCall == nil {
		fake.sharePrivateDomainReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sharePrivateDomainReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) UnsharePrivateDomain(arg1 string, arg2 string) error {
	fake.unsharePrivateDomainMutex.Lock()
	ret, specificReturn := fake.unsharePrivateDomainReturnsOnCall[len(fake.unsharePrivateDomainArgsForCall)]
	fake.unsharePrivateDomainArgsForCall = append(fake.unsharePrivateDomainArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.UnsharePrivateDomainStub
	fakeReturns := fake.unsharePrivateDomainReturns
	fake.recordInvocation("UnsharePrivateDomain", []interface{}{arg1, arg2})
	fake.unsharePrivateDomainMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOrganizationRepository) UnsharePrivateDomainCallCount() int {
	fake.unsharePrivateDomainMutex.RLock()
	defer fake.unsharePrivateDomainMutex.RUnlock()
	return len(fake.unsharePrivateDomainArgsForCall)
}

func (fake *FakeOrganizationRepository) UnsharePrivateDomainCalls(stub func(string, string) error) {
	fake.unsharePrivateDomainMutex.Lock()
	defer fake.unsharePrivateDomainMutex.Unlock()
	fake.UnsharePrivateDomainStub = stub
}

func (fake *FakeOrganizationRepository) UnsharePrivateDomainArgsForCall(i int) (string, string) {
	fake.unsharePrivateDomainMutex.RLock()
	defer fake.unsharePrivateDomainMutex.RUnlock()
	argsForCall := fake.unsharePrivateDomainArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOrganizationRepository) UnsharePrivateDomainReturns(result1 error) {
	fake.unsharePrivateDomainMutex.Lock()
	defer fake.unsharePrivateDomainMutex.Unlock()
	fake.UnsharePrivateDomainStub = nil
	fake.unsharePrivateDomainReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) UnsharePrivateDomainReturnsOnCall(i int, result1 error) {
	fake.unsharePrivateDomainMutex.Lock()
	defer fake.unsharePrivateDomainMutex.Unlock()
	fake.UnsharePrivateDomainStub = nil
	if fake.unsharePrivateDomainReturnsOnCall == nil {
		fake.unsharePrivateDomainReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unsharePrivateDomainReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	fake.getManyOrgsByGUIDMutex.RLock()
	defer fake.getManyOrgsByGUIDMutex.RUnlock()
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	fake.sharePrivateDomainMutex.RLock()
	defer fake.sharePrivateDomainMutex.RUnlock()
	fake.unsharePrivateDomainMutex.RLock()
	defer fake.unsharePrivateDomainMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOrganizationRepository) recordInvocation(key string, args []interface{}) {
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

var _ organizations.OrganizationRepository = new(FakeOrganizationRepository)
