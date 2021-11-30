// SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by counterfeiter. DO NOT EDIT.
package gitfakes

import (
	"context"
	"sync"

	"github.com/gardener/docforge/pkg/git"
	gita "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type FakeRepository struct {
	FetchContextStub        func(context.Context, *gita.FetchOptions) error
	fetchContextMutex       sync.RWMutex
	fetchContextArgsForCall []struct {
		arg1 context.Context
		arg2 *gita.FetchOptions
	}
	fetchContextReturns struct {
		result1 error
	}
	fetchContextReturnsOnCall map[int]struct {
		result1 error
	}
	ReferenceStub        func(plumbing.ReferenceName, bool) (*plumbing.Reference, error)
	referenceMutex       sync.RWMutex
	referenceArgsForCall []struct {
		arg1 plumbing.ReferenceName
		arg2 bool
	}
	referenceReturns struct {
		result1 *plumbing.Reference
		result2 error
	}
	referenceReturnsOnCall map[int]struct {
		result1 *plumbing.Reference
		result2 error
	}
	TagsStub        func() ([]string, error)
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct {
	}
	tagsReturns struct {
		result1 []string
		result2 error
	}
	tagsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	WorktreeStub        func() (git.RepositoryWorktree, error)
	worktreeMutex       sync.RWMutex
	worktreeArgsForCall []struct {
	}
	worktreeReturns struct {
		result1 git.RepositoryWorktree
		result2 error
	}
	worktreeReturnsOnCall map[int]struct {
		result1 git.RepositoryWorktree
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) FetchContext(arg1 context.Context, arg2 *gita.FetchOptions) error {
	fake.fetchContextMutex.Lock()
	ret, specificReturn := fake.fetchContextReturnsOnCall[len(fake.fetchContextArgsForCall)]
	fake.fetchContextArgsForCall = append(fake.fetchContextArgsForCall, struct {
		arg1 context.Context
		arg2 *gita.FetchOptions
	}{arg1, arg2})
	stub := fake.FetchContextStub
	fakeReturns := fake.fetchContextReturns
	fake.recordInvocation("FetchContext", []interface{}{arg1, arg2})
	fake.fetchContextMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRepository) FetchContextCallCount() int {
	fake.fetchContextMutex.RLock()
	defer fake.fetchContextMutex.RUnlock()
	return len(fake.fetchContextArgsForCall)
}

func (fake *FakeRepository) FetchContextCalls(stub func(context.Context, *gita.FetchOptions) error) {
	fake.fetchContextMutex.Lock()
	defer fake.fetchContextMutex.Unlock()
	fake.FetchContextStub = stub
}

func (fake *FakeRepository) FetchContextArgsForCall(i int) (context.Context, *gita.FetchOptions) {
	fake.fetchContextMutex.RLock()
	defer fake.fetchContextMutex.RUnlock()
	argsForCall := fake.fetchContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) FetchContextReturns(result1 error) {
	fake.fetchContextMutex.Lock()
	defer fake.fetchContextMutex.Unlock()
	fake.FetchContextStub = nil
	fake.fetchContextReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) FetchContextReturnsOnCall(i int, result1 error) {
	fake.fetchContextMutex.Lock()
	defer fake.fetchContextMutex.Unlock()
	fake.FetchContextStub = nil
	if fake.fetchContextReturnsOnCall == nil {
		fake.fetchContextReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.fetchContextReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) Reference(arg1 plumbing.ReferenceName, arg2 bool) (*plumbing.Reference, error) {
	fake.referenceMutex.Lock()
	ret, specificReturn := fake.referenceReturnsOnCall[len(fake.referenceArgsForCall)]
	fake.referenceArgsForCall = append(fake.referenceArgsForCall, struct {
		arg1 plumbing.ReferenceName
		arg2 bool
	}{arg1, arg2})
	stub := fake.ReferenceStub
	fakeReturns := fake.referenceReturns
	fake.recordInvocation("Reference", []interface{}{arg1, arg2})
	fake.referenceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) ReferenceCallCount() int {
	fake.referenceMutex.RLock()
	defer fake.referenceMutex.RUnlock()
	return len(fake.referenceArgsForCall)
}

func (fake *FakeRepository) ReferenceCalls(stub func(plumbing.ReferenceName, bool) (*plumbing.Reference, error)) {
	fake.referenceMutex.Lock()
	defer fake.referenceMutex.Unlock()
	fake.ReferenceStub = stub
}

func (fake *FakeRepository) ReferenceArgsForCall(i int) (plumbing.ReferenceName, bool) {
	fake.referenceMutex.RLock()
	defer fake.referenceMutex.RUnlock()
	argsForCall := fake.referenceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) ReferenceReturns(result1 *plumbing.Reference, result2 error) {
	fake.referenceMutex.Lock()
	defer fake.referenceMutex.Unlock()
	fake.ReferenceStub = nil
	fake.referenceReturns = struct {
		result1 *plumbing.Reference
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) ReferenceReturnsOnCall(i int, result1 *plumbing.Reference, result2 error) {
	fake.referenceMutex.Lock()
	defer fake.referenceMutex.Unlock()
	fake.ReferenceStub = nil
	if fake.referenceReturnsOnCall == nil {
		fake.referenceReturnsOnCall = make(map[int]struct {
			result1 *plumbing.Reference
			result2 error
		})
	}
	fake.referenceReturnsOnCall[i] = struct {
		result1 *plumbing.Reference
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Tags() ([]string, error) {
	fake.tagsMutex.Lock()
	ret, specificReturn := fake.tagsReturnsOnCall[len(fake.tagsArgsForCall)]
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct {
	}{})
	stub := fake.TagsStub
	fakeReturns := fake.tagsReturns
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *FakeRepository) TagsCalls(stub func() ([]string, error)) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = stub
}

func (fake *FakeRepository) TagsReturns(result1 []string, result2 error) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) TagsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	if fake.tagsReturnsOnCall == nil {
		fake.tagsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.tagsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Worktree() (git.RepositoryWorktree, error) {
	fake.worktreeMutex.Lock()
	ret, specificReturn := fake.worktreeReturnsOnCall[len(fake.worktreeArgsForCall)]
	fake.worktreeArgsForCall = append(fake.worktreeArgsForCall, struct {
	}{})
	stub := fake.WorktreeStub
	fakeReturns := fake.worktreeReturns
	fake.recordInvocation("Worktree", []interface{}{})
	fake.worktreeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) WorktreeCallCount() int {
	fake.worktreeMutex.RLock()
	defer fake.worktreeMutex.RUnlock()
	return len(fake.worktreeArgsForCall)
}

func (fake *FakeRepository) WorktreeCalls(stub func() (git.RepositoryWorktree, error)) {
	fake.worktreeMutex.Lock()
	defer fake.worktreeMutex.Unlock()
	fake.WorktreeStub = stub
}

func (fake *FakeRepository) WorktreeReturns(result1 git.RepositoryWorktree, result2 error) {
	fake.worktreeMutex.Lock()
	defer fake.worktreeMutex.Unlock()
	fake.WorktreeStub = nil
	fake.worktreeReturns = struct {
		result1 git.RepositoryWorktree
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) WorktreeReturnsOnCall(i int, result1 git.RepositoryWorktree, result2 error) {
	fake.worktreeMutex.Lock()
	defer fake.worktreeMutex.Unlock()
	fake.WorktreeStub = nil
	if fake.worktreeReturnsOnCall == nil {
		fake.worktreeReturnsOnCall = make(map[int]struct {
			result1 git.RepositoryWorktree
			result2 error
		})
	}
	fake.worktreeReturnsOnCall[i] = struct {
		result1 git.RepositoryWorktree
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchContextMutex.RLock()
	defer fake.fetchContextMutex.RUnlock()
	fake.referenceMutex.RLock()
	defer fake.referenceMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	fake.worktreeMutex.RLock()
	defer fake.worktreeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
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

var _ git.Repository = new(FakeRepository)