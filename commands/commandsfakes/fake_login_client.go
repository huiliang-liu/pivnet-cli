// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeLoginClient struct {
	LoginStub        func(string, string, string) error
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	loginReturns struct {
		result1 error
	}
	loginReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLoginClient) Login(arg1 string, arg2 string, arg3 string) error {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Login", []interface{}{arg1, arg2, arg3})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.loginReturns
	return fakeReturns.result1
}

func (fake *FakeLoginClient) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeLoginClient) LoginCalls(stub func(string, string, string) error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = stub
}

func (fake *FakeLoginClient) LoginArgsForCall(i int) (string, string, string) {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	argsForCall := fake.loginArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLoginClient) LoginReturns(result1 error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLoginClient) LoginReturnsOnCall(i int, result1 error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLoginClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLoginClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.LoginClient = new(FakeLoginClient)
