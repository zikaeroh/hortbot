// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// ScheduledCommand is an object representing the database table.
type ScheduledCommand struct {
	ID             int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ChannelID      int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	CommandInfoID  int64     `boil:"command_info_id" json:"command_info_id" toml:"command_info_id" yaml:"command_info_id"`
	Enabled        bool      `boil:"enabled" json:"enabled" toml:"enabled" yaml:"enabled"`
	CronExpression string    `boil:"cron_expression" json:"cron_expression" toml:"cron_expression" yaml:"cron_expression"`
	MessageDiff    int64     `boil:"message_diff" json:"message_diff" toml:"message_diff" yaml:"message_diff"`
	LastCount      int64     `boil:"last_count" json:"last_count" toml:"last_count" yaml:"last_count"`
	Creator        string    `boil:"creator" json:"creator" toml:"creator" yaml:"creator"`
	Editor         string    `boil:"editor" json:"editor" toml:"editor" yaml:"editor"`

	R *scheduledCommandR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L scheduledCommandL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ScheduledCommandColumns = struct {
	ID             string
	CreatedAt      string
	UpdatedAt      string
	ChannelID      string
	CommandInfoID  string
	Enabled        string
	CronExpression string
	MessageDiff    string
	LastCount      string
	Creator        string
	Editor         string
}{
	ID:             "id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	ChannelID:      "channel_id",
	CommandInfoID:  "command_info_id",
	Enabled:        "enabled",
	CronExpression: "cron_expression",
	MessageDiff:    "message_diff",
	LastCount:      "last_count",
	Creator:        "creator",
	Editor:         "editor",
}

// Generated where

var ScheduledCommandWhere = struct {
	ID             whereHelperint64
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
	ChannelID      whereHelperint64
	CommandInfoID  whereHelperint64
	Enabled        whereHelperbool
	CronExpression whereHelperstring
	MessageDiff    whereHelperint64
	LastCount      whereHelperint64
	Creator        whereHelperstring
	Editor         whereHelperstring
}{
	ID:             whereHelperint64{field: "\"scheduled_commands\".\"id\""},
	CreatedAt:      whereHelpertime_Time{field: "\"scheduled_commands\".\"created_at\""},
	UpdatedAt:      whereHelpertime_Time{field: "\"scheduled_commands\".\"updated_at\""},
	ChannelID:      whereHelperint64{field: "\"scheduled_commands\".\"channel_id\""},
	CommandInfoID:  whereHelperint64{field: "\"scheduled_commands\".\"command_info_id\""},
	Enabled:        whereHelperbool{field: "\"scheduled_commands\".\"enabled\""},
	CronExpression: whereHelperstring{field: "\"scheduled_commands\".\"cron_expression\""},
	MessageDiff:    whereHelperint64{field: "\"scheduled_commands\".\"message_diff\""},
	LastCount:      whereHelperint64{field: "\"scheduled_commands\".\"last_count\""},
	Creator:        whereHelperstring{field: "\"scheduled_commands\".\"creator\""},
	Editor:         whereHelperstring{field: "\"scheduled_commands\".\"editor\""},
}

// ScheduledCommandRels is where relationship names are stored.
var ScheduledCommandRels = struct {
	Channel     string
	CommandInfo string
}{
	Channel:     "Channel",
	CommandInfo: "CommandInfo",
}

// scheduledCommandR is where relationships are stored.
type scheduledCommandR struct {
	Channel     *Channel
	CommandInfo *CommandInfo
}

// NewStruct creates a new relationship struct
func (*scheduledCommandR) NewStruct() *scheduledCommandR {
	return &scheduledCommandR{}
}

// scheduledCommandL is where Load methods for each relationship are stored.
type scheduledCommandL struct{}

var (
	scheduledCommandAllColumns            = []string{"id", "created_at", "updated_at", "channel_id", "command_info_id", "enabled", "cron_expression", "message_diff", "last_count", "creator", "editor"}
	scheduledCommandColumnsWithoutDefault = []string{"channel_id", "command_info_id", "enabled", "cron_expression", "last_count", "creator", "editor"}
	scheduledCommandColumnsWithDefault    = []string{"id", "created_at", "updated_at", "message_diff"}
	scheduledCommandPrimaryKeyColumns     = []string{"id"}
)

type (
	// ScheduledCommandSlice is an alias for a slice of pointers to ScheduledCommand.
	// This should generally be used opposed to []ScheduledCommand.
	ScheduledCommandSlice []*ScheduledCommand

	scheduledCommandQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	scheduledCommandType                 = reflect.TypeOf(&ScheduledCommand{})
	scheduledCommandMapping              = queries.MakeStructMapping(scheduledCommandType)
	scheduledCommandPrimaryKeyMapping, _ = queries.BindMapping(scheduledCommandType, scheduledCommandMapping, scheduledCommandPrimaryKeyColumns)
	scheduledCommandInsertCacheMut       sync.RWMutex
	scheduledCommandInsertCache          = make(map[string]insertCache)
	scheduledCommandUpdateCacheMut       sync.RWMutex
	scheduledCommandUpdateCache          = make(map[string]updateCache)
	scheduledCommandUpsertCacheMut       sync.RWMutex
	scheduledCommandUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single scheduledCommand record from the query.
func (q scheduledCommandQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ScheduledCommand, error) {
	o := &ScheduledCommand{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for scheduled_commands")
	}

	return o, nil
}

// All returns all ScheduledCommand records from the query.
func (q scheduledCommandQuery) All(ctx context.Context, exec boil.ContextExecutor) (ScheduledCommandSlice, error) {
	var o []*ScheduledCommand

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ScheduledCommand slice")
	}

	return o, nil
}

// Count returns the count of all ScheduledCommand records in the query.
func (q scheduledCommandQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count scheduled_commands rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q scheduledCommandQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if scheduled_commands exists")
	}

	return count > 0, nil
}

