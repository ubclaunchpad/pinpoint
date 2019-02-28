// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"github.com/ubclaunchpad/pinpoint/protobuf/response"
	"google.golang.org/grpc"
)

type FakeCoreClient struct {
	CreateAccountStub        func(context.Context, *request.CreateAccount, ...grpc.CallOption) (*response.Message, error)
	createAccountMutex       sync.RWMutex
	createAccountArgsForCall []struct {
		arg1 context.Context
		arg2 *request.CreateAccount
		arg3 []grpc.CallOption
	}
	createAccountReturns struct {
		result1 *response.Message
		result2 error
	}
	createAccountReturnsOnCall map[int]struct {
		result1 *response.Message
		result2 error
	}
	CreateEventStub        func(context.Context, *request.CreateEvent, ...grpc.CallOption) (*response.Message, error)
	createEventMutex       sync.RWMutex
	createEventArgsForCall []struct {
		arg1 context.Context
		arg2 *request.CreateEvent
		arg3 []grpc.CallOption
	}
	createEventReturns struct {
		result1 *response.Message
		result2 error
	}
	createEventReturnsOnCall map[int]struct {
		result1 *response.Message
		result2 error
	}
	GetStatusStub        func(context.Context, *request.Status, ...grpc.CallOption) (*response.Status, error)
	getStatusMutex       sync.RWMutex
	getStatusArgsForCall []struct {
		arg1 context.Context
		arg2 *request.Status
		arg3 []grpc.CallOption
	}
	getStatusReturns struct {
		result1 *response.Status
		result2 error
	}
	getStatusReturnsOnCall map[int]struct {
		result1 *response.Status
		result2 error
	}
	HandshakeStub        func(context.Context, *request.Empty, ...grpc.CallOption) (*response.Empty, error)
	handshakeMutex       sync.RWMutex
	handshakeArgsForCall []struct {
		arg1 context.Context
		arg2 *request.Empty
		arg3 []grpc.CallOption
	}
	handshakeReturns struct {
		result1 *response.Empty
		result2 error
	}
	handshakeReturnsOnCall map[int]struct {
		result1 *response.Empty
		result2 error
	}
	LoginStub        func(context.Context, *request.Login, ...grpc.CallOption) (*response.Message, error)
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		arg1 context.Context
		arg2 *request.Login
		arg3 []grpc.CallOption
	}
	loginReturns struct {
		result1 *response.Message
		result2 error
	}
	loginReturnsOnCall map[int]struct {
		result1 *response.Message
		result2 error
	}
	VerifyStub        func(context.Context, *request.Verify, ...grpc.CallOption) (*response.Message, error)
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		arg1 context.Context
		arg2 *request.Verify
		arg3 []grpc.CallOption
	}
	verifyReturns struct {
		result1 *response.Message
		result2 error
	}
	verifyReturnsOnCall map[int]struct {
		result1 *response.Message
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCoreClient) CreateAccount(arg1 context.Context, arg2 *request.CreateAccount, arg3 ...grpc.CallOption) (*response.Message, error) {
	fake.createAccountMutex.Lock()
	ret, specificReturn := fake.createAccountReturnsOnCall[len(fake.createAccountArgsForCall)]
	fake.createAccountArgsForCall = append(fake.createAccountArgsForCall, struct {
		arg1 context.Context
		arg2 *request.CreateAccount
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateAccount", []interface{}{arg1, arg2, arg3})
	fake.createAccountMutex.Unlock()
	if fake.CreateAccountStub != nil {
		return fake.CreateAccountStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createAccountReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCoreClient) CreateAccountCallCount() int {
	fake.createAccountMutex.RLock()
	defer fake.createAccountMutex.RUnlock()
	return len(fake.createAccountArgsForCall)
}

func (fake *FakeCoreClient) CreateAccountCalls(stub func(context.Context, *request.CreateAccount, ...grpc.CallOption) (*response.Message, error)) {
	fake.createAccountMutex.Lock()
	defer fake.createAccountMutex.Unlock()
	fake.CreateAccountStub = stub
}

func (fake *FakeCoreClient) CreateAccountArgsForCall(i int) (context.Context, *request.CreateAccount, []grpc.CallOption) {
	fake.createAccountMutex.RLock()
	defer fake.createAccountMutex.RUnlock()
	argsForCall := fake.createAccountArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCoreClient) CreateAccountReturns(result1 *response.Message, result2 error) {
	fake.createAccountMutex.Lock()
	defer fake.createAccountMutex.Unlock()
	fake.CreateAccountStub = nil
	fake.createAccountReturns = struct {
		result1 *response.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) CreateAccountReturnsOnCall(i int, result1 *response.Message, result2 error) {
	fake.createAccountMutex.Lock()
	defer fake.createAccountMutex.Unlock()
	fake.CreateAccountStub = nil
	if fake.createAccountReturnsOnCall == nil {
		fake.createAccountReturnsOnCall = make(map[int]struct {
			result1 *response.Message
			result2 error
		})
	}
	fake.createAccountReturnsOnCall[i] = struct {
		result1 *response.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) CreateEvent(arg1 context.Context, arg2 *request.CreateEvent, arg3 ...grpc.CallOption) (*response.Message, error) {
	fake.createEventMutex.Lock()
	ret, specificReturn := fake.createEventReturnsOnCall[len(fake.createEventArgsForCall)]
	fake.createEventArgsForCall = append(fake.createEventArgsForCall, struct {
		arg1 context.Context
		arg2 *request.CreateEvent
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateEvent", []interface{}{arg1, arg2, arg3})
	fake.createEventMutex.Unlock()
	if fake.CreateEventStub != nil {
		return fake.CreateEventStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createEventReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCoreClient) CreateEventCallCount() int {
	fake.createEventMutex.RLock()
	defer fake.createEventMutex.RUnlock()
	return len(fake.createEventArgsForCall)
}

func (fake *FakeCoreClient) CreateEventCalls(stub func(context.Context, *request.CreateEvent, ...grpc.CallOption) (*response.Message, error)) {
	fake.createEventMutex.Lock()
	defer fake.createEventMutex.Unlock()
	fake.CreateEventStub = stub
}

func (fake *FakeCoreClient) CreateEventArgsForCall(i int) (context.Context, *request.CreateEvent, []grpc.CallOption) {
	fake.createEventMutex.RLock()
	defer fake.createEventMutex.RUnlock()
	argsForCall := fake.createEventArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCoreClient) CreateEventReturns(result1 *response.Message, result2 error) {
	fake.createEventMutex.Lock()
	defer fake.createEventMutex.Unlock()
	fake.CreateEventStub = nil
	fake.createEventReturns = struct {
		result1 *response.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) CreateEventReturnsOnCall(i int, result1 *response.Message, result2 error) {
	fake.createEventMutex.Lock()
	defer fake.createEventMutex.Unlock()
	fake.CreateEventStub = nil
	if fake.createEventReturnsOnCall == nil {
		fake.createEventReturnsOnCall = make(map[int]struct {
			result1 *response.Message
			result2 error
		})
	}
	fake.createEventReturnsOnCall[i] = struct {
		result1 *response.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) GetStatus(arg1 context.Context, arg2 *request.Status, arg3 ...grpc.CallOption) (*response.Status, error) {
	fake.getStatusMutex.Lock()
	ret, specificReturn := fake.getStatusReturnsOnCall[len(fake.getStatusArgsForCall)]
	fake.getStatusArgsForCall = append(fake.getStatusArgsForCall, struct {
		arg1 context.Context
		arg2 *request.Status
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetStatus", []interface{}{arg1, arg2, arg3})
	fake.getStatusMutex.Unlock()
	if fake.GetStatusStub != nil {
		return fake.GetStatusStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStatusReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCoreClient) GetStatusCallCount() int {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	return len(fake.getStatusArgsForCall)
}

func (fake *FakeCoreClient) GetStatusCalls(stub func(context.Context, *request.Status, ...grpc.CallOption) (*response.Status, error)) {
	fake.getStatusMutex.Lock()
	defer fake.getStatusMutex.Unlock()
	fake.GetStatusStub = stub
}

func (fake *FakeCoreClient) GetStatusArgsForCall(i int) (context.Context, *request.Status, []grpc.CallOption) {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	argsForCall := fake.getStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCoreClient) GetStatusReturns(result1 *response.Status, result2 error) {
	fake.getStatusMutex.Lock()
	defer fake.getStatusMutex.Unlock()
	fake.GetStatusStub = nil
	fake.getStatusReturns = struct {
		result1 *response.Status
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) GetStatusReturnsOnCall(i int, result1 *response.Status, result2 error) {
	fake.getStatusMutex.Lock()
	defer fake.getStatusMutex.Unlock()
	fake.GetStatusStub = nil
	if fake.getStatusReturnsOnCall == nil {
		fake.getStatusReturnsOnCall = make(map[int]struct {
			result1 *response.Status
			result2 error
		})
	}
	fake.getStatusReturnsOnCall[i] = struct {
		result1 *response.Status
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) Handshake(arg1 context.Context, arg2 *request.Empty, arg3 ...grpc.CallOption) (*response.Empty, error) {
	fake.handshakeMutex.Lock()
	ret, specificReturn := fake.handshakeReturnsOnCall[len(fake.handshakeArgsForCall)]
	fake.handshakeArgsForCall = append(fake.handshakeArgsForCall, struct {
		arg1 context.Context
		arg2 *request.Empty
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("Handshake", []interface{}{arg1, arg2, arg3})
	fake.handshakeMutex.Unlock()
	if fake.HandshakeStub != nil {
		return fake.HandshakeStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.handshakeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCoreClient) HandshakeCallCount() int {
	fake.handshakeMutex.RLock()
	defer fake.handshakeMutex.RUnlock()
	return len(fake.handshakeArgsForCall)
}

func (fake *FakeCoreClient) HandshakeCalls(stub func(context.Context, *request.Empty, ...grpc.CallOption) (*response.Empty, error)) {
	fake.handshakeMutex.Lock()
	defer fake.handshakeMutex.Unlock()
	fake.HandshakeStub = stub
}

func (fake *FakeCoreClient) HandshakeArgsForCall(i int) (context.Context, *request.Empty, []grpc.CallOption) {
	fake.handshakeMutex.RLock()
	defer fake.handshakeMutex.RUnlock()
	argsForCall := fake.handshakeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCoreClient) HandshakeReturns(result1 *response.Empty, result2 error) {
	fake.handshakeMutex.Lock()
	defer fake.handshakeMutex.Unlock()
	fake.HandshakeStub = nil
	fake.handshakeReturns = struct {
		result1 *response.Empty
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) HandshakeReturnsOnCall(i int, result1 *response.Empty, result2 error) {
	fake.handshakeMutex.Lock()
	defer fake.handshakeMutex.Unlock()
	fake.HandshakeStub = nil
	if fake.handshakeReturnsOnCall == nil {
		fake.handshakeReturnsOnCall = make(map[int]struct {
			result1 *response.Empty
			result2 error
		})
	}
	fake.handshakeReturnsOnCall[i] = struct {
		result1 *response.Empty
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) Login(arg1 context.Context, arg2 *request.Login, arg3 ...grpc.CallOption) (*response.Message, error) {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		arg1 context.Context
		arg2 *request.Login
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("Login", []interface{}{arg1, arg2, arg3})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loginReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCoreClient) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeCoreClient) LoginCalls(stub func(context.Context, *request.Login, ...grpc.CallOption) (*response.Message, error)) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = stub
}

func (fake *FakeCoreClient) LoginArgsForCall(i int) (context.Context, *request.Login, []grpc.CallOption) {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	argsForCall := fake.loginArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCoreClient) LoginReturns(result1 *response.Message, result2 error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 *response.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) LoginReturnsOnCall(i int, result1 *response.Message, result2 error) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 *response.Message
			result2 error
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 *response.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) Verify(arg1 context.Context, arg2 *request.Verify, arg3 ...grpc.CallOption) (*response.Message, error) {
	fake.verifyMutex.Lock()
	ret, specificReturn := fake.verifyReturnsOnCall[len(fake.verifyArgsForCall)]
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		arg1 context.Context
		arg2 *request.Verify
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("Verify", []interface{}{arg1, arg2, arg3})
	fake.verifyMutex.Unlock()
	if fake.VerifyStub != nil {
		return fake.VerifyStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.verifyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCoreClient) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *FakeCoreClient) VerifyCalls(stub func(context.Context, *request.Verify, ...grpc.CallOption) (*response.Message, error)) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = stub
}

func (fake *FakeCoreClient) VerifyArgsForCall(i int) (context.Context, *request.Verify, []grpc.CallOption) {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	argsForCall := fake.verifyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCoreClient) VerifyReturns(result1 *response.Message, result2 error) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 *response.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) VerifyReturnsOnCall(i int, result1 *response.Message, result2 error) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = nil
	if fake.verifyReturnsOnCall == nil {
		fake.verifyReturnsOnCall = make(map[int]struct {
			result1 *response.Message
			result2 error
		})
	}
	fake.verifyReturnsOnCall[i] = struct {
		result1 *response.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createAccountMutex.RLock()
	defer fake.createAccountMutex.RUnlock()
	fake.createEventMutex.RLock()
	defer fake.createEventMutex.RUnlock()
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	fake.handshakeMutex.RLock()
	defer fake.handshakeMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCoreClient) recordInvocation(key string, args []interface{}) {
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

var _ pinpoint.CoreClient = new(FakeCoreClient)
