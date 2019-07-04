package repeat_test

import (
	"context"
	"testing"
	"time"

	"github.com/angadn/cronexpr"
	"github.com/fortytw2/leaktest"
	"github.com/hortbot/hortbot/internal/pkg/repeat"
	"github.com/leononame/clock"
	"gotest.tools/assert"
)

var startTime = mustParseTime("2000-10-01T03:11:00Z")

func TestCronAdd(t *testing.T) {
	defer leaktest.Check(t)()

	clk := clock.NewMock()
	clk.Set(startTime)

	runEveryHour := cronexpr.MustParse("0 0 * * * * *")

	count := 0
	fn := func(ctx context.Context, id int64) {
		count++
	}

	r := repeat.New(context.Background(), clk)

	r.AddCron(0, fn, runEveryHour)

	clk.Forward(time.Hour)
	clk.Forward(time.Hour)
	clk.Forward(time.Hour)
	clk.Forward(time.Hour)
	clk.Forward(time.Hour)

	time.Sleep(50 * time.Millisecond)

	r.Stop()

	assert.Equal(t, count, 5)
}

func TestCronAddTwice(t *testing.T) {
	defer leaktest.Check(t)()

	clk := clock.NewMock()
	clk.Set(startTime)

	runEveryHour := cronexpr.MustParse("0 0 * * * * *")

	count := 0
	fn := func(ctx context.Context, id int64) {
		count++
	}

	r := repeat.New(context.Background(), clk)

	r.AddCron(0, fn, runEveryHour)

	clk.Forward(time.Hour)
	clk.Forward(time.Hour)
	clk.Forward(time.Second)
	time.Sleep(50 * time.Millisecond)

	runEveryDay := cronexpr.MustParse("@daily")

	r.AddCron(0, fn, runEveryDay)
	time.Sleep(50 * time.Millisecond)

	clk.Forward(time.Hour)
	clk.Forward(time.Hour)
	clk.Forward(time.Hour)

	time.Sleep(50 * time.Millisecond)

	r.Stop()

	assert.Equal(t, count, 2)
}

func TestCronAddRemove(t *testing.T) {
	defer leaktest.Check(t)()

	clk := clock.NewMock()
	clk.Set(startTime)

	runEveryHour := cronexpr.MustParse("0 0 * * * * *")

	count := 0
	fn := func(ctx context.Context, id int64) {
		count++
	}

	r := repeat.New(context.Background(), clk)

	r.AddCron(0, fn, runEveryHour)

	clk.Forward(time.Hour)
	clk.Forward(time.Hour)
	clk.Forward(time.Second)
	time.Sleep(50 * time.Millisecond)

	r.RemoveCron(0)
	time.Sleep(50 * time.Millisecond)

	clk.Forward(time.Hour)
	clk.Forward(time.Hour)
	clk.Forward(time.Hour)

	time.Sleep(50 * time.Millisecond)

	r.Stop()

	assert.Equal(t, count, 2)
}

func mustParseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}
