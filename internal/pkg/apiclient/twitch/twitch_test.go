package twitch_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hortbot/hortbot/internal/pkg/apiclient/twitch"
	"github.com/hortbot/hortbot/internal/pkg/oauth2x"
	"golang.org/x/oauth2"
	"gotest.tools/v3/assert"
)

var (
	tokenCmp = cmp.Comparer(func(x, y oauth2.Token) bool {
		return oauth2x.Equals(&x, &y)
	})

	expectedErrors = map[int]error{
		401: twitch.ErrNotAuthorized,
		404: twitch.ErrNotFound,
		418: twitch.ErrUnknown,
		500: twitch.ErrServerError,
	}
)

const (
	clientID     = "client-id"
	clientSecret = "client-secret"
	redirectURL  = "http://localhost/auth/twitch/callback"
)

func TestNewPanic(t *testing.T) {
	checkPanic := func(fn func()) interface{} {
		var recovered interface{}

		func() {
			defer func() {
				recovered = recover()
			}()

			fn()
		}()

		return recovered
	}

	assert.Equal(t, checkPanic(func() {
		twitch.New(clientID, clientSecret, redirectURL)
	}), nil)

	assert.Equal(t, checkPanic(func() {
		twitch.New("", clientSecret, redirectURL)
	}), "empty clientID")

	assert.Equal(t, checkPanic(func() {
		twitch.New(clientID, "", redirectURL)
	}), "empty clientSecret")

	assert.Equal(t, checkPanic(func() {
		twitch.New(clientID, clientSecret, "")
	}), "empty redirectURL")
}

func TestAuthExchange(t *testing.T) {
	ctx, cancel := testContext(t)
	defer cancel()

	ft := newFakeTwitch(t)
	cli := ft.client()

	const state = "some-state"

	tw := twitch.New(clientID, clientSecret, redirectURL, twitch.HTTPClient(cli))

	assert.Equal(t,
		tw.AuthCodeURL(state, nil),
		fmt.Sprintf("https://id.twitch.tv/oauth2/authorize?access_type=offline&client_id=%s&force_verify=true&redirect_uri=%s&response_type=code&state=%s", clientID, url.QueryEscape(redirectURL), state),
	)

	code := ft.codeForUser(1234)

	tok, err := tw.Exchange(ctx, code)
	assert.NilError(t, err)
	assert.DeepEqual(t, tok, ft.tokenForCode(code), tokenCmp)

	assert.Equal(t,
		tw.AuthCodeURL(state, []string{"user_follows_edit"}),
		fmt.Sprintf("https://id.twitch.tv/oauth2/authorize?access_type=offline&client_id=%s&force_verify=true&redirect_uri=%s&response_type=code&scope=user_follows_edit&state=%s", clientID, url.QueryEscape(redirectURL), state),
	)
}
