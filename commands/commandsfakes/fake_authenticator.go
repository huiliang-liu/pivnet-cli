// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/auth"
	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeAuthenticator struct {
	AuthenticateClientStub        func(auth.AuthClient) error
	authenticateClientMutex       sync.RWMutex
	authenticateClientArgsForCall []struct {
		arg1 auth.AuthClient
	}
	authenticateClientReturns struct {
		result1 error
	}
	authenticateClientReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthenticator) AuthenticateClient(arg1 auth.AuthClient) error {
	fake.authenticateClientMutex.Lock()
	ret, specificReturn := fake.authenticateClientReturnsOnCall[len(fake.authenticateClientArgsForCall)]
	fake.authenticateClientArgsForCall = append(fake.authenticateClientArgsForCall, struct {
		arg1 auth.AuthClient
	}{arg1})
	fake.recordInvocation("AuthenticateClient", []interface{}{arg1})
	fake.authenticateClientMutex.Unlock()
	if fake.AuthenticateClientStub != nil {
		return fake.AuthenticateClientStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.authenticateClientReturns
	return fakeReturns.result1
}

func (fake *FakeAuthenticator) AuthenticateClientCallCount() int {
	fake.authenticateClientMutex.RLock()
	defer fake.authenticateClientMutex.RUnlock()
	return len(fake.authenticateClientArgsForCall)
}

func (fake *FakeAuthenticator) AuthenticateClientCalls(stub func(auth.AuthClient) error) {
	fake.authenticateClientMutex.Lock()
	defer fake.authenticateClientMutex.Unlock()
	fake.AuthenticateClientStub = stub
}

func (fake *FakeAuthenticator) AuthenticateClientArgsForCall(i int) auth.AuthClient {
	fake.authenticateClientMutex.RLock()
	defer fake.authenticateClientMutex.RUnlock()
	argsForCall := fake.authenticateClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAuthenticator) AuthenticateClientReturns(result1 error) {
	fake.authenticateClientMutex.Lock()
	defer fake.authenticateClientMutex.Unlock()
	fake.AuthenticateClientStub = nil
	fake.authenticateClientReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthenticator) AuthenticateClientReturnsOnCall(i int, result1 error) {
	fake.authenticateClientMutex.Lock()
	defer fake.authenticateClientMutex.Unlock()
	fake.AuthenticateClientStub = nil
	if fake.authenticateClientReturnsOnCall == nil {
		fake.authenticateClientReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.authenticateClientReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthenticator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authenticateClientMutex.RLock()
	defer fake.authenticateClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthenticator) recordInvocation(key string, args []interface{}) {
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

var _ commands.Authenticator = new(FakeAuthenticator)