// Channel pointed to by the foreign key.
func (o *ScheduledCommand) Channel(mods ...qm.QueryMod) channelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := Channels(queryMods...)
	queries.SetFrom(query.Query, "\"channels\"")

	return query
}

// CommandInfo pointed to by the foreign key.
func (o *ScheduledCommand) CommandInfo(mods ...qm.QueryMod) commandInfoQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.CommandInfoID),
	}

	queryMods = append(queryMods, mods...)

	query := CommandInfos(queryMods...)
	queries.SetFrom(query.Query, "\"command_infos\"")

	return query
}

// LoadChannel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (scheduledCommandL) LoadChannel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeScheduledCommand interface{}, mods queries.Applicator) error {
	var slice []*ScheduledCommand
	var object *ScheduledCommand

	if singular {
		object = maybeScheduledCommand.(*ScheduledCommand)
	} else {
		slice = *maybeScheduledCommand.(*[]*ScheduledCommand)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &scheduledCommandR{}
		}
		args = append(args, object.ChannelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &scheduledCommandR{}
			}

			for _, a := range args {
				if a == obj.ChannelID {
					continue Outer
				}
			}

			args = append(args, obj.ChannelID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`channels`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Channel")
	}

	var resultSlice []*Channel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Channel")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for channels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for channels")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Channel = foreign
		if foreign.R == nil {
			foreign.R = &channelR{}
		}
		foreign.R.ScheduledCommands = append(foreign.R.ScheduledCommands, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChannelID == foreign.ID {
				local.R.Channel = foreign
				if foreign.R == nil {
					foreign.R = &channelR{}
				}
				foreign.R.ScheduledCommands = append(foreign.R.ScheduledCommands, local)
				break
			}
		}
	}

	return nil
}

// LoadCommandInfo allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (scheduledCommandL) LoadCommandInfo(ctx context.Context, e boil.ContextExecutor, singular bool, maybeScheduledCommand interface{}, mods queries.Applicator) error {
	var slice []*ScheduledCommand
	var object *ScheduledCommand

	if singular {
		object = maybeScheduledCommand.(*ScheduledCommand)
	} else {
		slice = *maybeScheduledCommand.(*[]*ScheduledCommand)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &scheduledCommandR{}
		}
		args = append(args, object.CommandInfoID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &scheduledCommandR{}
			}

			for _, a := range args {
				if a == obj.CommandInfoID {
					continue Outer
				}
			}

			args = append(args, obj.CommandInfoID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`command_infos`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CommandInfo")
	}

	var resultSlice []*CommandInfo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CommandInfo")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for command_infos")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for command_infos")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.CommandInfo = foreign
		if foreign.R == nil {
			foreign.R = &commandInfoR{}
		}
		foreign.R.ScheduledCommand = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CommandInfoID == foreign.ID {
				local.R.CommandInfo = foreign
				if foreign.R == nil {
					foreign.R = &commandInfoR{}
				}
				foreign.R.ScheduledCommand = local
				break
			}
		}
	}

	return nil
}

// SetChannel of the scheduledCommand to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.ScheduledCommands.
func (o *ScheduledCommand) SetChannel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Channel) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"scheduled_commands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
		strmangle.WhereClause("\"", "\"", 2, scheduledCommandPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ChannelID = related.ID
	if o.R == nil {
		o.R = &scheduledCommandR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &channelR{
			ScheduledCommands: ScheduledCommandSlice{o},
		}
	} else {
		related.R.ScheduledCommands = append(related.R.ScheduledCommands, o)
	}

	return nil
}

