// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"clerks/domain"
	"sync"
)

type FakeRandomUserRepository struct {
	FetchUsersStub        func() ([]domain.User, error)
	fetchUsersMutex       sync.RWMutex
	fetchUsersArgsForCall []struct {
	}
	fetchUsersReturns struct {
		result1 []domain.User
		result2 error
	}
	fetchUsersReturnsOnCall map[int]struct {
		result1 []domain.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRandomUserRepository) FetchUsers() ([]domain.User, error) {
	fake.fetchUsersMutex.Lock()
	ret, specificReturn := fake.fetchUsersReturnsOnCall[len(fake.fetchUsersArgsForCall)]
	fake.fetchUsersArgsForCall = append(fake.fetchUsersArgsForCall, struct {
	}{})
	stub := fake.FetchUsersStub
	fakeReturns := fake.fetchUsersReturns
	fake.recordInvocation("FetchUsers", []interface{}{})
	fake.fetchUsersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRandomUserRepository) FetchUsersCallCount() int {
	fake.fetchUsersMutex.RLock()
	defer fake.fetchUsersMutex.RUnlock()
	return len(fake.fetchUsersArgsForCall)
}

func (fake *FakeRandomUserRepository) FetchUsersCalls(stub func() ([]domain.User, error)) {
	fake.fetchUsersMutex.Lock()
	defer fake.fetchUsersMutex.Unlock()
	fake.FetchUsersStub = stub
}

func (fake *FakeRandomUserRepository) FetchUsersReturns(result1 []domain.User, result2 error) {
	fake.fetchUsersMutex.Lock()
	defer fake.fetchUsersMutex.Unlock()
	fake.FetchUsersStub = nil
	fake.fetchUsersReturns = struct {
		result1 []domain.User
		result2 error
	}{result1, result2}
}

func (fake *FakeRandomUserRepository) FetchUsersReturnsOnCall(i int, result1 []domain.User, result2 error) {
	fake.fetchUsersMutex.Lock()
	defer fake.fetchUsersMutex.Unlock()
	fake.FetchUsersStub = nil
	if fake.fetchUsersReturnsOnCall == nil {
		fake.fetchUsersReturnsOnCall = make(map[int]struct {
			result1 []domain.User
			result2 error
		})
	}
	fake.fetchUsersReturnsOnCall[i] = struct {
		result1 []domain.User
		result2 error
	}{result1, result2}
}

func (fake *FakeRandomUserRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchUsersMutex.RLock()
	defer fake.fetchUsersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRandomUserRepository) recordInvocation(key string, args []interface{}) {
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

var _ domain.RandomUserRepository = new(FakeRandomUserRepository)
