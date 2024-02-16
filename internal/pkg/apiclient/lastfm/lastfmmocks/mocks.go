// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package lastfmmocks

import (
	"context"
	"sync"

	"github.com/hortbot/hortbot/internal/pkg/apiclient/lastfm"
)

// Ensure, that APIMock does implement lastfm.API.
// If this is not the case, regenerate this file with moq.
var _ lastfm.API = &APIMock{}

// APIMock is a mock implementation of lastfm.API.
//
//	func TestSomethingThatUsesAPI(t *testing.T) {
//
//		// make and configure a mocked lastfm.API
//		mockedAPI := &APIMock{
//			RecentTracksFunc: func(ctx context.Context, user string, n int) ([]lastfm.Track, error) {
//				panic("mock out the RecentTracks method")
//			},
//		}
//
//		// use mockedAPI in code that requires lastfm.API
//		// and then make assertions.
//
//	}
type APIMock struct {
	// RecentTracksFunc mocks the RecentTracks method.
	RecentTracksFunc func(ctx context.Context, user string, n int) ([]lastfm.Track, error)

	// calls tracks calls to the methods.
	calls struct {
		// RecentTracks holds details about calls to the RecentTracks method.
		RecentTracks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// User is the user argument value.
			User string
			// N is the n argument value.
			N int
		}
	}
	lockRecentTracks sync.RWMutex
}

// RecentTracks calls RecentTracksFunc.
func (mock *APIMock) RecentTracks(ctx context.Context, user string, n int) ([]lastfm.Track, error) {
	if mock.RecentTracksFunc == nil {
		panic("APIMock.RecentTracksFunc: method is nil but API.RecentTracks was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		User string
		N    int
	}{
		Ctx:  ctx,
		User: user,
		N:    n,
	}
	mock.lockRecentTracks.Lock()
	mock.calls.RecentTracks = append(mock.calls.RecentTracks, callInfo)
	mock.lockRecentTracks.Unlock()
	return mock.RecentTracksFunc(ctx, user, n)
}

// RecentTracksCalls gets all the calls that were made to RecentTracks.
// Check the length with:
//
//	len(mockedAPI.RecentTracksCalls())
func (mock *APIMock) RecentTracksCalls() []struct {
	Ctx  context.Context
	User string
	N    int
} {
	var calls []struct {
		Ctx  context.Context
		User string
		N    int
	}
	mock.lockRecentTracks.RLock()
	calls = mock.calls.RecentTracks
	mock.lockRecentTracks.RUnlock()
	return calls
}
