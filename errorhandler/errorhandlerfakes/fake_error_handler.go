// Code generated by counterfeiter. DO NOT EDIT.
package errorhandlerfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/errorhandler"
)

type FakeErrorHandler struct {
	HandleErrorStub        func(err error) error
	handleErrorMutex       sync.RWMutex
	handleErrorArgsForCall []struct {
		err error
	}
	handleErrorReturns struct {
		result1 error
	}
	handleErrorReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeErrorHandler) HandleError(err error) error {
	fake.handleErrorMutex.Lock()
	ret, specificReturn := fake.handleErrorReturnsOnCall[len(fake.handleErrorArgsForCall)]
	fake.handleErrorArgsForCall = append(fake.handleErrorArgsForCall, struct {
		err error
	}{err})
	fake.recordInvocation("HandleError", []interface{}{err})
	fake.handleErrorMutex.Unlock()
	if fake.HandleErrorStub != nil {
		return fake.HandleErrorStub(err)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.handleErrorReturns.result1
}

func (fake *FakeErrorHandler) HandleErrorCallCount() int {
	fake.handleErrorMutex.RLock()
	defer fake.handleErrorMutex.RUnlock()
	return len(fake.handleErrorArgsForCall)
}

func (fake *FakeErrorHandler) HandleErrorArgsForCall(i int) error {
	fake.handleErrorMutex.RLock()
	defer fake.handleErrorMutex.RUnlock()
	return fake.handleErrorArgsForCall[i].err
}

func (fake *FakeErrorHandler) HandleErrorReturns(result1 error) {
	fake.HandleErrorStub = nil
	fake.handleErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeErrorHandler) HandleErrorReturnsOnCall(i int, result1 error) {
	fake.HandleErrorStub = nil
	if fake.handleErrorReturnsOnCall == nil {
		fake.handleErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.handleErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeErrorHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleErrorMutex.RLock()
	defer fake.handleErrorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeErrorHandler) recordInvocation(key string, args []interface{}) {
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

var _ errorhandler.ErrorHandler = new(FakeErrorHandler)
