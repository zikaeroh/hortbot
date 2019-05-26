package bot_test

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	"github.com/bmatcuk/doublestar"
	"github.com/gofrs/uuid"
	"github.com/hortbot/hortbot/internal/bot"
	"github.com/hortbot/hortbot/internal/bot/botfakes"
	"github.com/hortbot/hortbot/internal/ctxlog"
	"github.com/hortbot/hortbot/internal/db/models"
	"github.com/hortbot/hortbot/internal/dedupe"
	dedupemem "github.com/hortbot/hortbot/internal/dedupe/memory"
	"github.com/hortbot/hortbot/internal/testutil"
	"github.com/jakebailey/irc"
	"github.com/volatiletech/sqlboiler/boil"
	"gotest.tools/assert"
)

func TestScripts(t *testing.T) {
	files, err := doublestar.Glob(filepath.Join("testdata", "script", "**", "*.txt"))
	assert.NilError(t, err)
	assert.Assert(t, len(files) != 0)

	prefix := filepath.Join("testdata", "script")

	for _, file := range files {
		file := file
		name := strings.TrimSuffix(strings.TrimPrefix(file, prefix)[1:], ".txt")
		t.Run(name, func(t *testing.T) {
			testScriptFile(t, file)
		})
	}
}

func testScriptFile(t *testing.T, filename string) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic: %v", r)
			t.Logf("%s", debug.Stack())
			t.FailNow()
		}
	}()

	ctx := ctxlog.WithLogger(context.Background(), testutil.Logger(t))

	resetDatabase(t)

	f, err := os.Open(filename)
	assert.NilError(t, err)
	defer f.Close()

	sender := &botfakes.FakeMessageSender{}

	bc := bot.Config{
		DB:     db,
		Dedupe: dedupe.NeverSeen,
		Sender: sender,
	}

	var b *bot.Bot
	counts := make(map[string]int)
	var actions []func()
	var cleanup []func()

	ensureBot := func() {
		if b != nil {
			return
		}
		b = bot.New(&bc)
	}

	scanner := bufio.NewScanner(f)

	lineNum := 0
	sentBefore := 0

	for scanner.Scan() {
		lineNum++
		lineNum := lineNum // Shadow lineNum, otherwise the closures below will get the last line's number

		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" || line[0] == '#' {
			continue
		}

		directive := strings.SplitN(line, " ", 2)
		me := false

		switch directive[0] {
		case "skip":
			if len(directive) > 1 {
				reason := strings.TrimSpace(directive[1])
				t.Skipf("line %d: %s", lineNum, reason)
			} else {
				t.Skipf("line %d", lineNum)
			}

		case "boil_debug":
			oldMode := boil.DebugMode
			oldWriter := boil.DebugWriter

			boil.DebugMode = true
			boil.DebugWriter = testutil.Writer{T: t}

			defer func() {
				boil.DebugMode = oldMode
				boil.DebugWriter = oldWriter
			}()

		case "bot_config":
			actions = append(actions, func() {
				assert.Assert(t, b == nil, "bot has already been created, cannot configure")

				var bcj struct {
					Prefix string
					Bullet string
					Dedupe string
				}

				assert.NilError(t, json.Unmarshal([]byte(directive[1]), &bcj), "line %d", lineNum)

				bc.Prefix = bcj.Prefix
				bc.Bullet = bcj.Bullet

				switch bcj.Dedupe {
				case "", "never":
					bc.Dedupe = dedupe.NeverSeen

				case "memory":
					d := dedupemem.New(time.Minute, 5*time.Minute)
					cleanup = append(cleanup, d.Stop)
					bc.Dedupe = d

				default:
					t.Fatalf("line %d: unknown dedupe type %s", lineNum, bcj.Dedupe)
				}
			})

		case "insert_channel":
			var channel models.Channel
			assert.NilError(t, json.Unmarshal([]byte(directive[1]), &channel), "line %d", lineNum)

			actions = append(actions, func() {
				assert.NilError(t, channel.Insert(ctx, db, boil.Infer()), "line %d", lineNum)
			})

		case "insert_simple_command":
			var sc models.SimpleCommand
			assert.NilError(t, json.Unmarshal([]byte(directive[1]), &sc), "line %d", lineNum)

			actions = append(actions, func() {
				assert.NilError(t, sc.Insert(ctx, db, boil.Infer()), "line %d", lineNum)
			})

		case "handle_me":
			me = true
			fallthrough
		case "handle":
			args := strings.SplitN(directive[1], " ", 2)
			assert.Assert(t, len(args) == 2, "line %d", lineNum)

			origin := args[0]
			mRaw := args[1]

			var m *irc.Message

			if mRaw != "nil" {
				u := uuid.Must(uuid.NewV4())
				mRaw = strings.ReplaceAll(mRaw, "__UUID__", u.String())

				var err error
				m, err = irc.ParseMessage(mRaw)
				assert.NilError(t, err, "line %d", lineNum)

				if me {
					m.Trailing, _ = irc.EncodeCTCP("ACTION", m.Trailing)
				}
			}
			actions = append(actions, func() {
				ensureBot()
				sentBefore = sender.SendMessageCallCount()
				b.Handle(ctx, origin, m)
			})

		case "send":
			callNum := counts["send"]
			counts["send"]++

			sent := strings.SplitN(directive[1], " ", 3)
			assert.Assert(t, len(sent) == 3, "line %d", lineNum)

			actions = append(actions, func() {
				assert.Assert(t, sender.SendMessageCallCount() > callNum, "SendMessage not called: line %d", lineNum)
				origin, target, message := sender.SendMessageArgsForCall(callNum)
				assert.Equal(t, origin, sent[0], "line %d", lineNum)
				assert.Equal(t, target, sent[1], "line %d", lineNum)
				assert.Equal(t, message, sent[2], "line %d", lineNum)
			})

		case "no_send":
			actions = append(actions, func() {
				sentAfter := sender.SendMessageCallCount()

				if sentBefore != sentAfter {
					origin, target, message := sender.SendMessageArgsForCall(sentAfter - 1)
					t.Errorf("sent message: origin=%s, target=%s, message=%s: line %d", origin, target, message, lineNum)
					t.FailNow()
				}
			})

		default:
			t.Fatalf("line %d: unknown directive %s", lineNum, directive[0])
		}
	}

	assert.NilError(t, scanner.Err())

	if len(actions) == 0 {
		t.Error("no actions were provided")
		t.FailNow()
	}

	for _, fn := range actions {
		fn()
	}

	for i := len(cleanup) - 1; i >= 0; i-- {
		cleanup[i]()
	}

	assert.Equal(t, sender.SendMessageCallCount(), counts["send"])
}
