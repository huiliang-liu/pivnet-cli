// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeFileGroupClient struct {
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
	GetStub        func(productSlug string, productFileID int) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		productSlug   string
		productFileID int
	}
	getReturns struct {
		result1 error
	}
	getReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(productSlug string, name string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		productSlug string
		name        string
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(productSlug string, productFileID int, name *string) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		productSlug   string
		productFileID int
		name          *string
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(productSlug string, productFileID int) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		productSlug   string
		productFileID int
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	AddToReleaseStub        func(productSlug string, productFileID int, releaseVersion string) error
	addToReleaseMutex       sync.RWMutex
	addToReleaseArgsForCall []struct {
		productSlug    string
		productFileID  int
		releaseVersion string
	}
	addToReleaseReturns struct {
		result1 error
	}
	addToReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveFromReleaseStub        func(productSlug string, productFileID int, releaseVersion string) error
	removeFromReleaseMutex       sync.RWMutex
	removeFromReleaseArgsForCall []struct {
		productSlug    string
		productFileID  int
		releaseVersion string
	}
	removeFromReleaseReturns struct {
		result1 error
	}
	removeFromReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFileGroupClient) List(productSlug string, releaseVersion string) error {
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

func (fake *FakeFileGroupClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeFileGroupClient) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].productSlug, fake.listArgsForCall[i].releaseVersion
}

func (fake *FakeFileGroupClient) ListReturns(result1 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileGroupClient) ListReturnsOnCall(i int, result1 error) {
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

func (fake *FakeFileGroupClient) Get(productSlug string, productFileID int) error {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		productSlug   string
		productFileID int
	}{productSlug, productFileID})
	fake.recordInvocation("Get", []interface{}{productSlug, productFileID})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(productSlug, productFileID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getReturns.result1
}

func (fake *FakeFileGroupClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeFileGroupClient) GetArgsForCall(i int) (string, int) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].productSlug, fake.getArgsForCall[i].productFileID
}

func (fake *FakeFileGroupClient) GetReturns(result1 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileGroupClient) GetReturnsOnCall(i int, result1 error) {
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

func (fake *FakeFileGroupClient) Create(productSlug string, name string) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		productSlug string
		name        string
	}{productSlug, name})
	fake.recordInvocation("Create", []interface{}{productSlug, name})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(productSlug, name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *FakeFileGroupClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeFileGroupClient) CreateArgsForCall(i int) (string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].productSlug, fake.createArgsForCall[i].name
}

func (fake *FakeFileGroupClient) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileGroupClient) CreateReturnsOnCall(i int, result1 error) {
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

func (fake *FakeFileGroupClient) Update(productSlug string, productFileID int, name *string) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		productSlug   string
		productFileID int
		name          *string
	}{productSlug, productFileID, name})
	fake.recordInvocation("Update", []interface{}{productSlug, productFileID, name})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(productSlug, productFileID, name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateReturns.result1
}

func (fake *FakeFileGroupClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeFileGroupClient) UpdateArgsForCall(i int) (string, int, *string) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].productSlug, fake.updateArgsForCall[i].productFileID, fake.updateArgsForCall[i].name
}

func (fake *FakeFileGroupClient) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileGroupClient) UpdateReturnsOnCall(i int, result1 error) {
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

func (fake *FakeFileGroupClient) Delete(productSlug string, productFileID int) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		productSlug   string
		productFileID int
	}{productSlug, productFileID})
	fake.recordInvocation("Delete", []interface{}{productSlug, productFileID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(productSlug, productFileID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeFileGroupClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeFileGroupClient) DeleteArgsForCall(i int) (string, int) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].productSlug, fake.deleteArgsForCall[i].productFileID
}

func (fake *FakeFileGroupClient) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileGroupClient) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *FakeFileGroupClient) AddToRelease(productSlug string, productFileID int, releaseVersion string) error {
	fake.addToReleaseMutex.Lock()
	ret, specificReturn := fake.addToReleaseReturnsOnCall[len(fake.addToReleaseArgsForCall)]
	fake.addToReleaseArgsForCall = append(fake.addToReleaseArgsForCall, struct {
		productSlug    string
		productFileID  int
		releaseVersion string
	}{productSlug, productFileID, releaseVersion})
	fake.recordInvocation("AddToRelease", []interface{}{productSlug, productFileID, releaseVersion})
	fake.addToReleaseMutex.Unlock()
	if fake.AddToReleaseStub != nil {
		return fake.AddToReleaseStub(productSlug, productFileID, releaseVersion)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addToReleaseReturns.result1
}

func (fake *FakeFileGroupClient) AddToReleaseCallCount() int {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	return len(fake.addToReleaseArgsForCall)
}

func (fake *FakeFileGroupClient) AddToReleaseArgsForCall(i int) (string, int, string) {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	return fake.addToReleaseArgsForCall[i].productSlug, fake.addToReleaseArgsForCall[i].productFileID, fake.addToReleaseArgsForCall[i].releaseVersion
}

func (fake *FakeFileGroupClient) AddToReleaseReturns(result1 error) {
	fake.AddToReleaseStub = nil
	fake.addToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileGroupClient) AddToReleaseReturnsOnCall(i int, result1 error) {
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

func (fake *FakeFileGroupClient) RemoveFromRelease(productSlug string, productFileID int, releaseVersion string) error {
	fake.removeFromReleaseMutex.Lock()
	ret, specificReturn := fake.removeFromReleaseReturnsOnCall[len(fake.removeFromReleaseArgsForCall)]
	fake.removeFromReleaseArgsForCall = append(fake.removeFromReleaseArgsForCall, struct {
		productSlug    string
		productFileID  int
		releaseVersion string
	}{productSlug, productFileID, releaseVersion})
	fake.recordInvocation("RemoveFromRelease", []interface{}{productSlug, productFileID, releaseVersion})
	fake.removeFromReleaseMutex.Unlock()
	if fake.RemoveFromReleaseStub != nil {
		return fake.RemoveFromReleaseStub(productSlug, productFileID, releaseVersion)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeFromReleaseReturns.result1
}

func (fake *FakeFileGroupClient) RemoveFromReleaseCallCount() int {
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	return len(fake.removeFromReleaseArgsForCall)
}

func (fake *FakeFileGroupClient) RemoveFromReleaseArgsForCall(i int) (string, int, string) {
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	return fake.removeFromReleaseArgsForCall[i].productSlug, fake.removeFromReleaseArgsForCall[i].productFileID, fake.removeFromReleaseArgsForCall[i].releaseVersion
}

func (fake *FakeFileGroupClient) RemoveFromReleaseReturns(result1 error) {
	fake.RemoveFromReleaseStub = nil
	fake.removeFromReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFileGroupClient) RemoveFromReleaseReturnsOnCall(i int, result1 error) {
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

func (fake *FakeFileGroupClient) Invocations() map[string][][]interface{} {
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
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFileGroupClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.FileGroupClient = new(FakeFileGroupClient)
