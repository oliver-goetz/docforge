// SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by counterfeiter. DO NOT EDIT.
package reactorfakes

import (
	"net/url"
	"sync"

	"github.com/gardener/docforge/pkg/reactor"
)

type FakeValidator struct {
	ValidateLinkStub        func(*url.URL, string, string) bool
	validateLinkMutex       sync.RWMutex
	validateLinkArgsForCall []struct {
		arg1 *url.URL
		arg2 string
		arg3 string
	}
	validateLinkReturns struct {
		result1 bool
	}
	validateLinkReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeValidator) ValidateLink(arg1 *url.URL, arg2 string, arg3 string) bool {
	fake.validateLinkMutex.Lock()
	ret, specificReturn := fake.validateLinkReturnsOnCall[len(fake.validateLinkArgsForCall)]
	fake.validateLinkArgsForCall = append(fake.validateLinkArgsForCall, struct {
		arg1 *url.URL
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.ValidateLinkStub
	fakeReturns := fake.validateLinkReturns
	fake.recordInvocation("ValidateLink", []interface{}{arg1, arg2, arg3})
	fake.validateLinkMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeValidator) ValidateLinkCallCount() int {
	fake.validateLinkMutex.RLock()
	defer fake.validateLinkMutex.RUnlock()
	return len(fake.validateLinkArgsForCall)
}

func (fake *FakeValidator) ValidateLinkCalls(stub func(*url.URL, string, string) bool) {
	fake.validateLinkMutex.Lock()
	defer fake.validateLinkMutex.Unlock()
	fake.ValidateLinkStub = stub
}

func (fake *FakeValidator) ValidateLinkArgsForCall(i int) (*url.URL, string, string) {
	fake.validateLinkMutex.RLock()
	defer fake.validateLinkMutex.RUnlock()
	argsForCall := fake.validateLinkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeValidator) ValidateLinkReturns(result1 bool) {
	fake.validateLinkMutex.Lock()
	defer fake.validateLinkMutex.Unlock()
	fake.ValidateLinkStub = nil
	fake.validateLinkReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeValidator) ValidateLinkReturnsOnCall(i int, result1 bool) {
	fake.validateLinkMutex.Lock()
	defer fake.validateLinkMutex.Unlock()
	fake.ValidateLinkStub = nil
	if fake.validateLinkReturnsOnCall == nil {
		fake.validateLinkReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.validateLinkReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateLinkMutex.RLock()
	defer fake.validateLinkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeValidator) recordInvocation(key string, args []interface{}) {
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

var _ reactor.Validator = new(FakeValidator)
