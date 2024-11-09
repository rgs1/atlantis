// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime (interfaces: ExternalTeamAllowlistRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockExternalTeamAllowlistRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockExternalTeamAllowlistRunner(options ...pegomock.Option) *MockExternalTeamAllowlistRunner {
	mock := &MockExternalTeamAllowlistRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockExternalTeamAllowlistRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockExternalTeamAllowlistRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockExternalTeamAllowlistRunner) Run(ctx models.TeamAllowlistCheckerContext, shell string, shellArgs string, command string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockExternalTeamAllowlistRunner().")
	}
	_params := []pegomock.Param{ctx, shell, shellArgs, command}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Run", _params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
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

func (mock *MockExternalTeamAllowlistRunner) VerifyWasCalledOnce() *VerifierMockExternalTeamAllowlistRunner {
	return &VerifierMockExternalTeamAllowlistRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockExternalTeamAllowlistRunner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockExternalTeamAllowlistRunner {
	return &VerifierMockExternalTeamAllowlistRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockExternalTeamAllowlistRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockExternalTeamAllowlistRunner {
	return &VerifierMockExternalTeamAllowlistRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockExternalTeamAllowlistRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockExternalTeamAllowlistRunner {
	return &VerifierMockExternalTeamAllowlistRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockExternalTeamAllowlistRunner struct {
	mock                   *MockExternalTeamAllowlistRunner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockExternalTeamAllowlistRunner) Run(ctx models.TeamAllowlistCheckerContext, shell string, shellArgs string, command string) *MockExternalTeamAllowlistRunner_Run_OngoingVerification {
	_params := []pegomock.Param{ctx, shell, shellArgs, command}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", _params, verifier.timeout)
	return &MockExternalTeamAllowlistRunner_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockExternalTeamAllowlistRunner_Run_OngoingVerification struct {
	mock              *MockExternalTeamAllowlistRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockExternalTeamAllowlistRunner_Run_OngoingVerification) GetCapturedArguments() (models.TeamAllowlistCheckerContext, string, string, string) {
	ctx, shell, shellArgs, command := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], shell[len(shell)-1], shellArgs[len(shellArgs)-1], command[len(command)-1]
}

func (c *MockExternalTeamAllowlistRunner_Run_OngoingVerification) GetAllCapturedArguments() (_param0 []models.TeamAllowlistCheckerContext, _param1 []string, _param2 []string, _param3 []string) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]models.TeamAllowlistCheckerContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(models.TeamAllowlistCheckerContext)
			}
		}
		if len(_params) > 1 {
			_param1 = make([]string, len(c.methodInvocations))
			for u, param := range _params[1] {
				_param1[u] = param.(string)
			}
		}
		if len(_params) > 2 {
			_param2 = make([]string, len(c.methodInvocations))
			for u, param := range _params[2] {
				_param2[u] = param.(string)
			}
		}
		if len(_params) > 3 {
			_param3 = make([]string, len(c.methodInvocations))
			for u, param := range _params[3] {
				_param3[u] = param.(string)
			}
		}
	}
	return
}
