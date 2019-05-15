// Code generated by counterfeiter. DO NOT EDIT.
package botfakes

import (
	"sync"

	"github.com/hortbot/hortbot/internal/bot"
)

type FakeMessageSender struct {
	SendMessageStub        func(string, string) error
	sendMessageMutex       sync.RWMutex
	sendMessageArgsForCall []struct {
		arg1 string
		arg2 string
	}
	sendMessageReturns struct {
		result1 error
	}
	sendMessageReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMessageSender) SendMessage(arg1 string, arg2 string) error {
	fake.sendMessageMutex.Lock()
	ret, specificReturn := fake.sendMessageReturnsOnCall[len(fake.sendMessageArgsForCall)]
	fake.sendMessageArgsForCall = append(fake.sendMessageArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SendMessage", []interface{}{arg1, arg2})
	fake.sendMessageMutex.Unlock()
	if fake.SendMessageStub != nil {
		return fake.SendMessageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendMessageReturns
	return fakeReturns.result1
}

func (fake *FakeMessageSender) SendMessageCallCount() int {
	fake.sendMessageMutex.RLock()
	defer fake.sendMessageMutex.RUnlock()
	return len(fake.sendMessageArgsForCall)
}

func (fake *FakeMessageSender) SendMessageCalls(stub func(string, string) error) {
	fake.sendMessageMutex.Lock()
	defer fake.sendMessageMutex.Unlock()
	fake.SendMessageStub = stub
}

func (fake *FakeMessageSender) SendMessageArgsForCall(i int) (string, string) {
	fake.sendMessageMutex.RLock()
	defer fake.sendMessageMutex.RUnlock()
	argsForCall := fake.sendMessageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMessageSender) SendMessageReturns(result1 error) {
	fake.sendMessageMutex.Lock()
	defer fake.sendMessageMutex.Unlock()
	fake.SendMessageStub = nil
	fake.sendMessageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMessageSender) SendMessageReturnsOnCall(i int, result1 error) {
	fake.sendMessageMutex.Lock()
	defer fake.sendMessageMutex.Unlock()
	fake.SendMessageStub = nil
	if fake.sendMessageReturnsOnCall == nil {
		fake.sendMessageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendMessageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMessageSender) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendMessageMutex.RLock()
	defer fake.sendMessageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMessageSender) recordInvocation(key string, args []interface{}) {
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

var _ bot.MessageSender = new(FakeMessageSender)
