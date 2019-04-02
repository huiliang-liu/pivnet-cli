// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeUserGroupClient struct {
	ListStub        func(productSlug string, releaseVersion string) error
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	listReturns struct {
		result1 error
	}
	listReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(userGroupID int) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		userGroupID int
	}
	getReturns struct {
		result1 error
	}
	getReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(name string, description string, members []string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		name        string
		description string
		members     []string
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(userGroupID int, name *string, description *string) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		userGroupID int
		name        *string
		description *string
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	AddToReleaseStub        func(productSlug string, releaseVersion string, userGroupID int) error
	addToReleaseMutex       sync.RWMutex
	addToReleaseArgsForCall []struct {
		productSlug    string
		releaseVersion string
		userGroupID    int
	}
	addToReleaseReturns struct {
		result1 error
	}
	addToReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(userGroupID int) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		userGroupID int
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveFromReleaseStub        func(productSlug string, releaseVersion string, userGroupID int) error
	removeFromReleaseMutex       sync.RWMutex
	removeFromReleaseArgsForCall []struct {
		productSlug    string
		releaseVersion string
		userGroupID    int
	}
	removeFromReleaseReturns struct {
		result1 error
	}
	removeFromReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	AddUserGroupMemberStub        func(userGroupID int, memberEmailAddress string, admin bool) error
	addUserGroupMemberMutex       sync.RWMutex
	addUserGroupMemberArgsForCall []struct {
		userGroupID        int
		memberEmailAddress string
		admin              bool
	}
	addUserGroupMemberReturns struct {
		result1 error
	}
	addUserGroupMemberReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveUserGroupMemberStub        func(userGroupID int, memberEmailAddress string) error
	removeUserGroupMemberMutex       sync.RWMutex
	removeUserGroupMemberArgsForCall []struct {
		userGroupID        int
		memberEmailAddress string
	}
	removeUserGroupMemberReturns struct {
		result1 error
	}
	removeUserGroupMemberReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserGroupClient) List(productSlug string, releaseVersion string) error {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("List", []interface{}{productSlug, releaseVersion})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(productSlug, releaseVersion)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.listReturns.result1
}

func (fake *FakeUserGroupClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeUserGroupClient) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].productSlug, fake.listArgsForCall[i].releaseVersion
}

func (fake *FakeUserGroupClient) ListReturns(result1 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) ListReturnsOnCall(i int, result1 error) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) Get(userGroupID int) error {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		userGroupID int
	}{userGroupID})
	fake.recordInvocation("Get", []interface{}{userGroupID})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(userGroupID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getReturns.result1
}

func (fake *FakeUserGroupClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeUserGroupClient) GetArgsForCall(i int) int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].userGroupID
}

func (fake *FakeUserGroupClient) GetReturns(result1 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) GetReturnsOnCall(i int, result1 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) Create(name string, description string, members []string) error {
	var membersCopy []string
	if members != nil {
		membersCopy = make([]string, len(members))
		copy(membersCopy, members)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		name        string
		description string
		members     []string
	}{name, description, membersCopy})
	fake.recordInvocation("Create", []interface{}{name, description, membersCopy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(name, description, members)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *FakeUserGroupClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeUserGroupClient) CreateArgsForCall(i int) (string, string, []string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].name, fake.createArgsForCall[i].description, fake.createArgsForCall[i].members
}

func (fake *FakeUserGroupClient) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) CreateReturnsOnCall(i int, result1 error) {
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

func (fake *FakeUserGroupClient) Update(userGroupID int, name *string, description *string) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		userGroupID int
		name        *string
		description *string
	}{userGroupID, name, description})
	fake.recordInvocation("Update", []interface{}{userGroupID, name, description})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(userGroupID, name, description)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateReturns.result1
}

func (fake *FakeUserGroupClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeUserGroupClient) UpdateArgsForCall(i int) (int, *string, *string) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].userGroupID, fake.updateArgsForCall[i].name, fake.updateArgsForCall[i].description
}

