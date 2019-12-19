package bnsq_test

import (
	"context"
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
	"github.com/hortbot/hortbot/internal/bnsq"
	"github.com/hortbot/hortbot/internal/pkg/correlation"
	"github.com/hortbot/hortbot/internal/pkg/ctxlog"
	"github.com/hortbot/hortbot/internal/pkg/docker/dnsq"
	"github.com/hortbot/hortbot/internal/pkg/errgroupx"
	"github.com/hortbot/hortbot/internal/pkg/testutil"
	"github.com/rs/xid"
	"go.opencensus.io/trace"
	"gotest.tools/v3/assert"
)

func TestNotify(t *testing.T) {
	defer leaktest.Check(t)()

	addr, cleanup, err := dnsq.New()
	assert.NilError(t, err)
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger := testutil.Logger(t)
	ctx = ctxlog.WithLogger(ctx, logger)

	const (
		botName = "hortbot"
		channel = "blue"
	)

	received := make(chan *bnsq.ChannelUpdatesNotification, 10)
	cids := make(chan xid.ID, 10)
	spans := make(chan trace.SpanContext)

	publisher := bnsq.NewNotifyPublisher(addr)

	subscriber := bnsq.NotifySubscriber{
		Addr:    addr,
		BotName: botName,
		Channel: channel,
		OnNotifyChannelUpdates: func(n *bnsq.ChannelUpdatesNotification, m *bnsq.Metadata) error {
			received <- n
			cids <- correlation.FromContext(m.Correlate(context.Background()))
			spans <- m.ParentSpan()
			return nil
		},
	}

	g := errgroupx.FromContext(ctx)

	g.Go(publisher.Run)
	g.Go(subscriber.Run)

	id := xid.New()
	ctx, span := trace.StartSpan(ctx, "TwitchToken")

	assert.NilError(t, publisher.NotifyChannelUpdates(correlation.WithID(ctx, id), botName))
	assert.NilError(t, publisher.NotifyChannelUpdates(ctx, "wrong"))

	got := <-received
	gotID := <-cids
	gotSpan := <-spans

	g.Stop()

	assert.Equal(t, len(received), 0)

	assert.DeepEqual(t, got, &bnsq.ChannelUpdatesNotification{
		BotName: botName,
	})

	assert.Equal(t, gotID, id)
	assert.Equal(t, gotSpan.TraceID, span.SpanContext().TraceID)
}