// SetCommandInfo of the scheduledCommand to the related item.
// Sets o.R.CommandInfo to related.
// Adds o to related.R.ScheduledCommand.
func (o *ScheduledCommand) SetCommandInfo(ctx context.Context, exec boil.ContextExecutor, insert bool, related *CommandInfo) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"scheduled_commands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"command_info_id"}),
		strmangle.WhereClause("\"", "\"", 2, scheduledCommandPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CommandInfoID = related.ID
	if o.R == nil {
		o.R = &scheduledCommandR{
			CommandInfo: related,
		}
	} else {
		o.R.CommandInfo = related
	}

	if related.R == nil {
		related.R = &commandInfoR{
			ScheduledCommand: o,
		}
	} else {
		related.R.ScheduledCommand = o
	}

	return nil
}

// ScheduledCommands retrieves all the records using an executor.
func ScheduledCommands(mods ...qm.QueryMod) scheduledCommandQuery {
	mods = append(mods, qm.From("\"scheduled_commands\""))
	return scheduledCommandQuery{NewQuery(mods...)}
}

// FindScheduledCommand retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindScheduledCommand(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ScheduledCommand, error) {
	scheduledCommandObj := &ScheduledCommand{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"scheduled_commands\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, scheduledCommandObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from scheduled_commands")
	}

	return scheduledCommandObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ScheduledCommand) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no scheduled_commands provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(scheduledCommandColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	scheduledCommandInsertCacheMut.RLock()
	cache, cached := scheduledCommandInsertCache[key]
	scheduledCommandInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			scheduledCommandAllColumns,
			scheduledCommandColumnsWithDefault,
			scheduledCommandColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(scheduledCommandType, scheduledCommandMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(scheduledCommandType, scheduledCommandMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"scheduled_commands\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"scheduled_commands\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into scheduled_commands")
	}

	if !cached {
		scheduledCommandInsertCacheMut.Lock()
		scheduledCommandInsertCache[key] = cache
		scheduledCommandInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the ScheduledCommand.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ScheduledCommand) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	scheduledCommandUpdateCacheMut.RLock()
	cache, cached := scheduledCommandUpdateCache[key]
	scheduledCommandUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			scheduledCommandAllColumns,
			scheduledCommandPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update scheduled_commands, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"scheduled_commands\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, scheduledCommandPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(scheduledCommandType, scheduledCommandMapping, append(wl, scheduledCommandPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update scheduled_commands row")
	}

	if !cached {
		scheduledCommandUpdateCacheMut.Lock()
		scheduledCommandUpdateCache[key] = cache
		scheduledCommandUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q scheduledCommandQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for scheduled_commands")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ScheduledCommandSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scheduledCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"scheduled_commands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, scheduledCommandPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in scheduledCommand slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ScheduledCommand) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no scheduled_commands provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(scheduledCommandColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	scheduledCommandUpsertCacheMut.RLock()
	cache, cached := scheduledCommandUpsertCache[key]
	scheduledCommandUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			scheduledCommandAllColumns,
			scheduledCommandColumnsWithDefault,
			scheduledCommandColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			scheduledCommandAllColumns,
			scheduledCommandPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert scheduled_commands, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(scheduledCommandPrimaryKeyColumns))
			copy(conflict, scheduledCommandPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"scheduled_commands\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(scheduledCommandType, scheduledCommandMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(scheduledCommandType, scheduledCommandMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert scheduled_commands")
	}

	if !cached {
		scheduledCommandUpsertCacheMut.Lock()
		scheduledCommandUpsertCache[key] = cache
		scheduledCommandUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single ScheduledCommand record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ScheduledCommand) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no ScheduledCommand provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), scheduledCommandPrimaryKeyMapping)
	sql := "DELETE FROM \"scheduled_commands\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from scheduled_commands")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q scheduledCommandQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no scheduledCommandQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from scheduled_commands")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ScheduledCommandSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scheduledCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"scheduled_commands\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, scheduledCommandPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from scheduledCommand slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ScheduledCommand) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindScheduledCommand(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ScheduledCommandSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ScheduledCommandSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scheduledCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"scheduled_commands\".* FROM \"scheduled_commands\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, scheduledCommandPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ScheduledCommandSlice")
	}

	*o = slice

	return nil
}

// ScheduledCommandExists checks if the ScheduledCommand row exists.
func ScheduledCommandExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"scheduled_commands\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if scheduled_commands exists")
	}

	return exists, nil
}