func (fake *FakeUserGroupClient) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) UpdateReturnsOnCall(i int, result1 error) {
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

func (fake *FakeUserGroupClient) AddToRelease(productSlug string, releaseVersion string, userGroupID int) error {
	fake.addToReleaseMutex.Lock()
	ret, specificReturn := fake.addToReleaseReturnsOnCall[len(fake.addToReleaseArgsForCall)]
	fake.addToReleaseArgsForCall = append(fake.addToReleaseArgsForCall, struct {
		productSlug    string
		releaseVersion string
		userGroupID    int
	}{productSlug, releaseVersion, userGroupID})
	fake.recordInvocation("AddToRelease", []interface{}{productSlug, releaseVersion, userGroupID})
	fake.addToReleaseMutex.Unlock()
	if fake.AddToReleaseStub != nil {
		return fake.AddToReleaseStub(productSlug, releaseVersion, userGroupID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addToReleaseReturns.result1
}

func (fake *FakeUserGroupClient) AddToReleaseCallCount() int {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	return len(fake.addToReleaseArgsForCall)
}

func (fake *FakeUserGroupClient) AddToReleaseArgsForCall(i int) (string, string, int) {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	return fake.addToReleaseArgsForCall[i].productSlug, fake.addToReleaseArgsForCall[i].releaseVersion, fake.addToReleaseArgsForCall[i].userGroupID
}

func (fake *FakeUserGroupClient) AddToReleaseReturns(result1 error) {
	fake.AddToReleaseStub = nil
	fake.addToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) AddToReleaseReturnsOnCall(i int, result1 error) {
	fake.AddToReleaseStub = nil
	if fake.addToReleaseReturnsOnCall == nil {
		fake.addToReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addToReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) Delete(userGroupID int) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		userGroupID int
	}{userGroupID})
	fake.recordInvocation("Delete", []interface{}{userGroupID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(userGroupID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeUserGroupClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeUserGroupClient) DeleteArgsForCall(i int) int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].userGroupID
}

func (fake *FakeUserGroupClient) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *FakeUserGroupClient) RemoveFromRelease(productSlug string, releaseVersion string, userGroupID int) error {
	fake.removeFromReleaseMutex.Lock()
	ret, specificReturn := fake.removeFromReleaseReturnsOnCall[len(fake.removeFromReleaseArgsForCall)]
	fake.removeFromReleaseArgsForCall = append(fake.removeFromReleaseArgsForCall, struct {
		productSlug    string
		releaseVersion string
		userGroupID    int
	}{productSlug, releaseVersion, userGroupID})
	fake.recordInvocation("RemoveFromRelease", []interface{}{productSlug, releaseVersion, userGroupID})
	fake.removeFromReleaseMutex.Unlock()
	if fake.RemoveFromReleaseStub != nil {
		return fake.RemoveFromReleaseStub(productSlug, releaseVersion, userGroupID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeFromReleaseReturns.result1
}

func (fake *FakeUserGroupClient) RemoveFromReleaseCallCount() int {
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	return len(fake.removeFromReleaseArgsForCall)
}

func (fake *FakeUserGroupClient) RemoveFromReleaseArgsForCall(i int) (string, string, int) {
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	return fake.removeFromReleaseArgsForCall[i].productSlug, fake.removeFromReleaseArgsForCall[i].releaseVersion, fake.removeFromReleaseArgsForCall[i].userGroupID
}

func (fake *FakeUserGroupClient) RemoveFromReleaseReturns(result1 error) {
	fake.RemoveFromReleaseStub = nil
	fake.removeFromReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) RemoveFromReleaseReturnsOnCall(i int, result1 error) {
	fake.RemoveFromReleaseStub = nil
	if fake.removeFromReleaseReturnsOnCall == nil {
		fake.removeFromReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeFromReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) AddUserGroupMember(userGroupID int, memberEmailAddress string, admin bool) error {
	fake.addUserGroupMemberMutex.Lock()
	ret, specificReturn := fake.addUserGroupMemberReturnsOnCall[len(fake.addUserGroupMemberArgsForCall)]
	fake.addUserGroupMemberArgsForCall = append(fake.addUserGroupMemberArgsForCall, struct {
		userGroupID        int
		memberEmailAddress string
		admin              bool
	}{userGroupID, memberEmailAddress, admin})
	fake.recordInvocation("AddUserGroupMember", []interface{}{userGroupID, memberEmailAddress, admin})
	fake.addUserGroupMemberMutex.Unlock()
	if fake.AddUserGroupMemberStub != nil {
		return fake.AddUserGroupMemberStub(userGroupID, memberEmailAddress, admin)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addUserGroupMemberReturns.result1
}

func (fake *FakeUserGroupClient) AddUserGroupMemberCallCount() int {
	fake.addUserGroupMemberMutex.RLock()
	defer fake.addUserGroupMemberMutex.RUnlock()
	return len(fake.addUserGroupMemberArgsForCall)
}

func (fake *FakeUserGroupClient) AddUserGroupMemberArgsForCall(i int) (int, string, bool) {
	fake.addUserGroupMemberMutex.RLock()
	defer fake.addUserGroupMemberMutex.RUnlock()
	return fake.addUserGroupMemberArgsForCall[i].userGroupID, fake.addUserGroupMemberArgsForCall[i].memberEmailAddress, fake.addUserGroupMemberArgsForCall[i].admin
}

func (fake *FakeUserGroupClient) AddUserGroupMemberReturns(result1 error) {
	fake.AddUserGroupMemberStub = nil
	fake.addUserGroupMemberReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) AddUserGroupMemberReturnsOnCall(i int, result1 error) {
	fake.AddUserGroupMemberStub = nil
	if fake.addUserGroupMemberReturnsOnCall == nil {
		fake.addUserGroupMemberReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addUserGroupMemberReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) RemoveUserGroupMember(userGroupID int, memberEmailAddress string) error {
	fake.removeUserGroupMemberMutex.Lock()
	ret, specificReturn := fake.removeUserGroupMemberReturnsOnCall[len(fake.removeUserGroupMemberArgsForCall)]
	fake.removeUserGroupMemberArgsForCall = append(fake.removeUserGroupMemberArgsForCall, struct {
		userGroupID        int
		memberEmailAddress string
	}{userGroupID, memberEmailAddress})
	fake.recordInvocation("RemoveUserGroupMember", []interface{}{userGroupID, memberEmailAddress})
	fake.removeUserGroupMemberMutex.Unlock()
	if fake.RemoveUserGroupMemberStub != nil {
		return fake.RemoveUserGroupMemberStub(userGroupID, memberEmailAddress)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeUserGroupMemberReturns.result1
}

func (fake *FakeUserGroupClient) RemoveUserGroupMemberCallCount() int {
	fake.removeUserGroupMemberMutex.RLock()
	defer fake.removeUserGroupMemberMutex.RUnlock()
	return len(fake.removeUserGroupMemberArgsForCall)
}

func (fake *FakeUserGroupClient) RemoveUserGroupMemberArgsForCall(i int) (int, string) {
	fake.removeUserGroupMemberMutex.RLock()
	defer fake.removeUserGroupMemberMutex.RUnlock()
	return fake.removeUserGroupMemberArgsForCall[i].userGroupID, fake.removeUserGroupMemberArgsForCall[i].memberEmailAddress
}

func (fake *FakeUserGroupClient) RemoveUserGroupMemberReturns(result1 error) {
	fake.RemoveUserGroupMemberStub = nil
	fake.removeUserGroupMemberReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) RemoveUserGroupMemberReturnsOnCall(i int, result1 error) {
	fake.RemoveUserGroupMemberStub = nil
	if fake.removeUserGroupMemberReturnsOnCall == nil {
		fake.removeUserGroupMemberReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeUserGroupMemberReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserGroupClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	fake.addUserGroupMemberMutex.RLock()
	defer fake.addUserGroupMemberMutex.RUnlock()
	fake.removeUserGroupMemberMutex.RLock()
	defer fake.removeUserGroupMemberMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserGroupClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.UserGroupClient = new(FakeUserGroupClient)
