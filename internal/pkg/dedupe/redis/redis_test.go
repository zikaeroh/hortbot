package redis_test

import (
	"context"
	"testing"
	"time"

	redislib "github.com/hortbot/hortbot/internal/db/redis"
	"github.com/hortbot/hortbot/internal/pkg/dedupe/redis"
	"github.com/hortbot/hortbot/internal/pkg/testutil/miniredistest"
	"gotest.tools/v3/assert"
)

const id = "id"

func TestCheckNotFound(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, c, cleanup, err := miniredistest.New()
	assert.NilError(t, err)
	defer cleanup()

	rdb := redislib.New(c)

	d, err := redis.New(rdb, time.Second)
	assert.NilError(t, err)

	seen, err := d.Check(ctx, id)
	assert.Assert(t, !seen)
	assert.NilError(t, err)
}

func TestMarkThenCheck(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s, c, cleanup, err := miniredistest.New()
	assert.NilError(t, err)
	defer cleanup()

	rdb := redislib.New(c)

	d, err := redis.New(rdb, time.Minute)
	assert.NilError(t, err)

	assert.NilError(t, d.Mark(ctx, id))
	s.FastForward(time.Second)

	seen, err := d.Check(ctx, id)
	assert.Assert(t, seen)
	assert.NilError(t, err)
}

func TestMarkMarkThenCheck(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s, c, cleanup, err := miniredistest.New()
	assert.NilError(t, err)
	defer cleanup()

	rdb := redislib.New(c)

	d, err := redis.New(rdb, time.Minute)
	assert.NilError(t, err)

	assert.NilError(t, d.Mark(ctx, id))
	s.FastForward(time.Second)

	assert.NilError(t, d.Mark(ctx, id))
	s.FastForward(time.Second)

	seen, err := d.Check(ctx, id)
	assert.Assert(t, seen)
	assert.NilError(t, err)
}

func TestCheckAndMark(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s, c, cleanup, err := miniredistest.New()
	assert.NilError(t, err)
	defer cleanup()

	rdb := redislib.New(c)

	d, err := redis.New(rdb, time.Minute)
	assert.NilError(t, err)

	seen, err := d.CheckAndMark(ctx, id)
	assert.Assert(t, !seen)
	assert.NilError(t, err)

	s.FastForward(time.Second)

	seen, err = d.Check(ctx, id)
	assert.Assert(t, seen)
	assert.NilError(t, err)
}

func TestCheckAndMarkTwice(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s, c, cleanup, err := miniredistest.New()
	assert.NilError(t, err)
	defer cleanup()

	rdb := redislib.New(c)

	d, err := redis.New(rdb, time.Minute)
	assert.NilError(t, err)

	seen, err := d.CheckAndMark(ctx, id)
	assert.Assert(t, !seen)
	assert.NilError(t, err)

	s.FastForward(time.Second)

	seen, err = d.CheckAndMark(ctx, id)
	assert.Assert(t, seen)
	assert.NilError(t, err)
}

func TestExpire(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s, c, cleanup, err := miniredistest.New()
	assert.NilError(t, err)
	defer cleanup()

	rdb := redislib.New(c)

	d, err := redis.New(rdb, time.Second)
	assert.NilError(t, err)

	seen, err := d.CheckAndMark(ctx, id)
	assert.Assert(t, !seen)
	assert.NilError(t, err)

	s.FastForward(2 * time.Second)

	seen, err = d.Check(ctx, id)
	assert.Assert(t, !seen)
	assert.NilError(t, err)
}

func TestShortExpiry(t *testing.T) {
	t.Parallel()

	d, err := redis.New(nil, time.Millisecond)
	assert.Assert(t, d == nil)
	assert.Assert(t, err == redis.ErrExpiryTooShort)
}
