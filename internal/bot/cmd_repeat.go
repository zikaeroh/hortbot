package bot

import (
	"context"
	"database/sql"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gobuffalo/flect"
	"github.com/hortbot/hortbot/internal/db/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var repeatCommands = newHandlerMap(map[string]handlerFunc{
	"add":    {fn: cmdRepeatAdd, minLevel: levelModerator},
	"delete": {fn: cmdRepeatDelete, minLevel: levelModerator},
	"remove": {fn: cmdRepeatDelete, minLevel: levelModerator},
	"on":     {fn: cmdRepeatOnOff, minLevel: levelModerator},
	"off":    {fn: cmdRepeatOnOff, minLevel: levelModerator},
	"list":   {fn: cmdRepeatList, minLevel: levelModerator},
})

func cmdRepeat(ctx context.Context, s *session, cmd string, args string) error {
	subcommand, args := splitSpace(args)
	subcommand = strings.ToLower(subcommand)

	ok, err := repeatCommands.run(ctx, s, subcommand, args)
	if err != nil {
		return err
	}

	if !ok {
		return s.ReplyUsage("add|delete|on|off|list ...")
	}

	return nil
}

func cmdRepeatAdd(ctx context.Context, s *session, cmd string, args string) error {
	usage := func() error {
		return s.ReplyUsage("<name> <delay in seconds> [message difference]")
	}

	name, args := splitSpace(args)
	delayStr, messageDiffStr := splitSpace(args)
	name = cleanCommandName(name)

	if name == "" {
		return usage()
	}

	delay, err := strconv.Atoi(delayStr)
	if err != nil {
		return usage()
	}

	if delay < 30 {
		return s.Reply("Delay must be at least 30 seconds.")
	}

	messageDiff := int64(1)

	if messageDiffStr != "" {
		messageDiff, err = strconv.ParseInt(messageDiffStr, 10, 64)
		if err != nil {
			return usage()
		}

		if messageDiff <= 0 {
			return s.Reply("Message difference must be at least 1.")
		}
	}

	info, repeat, err := findRepeatedCommand(ctx, name, s)
	if err != nil {
		return err
	}

	if info == nil {
		return s.Replyf("Command '%s' does not exist.", name)
	}

	if !s.UserLevel.CanAccess(newAccessLevel(info.AccessLevel)) {
		al := flect.Pluralize(info.AccessLevel)
		return s.Replyf("Command '%s' is restricted to %s; only %s and above can modify its repeat.", name, al, al)
	}

	if repeat != nil {
		repeat.Delay = delay
		repeat.MessageDiff = messageDiff
		repeat.Enabled = true
		repeat.LastCount = s.N
		repeat.Editor = s.User

		columns := boil.Whitelist(
			models.RepeatedCommandColumns.UpdatedAt,
			models.RepeatedCommandColumns.Delay,
			models.RepeatedCommandColumns.MessageDiff,
			models.RepeatedCommandColumns.Enabled,
			models.RepeatedCommandColumns.LastCount,
			models.RepeatedCommandColumns.Editor,
		)

		if err := repeat.Update(ctx, s.Tx, columns); err != nil {
			return err
		}
	} else {
		repeat = &models.RepeatedCommand{
			ChannelID:     s.Channel.ID,
			CommandInfoID: info.ID,
			Enabled:       true,
			Delay:         delay,
			MessageDiff:   messageDiff,
			LastCount:     s.N,
			Creator:       s.User,
			Editor:        s.User,
		}

		if err := repeat.Insert(ctx, s.Tx, boil.Infer()); err != nil {
			return err
		}
	}

	s.Deps.UpdateRepeat(repeat.ID, true, time.Duration(delay)*time.Second, 0)

	dUnit := "message has passed."
	if messageDiff != 1 {
		dUnit = "messages have passed."
	}

	return s.Replyf("Command '%s' will now repeat every %d seconds if at least %d %s", name, delay, messageDiff, dUnit)
}

