package xkcd_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/hortbot/hortbot/internal/pkg/apiclient/xkcd"
	"github.com/hortbot/hortbot/internal/pkg/httpmockx"
	"github.com/jarcoal/httpmock"
	"gotest.tools/v3/assert"
)

const comic1 = `{
    "month": "1",
    "num": 1,
    "link": "",
    "year": "2006",
    "news": "",
    "safe_title": "Barrel - Part 1",
    "transcript": "[[A boy sits in a barrel which is floating in an ocean.]]\nBoy: I wonder where I'll float next?\n[[The barrel drifts into the distance. Nothing else can be seen.]]\n{{Alt: Don't we all.}}",
    "alt": "Don't we all.",
    "img": "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
    "title": "Barrel - Part 1",
    "day": "1"
}`

func TestGetComic(t *testing.T) {
	mt := httpmockx.NewMockTransport(t)

	errTest := errors.New("test error")

	mt.RegisterResponder(
		"GET",
		"https://xkcd.com/1/info.0.json",
		httpmock.NewStringResponder(200, comic1),
	)

	mt.RegisterResponder(
		"GET",
		"https://xkcd.com/77777/info.0.json",
		httpmock.NewStringResponder(404, ``),
	)

	mt.RegisterResponder(
		"GET",
		"https://xkcd.com/88888/info.0.json",
		httpmock.NewStringResponder(200, `{`),
	)

	mt.RegisterResponder(
		"GET",
		"https://xkcd.com/99999/info.0.json",
		func(_ *http.Request) (*http.Response, error) {
			return nil, errTest
		},
	)

	t.Run("OK", func(t *testing.T) {
		x := xkcd.New(&http.Client{Transport: mt})

		comic, err := x.GetComic(context.Background(), 1)
		assert.NilError(t, err)
		assert.DeepEqual(t, comic, &xkcd.Comic{
			Title: "Barrel - Part 1",
			Img:   "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
			Alt:   "Don't we all.",
		})
	})

	t.Run("Not found", func(t *testing.T) {
		x := xkcd.New(&http.Client{Transport: mt})

		_, err := x.GetComic(context.Background(), 77777)
		assert.Equal(t, err, xkcd.ErrNotFound)
	})

	t.Run("Client error", func(t *testing.T) {
		x := xkcd.New(&http.Client{Transport: mt})

		_, err := x.GetComic(context.Background(), 99999)
		assert.ErrorContains(t, err, errTest.Error())
	})

	t.Run("Decode error", func(t *testing.T) {
		x := xkcd.New(&http.Client{Transport: mt})

		_, err := x.GetComic(context.Background(), 88888)
		assert.ErrorContains(t, err, "unexpected EOF")
	})
}
