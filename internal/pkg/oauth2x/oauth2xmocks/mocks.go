// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package oauth2xmocks

import (
	"sync"

	"github.com/hortbot/hortbot/internal/pkg/oauth2x"
	"golang.org/x/oauth2"
)

// Ensure, that TokenSourceMock does implement oauth2x.TokenSource.
// If this is not the case, regenerate this file with moq.
var _ oauth2x.TokenSource = &TokenSourceMock{}

// TokenSourceMock is a mock implementation of oauth2x.TokenSource.
//
//	func TestSomethingThatUsesTokenSource(t *testing.T) {
//
//		// make and configure a mocked oauth2x.TokenSource
//		mockedTokenSource := &TokenSourceMock{
//			TokenFunc: func() (*oauth2.Token, error) {
//				panic("mock out the Token method")
//			},
//		}
//
//		// use mockedTokenSource in code that requires oauth2x.TokenSource
//		// and then make assertions.
//
//	}
type TokenSourceMock struct {
	// TokenFunc mocks the Token method.
	TokenFunc func() (*oauth2.Token, error)

	// calls tracks calls to the methods.
	calls struct {
		// Token holds details about calls to the Token method.
		Token []struct {
		}
	}
	lockToken sync.RWMutex
}

// Token calls TokenFunc.
func (mock *TokenSourceMock) Token() (*oauth2.Token, error) {
	if mock.TokenFunc == nil {
		panic("TokenSourceMock.TokenFunc: method is nil but TokenSource.Token was just called")
	}
	callInfo := struct {
	}{}
	mock.lockToken.Lock()
	mock.calls.Token = append(mock.calls.Token, callInfo)
	mock.lockToken.Unlock()
	return mock.TokenFunc()
}

// TokenCalls gets all the calls that were made to Token.
// Check the length with:
//
//	len(mockedTokenSource.TokenCalls())
func (mock *TokenSourceMock) TokenCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockToken.RLock()
	calls = mock.calls.Token
	mock.lockToken.RUnlock()
	return calls
}
