// Code generated by counterfeiter. DO NOT EDIT.
package jumpstartfakes

import (
	"lf-agent/jumpstart"
	"sync"
)

type FakeDatabaseClient struct {
	ApplyGTIDPurgedStub        func() error
	applyGTIDPurgedMutex       sync.RWMutex
	applyGTIDPurgedArgsForCall []struct{}
	applyGTIDPurgedReturns     struct {
		result1 error
	}
	applyGTIDPurgedReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDatabaseClient) ApplyGTIDPurged() error {
	fake.applyGTIDPurgedMutex.Lock()
	ret, specificReturn := fake.applyGTIDPurgedReturnsOnCall[len(fake.applyGTIDPurgedArgsForCall)]
	fake.applyGTIDPurgedArgsForCall = append(fake.applyGTIDPurgedArgsForCall, struct{}{})
	fake.recordInvocation("ApplyGTIDPurged", []interface{}{})
	fake.applyGTIDPurgedMutex.Unlock()
	if fake.ApplyGTIDPurgedStub != nil {
		return fake.ApplyGTIDPurgedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.applyGTIDPurgedReturns.result1
}

func (fake *FakeDatabaseClient) ApplyGTIDPurgedCallCount() int {
	fake.applyGTIDPurgedMutex.RLock()
	defer fake.applyGTIDPurgedMutex.RUnlock()
	return len(fake.applyGTIDPurgedArgsForCall)
}

func (fake *FakeDatabaseClient) ApplyGTIDPurgedReturns(result1 error) {
	fake.ApplyGTIDPurgedStub = nil
	fake.applyGTIDPurgedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDatabaseClient) ApplyGTIDPurgedReturnsOnCall(i int, result1 error) {
	fake.ApplyGTIDPurgedStub = nil
	if fake.applyGTIDPurgedReturnsOnCall == nil {
		fake.applyGTIDPurgedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyGTIDPurgedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDatabaseClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyGTIDPurgedMutex.RLock()
	defer fake.applyGTIDPurgedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDatabaseClient) recordInvocation(key string, args []interface{}) {
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

var _ jumpstart.DatabaseClient = new(FakeDatabaseClient)