func cmdRepeatDelete(ctx context.Context, s *session, cmd string, args string) error {
	name, _ := splitSpace(args)
	name = cleanCommandName(name)

	if name == "" {
		return s.ReplyUsage("<name>")
	}

	info, repeat, err := findRepeatedCommand(ctx, name, s)
	if err != nil {
		return err
	}

	if info == nil {
		return s.Replyf("Command '%s' does not exist.", name)
	}

	if repeat == nil {
		return s.Replyf("Command '%s' has no repeat.", name)
	}

	if !s.UserLevel.CanAccess(newAccessLevel(info.AccessLevel)) {
		al := flect.Pluralize(info.AccessLevel)
		return s.Replyf("Command '%s' is restricted to %s; only %s and above can modify its repeat.", name, al, al)
	}

	if err := repeat.Delete(ctx, s.Tx); err != nil {
		return err
	}

	s.Deps.UpdateRepeat(repeat.ID, false, 0, 0)

	return s.Replyf("Command '%s' will no longer repeat.", name)
}

func cmdRepeatOnOff(ctx context.Context, s *session, cmd string, args string) error {
	name, _ := splitSpace(args)
	name = cleanCommandName(name)

	if name == "" {
		return s.ReplyUsage("<name>")
	}

	enable := cmd == "on"

	info, repeat, err := findRepeatedCommand(ctx, name, s)
	if err != nil {
		return err
	}

	if info == nil {
		return s.Replyf("Command '%s' does not exist.", name)
	}

	if repeat == nil {
		return s.Replyf("Command '%s' has no repeat.", name)
	}

	if !s.UserLevel.CanAccess(newAccessLevel(info.AccessLevel)) {
		al := flect.Pluralize(info.AccessLevel)
		return s.Replyf("Command '%s' is restricted to %s; only %s and above can modify its repeat.", name, al, al)
	}

	if repeat.Enabled == enable {
		if enable {
			return s.Replyf("Repeated command '%s' is already enabled.", name)
		}
		return s.Replyf("Repeated command '%s' is already disabled.", name)
	}

	repeat.Enabled = enable
	repeat.LastCount = s.N
	repeat.Editor = s.User

	columns := boil.Whitelist(
		models.RepeatedCommandColumns.UpdatedAt,
		models.RepeatedCommandColumns.Enabled,
		models.RepeatedCommandColumns.LastCount,
		models.RepeatedCommandColumns.Editor,
	)

	if err := repeat.Update(ctx, s.Tx, columns); err != nil {
		return err
	}

	s.Deps.UpdateRepeat(repeat.ID, enable, time.Duration(repeat.Delay)*time.Second, 0)

	if enable {
		return s.Replyf("Repeated command '%s' is now enabled.", name)
	}

	return s.Replyf("Repeated command '%s' is now disabled.", name)
}

func cmdRepeatList(ctx context.Context, s *session, cmd string, args string) error {
	repeats, err := s.Channel.RepeatedCommands(
		qm.Load(models.RepeatedCommandRels.CommandInfo),
	).All(ctx, s.Tx)
	if err != nil {
		return err
	}

	if len(repeats) == 0 {
		return s.Reply("There are no repeated commands.")
	}

	sort.Slice(repeats, func(i, j int) bool {
		return repeats[i].R.CommandInfo.Name < repeats[j].R.CommandInfo.Name
	})

	var builder strings.Builder
	builder.WriteString("Repeated commands: ")

	for i, repeat := range repeats {
		if i != 0 {
			builder.WriteString(", ")
		}

		builder.WriteString(repeat.R.CommandInfo.Name)
		builder.WriteString(" [")

		if repeat.Enabled {
			builder.WriteString("ON")
		} else {
			builder.WriteString("OFF")
		}

		builder.WriteString("] (")
		builder.WriteString(strconv.Itoa(repeat.Delay))
		builder.WriteByte(')')
	}

	return s.Reply(builder.String())
}

func findRepeatedCommand(ctx context.Context, name string, s *session) (*models.CommandInfo, *models.RepeatedCommand, error) {
	info, err := s.Channel.CommandInfos(
		models.CommandInfoWhere.Name.EQ(name),
		qm.Load(models.CommandInfoRels.RepeatedCommand),
		qm.For("UPDATE"),
	).One(ctx, s.Tx)

	if err == sql.ErrNoRows {
		return nil, nil, nil
	}

	if err != nil {
		return nil, nil, err
	}

	return info, info.R.RepeatedCommand, nil
}
