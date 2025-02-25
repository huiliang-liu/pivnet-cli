// Code generated by counterfeiter. DO NOT EDIT.
package eulafakes

import (
	"sync"

	"github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/pivnet-cli/commands/eula"
)

type FakePivnetClient struct {
	AcceptEULAStub        func(productSlug string, releaseID int) error
	acceptEULAMutex       sync.RWMutex
	acceptEULAArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	acceptEULAReturns struct {
		result1 error
	}
	acceptEULAReturnsOnCall map[int]struct {
		result1 error
	}
	EULAsStub        func() ([]pivnet.EULA, error)
	eULAsMutex       sync.RWMutex
	eULAsArgsForCall []struct{}
	eULAsReturns     struct {
		result1 []pivnet.EULA
		result2 error
	}
	eULAsReturnsOnCall map[int]struct {
		result1 []pivnet.EULA
		result2 error
	}
	EULAStub        func(eulaSlug string) (pivnet.EULA, error)
	eULAMutex       sync.RWMutex
	eULAArgsForCall []struct {
		eulaSlug string
	}
	eULAReturns struct {
		result1 pivnet.EULA
		result2 error
	}
	eULAReturnsOnCall map[int]struct {
		result1 pivnet.EULA
		result2 error
	}
	ReleaseForVersionStub        func(productSlug string, releaseVersion string) (pivnet.Release, error)
	releaseForVersionMutex       sync.RWMutex
	releaseForVersionArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	releaseForVersionReturns struct {
		result1 pivnet.Release
		result2 error
	}
	releaseForVersionReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) AcceptEULA(productSlug string, releaseID int) error {
	fake.acceptEULAMutex.Lock()
	ret, specificReturn := fake.acceptEULAReturnsOnCall[len(fake.acceptEULAArgsForCall)]
	fake.acceptEULAArgsForCall = append(fake.acceptEULAArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("AcceptEULA", []interface{}{productSlug, releaseID})
	fake.acceptEULAMutex.Unlock()
	if fake.AcceptEULAStub != nil {
		return fake.AcceptEULAStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.acceptEULAReturns.result1
}

func (fake *FakePivnetClient) AcceptEULACallCount() int {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return len(fake.acceptEULAArgsForCall)
}

func (fake *FakePivnetClient) AcceptEULAArgsForCall(i int) (string, int) {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return fake.acceptEULAArgsForCall[i].productSlug, fake.acceptEULAArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) AcceptEULAReturns(result1 error) {
	fake.AcceptEULAStub = nil
	fake.acceptEULAReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AcceptEULAReturnsOnCall(i int, result1 error) {
	fake.AcceptEULAStub = nil
	if fake.acceptEULAReturnsOnCall == nil {
		fake.acceptEULAReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.acceptEULAReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) EULAs() ([]pivnet.EULA, error) {
	fake.eULAsMutex.Lock()
	ret, specificReturn := fake.eULAsReturnsOnCall[len(fake.eULAsArgsForCall)]
	fake.eULAsArgsForCall = append(fake.eULAsArgsForCall, struct{}{})
	fake.recordInvocation("EULAs", []interface{}{})
	fake.eULAsMutex.Unlock()
	if fake.EULAsStub != nil {
		return fake.EULAsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.eULAsReturns.result1, fake.eULAsReturns.result2
}

func (fake *FakePivnetClient) EULAsCallCount() int {
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	return len(fake.eULAsArgsForCall)
}

func (fake *FakePivnetClient) EULAsReturns(result1 []pivnet.EULA, result2 error) {
	fake.EULAsStub = nil
	fake.eULAsReturns = struct {
		result1 []pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) EULAsReturnsOnCall(i int, result1 []pivnet.EULA, result2 error) {
	fake.EULAsStub = nil
	if fake.eULAsReturnsOnCall == nil {
		fake.eULAsReturnsOnCall = make(map[int]struct {
			result1 []pivnet.EULA
			result2 error
		})
	}
	fake.eULAsReturnsOnCall[i] = struct {
		result1 []pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) EULA(eulaSlug string) (pivnet.EULA, error) {
	fake.eULAMutex.Lock()
	ret, specificReturn := fake.eULAReturnsOnCall[len(fake.eULAArgsForCall)]
	fake.eULAArgsForCall = append(fake.eULAArgsForCall, struct {
		eulaSlug string
	}{eulaSlug})
	fake.recordInvocation("EULA", []interface{}{eulaSlug})
	fake.eULAMutex.Unlock()
	if fake.EULAStub != nil {
		return fake.EULAStub(eulaSlug)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.eULAReturns.result1, fake.eULAReturns.result2
}

func (fake *FakePivnetClient) EULACallCount() int {
	fake.eULAMutex.RLock()
	defer fake.eULAMutex.RUnlock()
	return len(fake.eULAArgsForCall)
}

func (fake *FakePivnetClient) EULAArgsForCall(i int) string {
	fake.eULAMutex.RLock()
	defer fake.eULAMutex.RUnlock()
	return fake.eULAArgsForCall[i].eulaSlug
}

func (fake *FakePivnetClient) EULAReturns(result1 pivnet.EULA, result2 error) {
	fake.EULAStub = nil
	fake.eULAReturns = struct {
		result1 pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) EULAReturnsOnCall(i int, result1 pivnet.EULA, result2 error) {
	fake.EULAStub = nil
	if fake.eULAReturnsOnCall == nil {
		fake.eULAReturnsOnCall = make(map[int]struct {
			result1 pivnet.EULA
			result2 error
		})
	}
	fake.eULAReturnsOnCall[i] = struct {
		result1 pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersion(productSlug string, releaseVersion string) (pivnet.Release, error) {
	fake.releaseForVersionMutex.Lock()
	ret, specificReturn := fake.releaseForVersionReturnsOnCall[len(fake.releaseForVersionArgsForCall)]
	fake.releaseForVersionArgsForCall = append(fake.releaseForVersionArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("ReleaseForVersion", []interface{}{productSlug, releaseVersion})
	fake.releaseForVersionMutex.Unlock()
	if fake.ReleaseForVersionStub != nil {
		return fake.ReleaseForVersionStub(productSlug, releaseVersion)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.releaseForVersionReturns.result1, fake.releaseForVersionReturns.result2
}

func (fake *FakePivnetClient) ReleaseForVersionCallCount() int {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return len(fake.releaseForVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForVersionArgsForCall(i int) (string, string) {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return fake.releaseForVersionArgsForCall[i].productSlug, fake.releaseForVersionArgsForCall[i].releaseVersion
}

func (fake *FakePivnetClient) ReleaseForVersionReturns(result1 pivnet.Release, result2 error) {
	fake.ReleaseForVersionStub = nil
	fake.releaseForVersionReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersionReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.ReleaseForVersionStub = nil
	if fake.releaseForVersionReturnsOnCall == nil {
		fake.releaseForVersionReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.releaseForVersionReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	fake.eULAMutex.RLock()
	defer fake.eULAMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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

var _ eula.PivnetClient = new(FakePivnetClient)
