// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime/cache (interfaces: KeySerializer)

package mocks

import (
	go_version "github.com/hashicorp/go-version"
	pegomock "github.com/petergtz/pegomock/v4"
	"reflect"
	"time"
)

type MockKeySerializer struct {
	fail func(message string, callerSkip ...int)
}

func NewMockKeySerializer(options ...pegomock.Option) *MockKeySerializer {
	mock := &MockKeySerializer{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockKeySerializer) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockKeySerializer) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockKeySerializer) Serialize(key *go_version.Version) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockKeySerializer().")
	}
	_params := []pegomock.Param{key}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Serialize", _params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var _ret0 string
	var _ret1 error
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(string)
		}
		if _result[1] != nil {
			_ret1 = _result[1].(error)
		}
	}
	return _ret0, _ret1
}

func (mock *MockKeySerializer) VerifyWasCalledOnce() *VerifierMockKeySerializer {
	return &VerifierMockKeySerializer{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockKeySerializer) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockKeySerializer {
	return &VerifierMockKeySerializer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockKeySerializer) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockKeySerializer {
	return &VerifierMockKeySerializer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockKeySerializer) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockKeySerializer {
	return &VerifierMockKeySerializer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockKeySerializer struct {
	mock                   *MockKeySerializer
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockKeySerializer) Serialize(key *go_version.Version) *MockKeySerializer_Serialize_OngoingVerification {
	_params := []pegomock.Param{key}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Serialize", _params, verifier.timeout)
	return &MockKeySerializer_Serialize_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockKeySerializer_Serialize_OngoingVerification struct {
	mock              *MockKeySerializer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockKeySerializer_Serialize_OngoingVerification) GetCapturedArguments() *go_version.Version {
	key := c.GetAllCapturedArguments()
	return key[len(key)-1]
}

func (c *MockKeySerializer_Serialize_OngoingVerification) GetAllCapturedArguments() (_param0 []*go_version.Version) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]*go_version.Version, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(*go_version.Version)
			}
		}
	}
	return
}
