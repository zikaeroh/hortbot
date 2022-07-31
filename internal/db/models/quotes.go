// Code generated by SQLBoiler v4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Quote is an object representing the database table.
type Quote struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ChannelID int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	Num       int       `boil:"num" json:"num" toml:"num" yaml:"num"`
	Quote     string    `boil:"quote" json:"quote" toml:"quote" yaml:"quote"`
	Creator   string    `boil:"creator" json:"creator" toml:"creator" yaml:"creator"`
	Editor    string    `boil:"editor" json:"editor" toml:"editor" yaml:"editor"`

	R *quoteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L quoteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QuoteColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	ChannelID string
	Num       string
	Quote     string
	Creator   string
	Editor    string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	ChannelID: "channel_id",
	Num:       "num",
	Quote:     "quote",
	Creator:   "creator",
	Editor:    "editor",
}

var QuoteTableColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	ChannelID string
	Num       string
	Quote     string
	Creator   string
	Editor    string
}{
	ID:        "quotes.id",
	CreatedAt: "quotes.created_at",
	UpdatedAt: "quotes.updated_at",
	ChannelID: "quotes.channel_id",
	Num:       "quotes.num",
	Quote:     "quotes.quote",
	Creator:   "quotes.creator",
	Editor:    "quotes.editor",
}

// Generated where

var QuoteWhere = struct {
	ID        whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	ChannelID whereHelperint64
	Num       whereHelperint
	Quote     whereHelperstring
	Creator   whereHelperstring
	Editor    whereHelperstring
}{
	ID:        whereHelperint64{field: "\"quotes\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"quotes\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"quotes\".\"updated_at\""},
	ChannelID: whereHelperint64{field: "\"quotes\".\"channel_id\""},
	Num:       whereHelperint{field: "\"quotes\".\"num\""},
	Quote:     whereHelperstring{field: "\"quotes\".\"quote\""},
	Creator:   whereHelperstring{field: "\"quotes\".\"creator\""},
	Editor:    whereHelperstring{field: "\"quotes\".\"editor\""},
}

// QuoteRels is where relationship names are stored.
var QuoteRels = struct {
	Channel string
}{
	Channel: "Channel",
}

// quoteR is where relationships are stored.
type quoteR struct {
	Channel *Channel `boil:"Channel" json:"Channel" toml:"Channel" yaml:"Channel"`
}

// NewStruct creates a new relationship struct
func (*quoteR) NewStruct() *quoteR {
	return &quoteR{}
}

func (r *quoteR) GetChannel() *Channel {
	if r == nil {
		return nil
	}
	return r.Channel
}

// quoteL is where Load methods for each relationship are stored.
type quoteL struct{}

var (
	quoteAllColumns            = []string{"id", "created_at", "updated_at", "channel_id", "num", "quote", "creator", "editor"}
	quoteColumnsWithoutDefault = []string{"channel_id", "num", "quote", "creator", "editor"}
	quoteColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	quotePrimaryKeyColumns     = []string{"id"}
	quoteGeneratedColumns      = []string{}
)

type (
	// QuoteSlice is an alias for a slice of pointers to Quote.
	// This should almost always be used instead of []Quote.
	QuoteSlice []*Quote

	quoteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	quoteType                 = reflect.TypeOf(&Quote{})
	quoteMapping              = queries.MakeStructMapping(quoteType)
	quotePrimaryKeyMapping, _ = queries.BindMapping(quoteType, quoteMapping, quotePrimaryKeyColumns)
	quoteInsertCacheMut       sync.RWMutex
	quoteInsertCache          = make(map[string]insertCache)
	quoteUpdateCacheMut       sync.RWMutex
	quoteUpdateCache          = make(map[string]updateCache)
	quoteUpsertCacheMut       sync.RWMutex
	quoteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single quote record from the query.
func (q quoteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Quote, error) {
	o := &Quote{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for quotes")
	}

	return o, nil
}

// All returns all Quote records from the query.
func (q quoteQuery) All(ctx context.Context, exec boil.ContextExecutor) (QuoteSlice, error) {
	var o []*Quote

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Quote slice")
	}

	return o, nil
}

// Count returns the count of all Quote records in the query.
func (q quoteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count quotes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q quoteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if quotes exists")
	}

	return count > 0, nil
}

// Channel pointed to by the foreign key.
func (o *Quote) Channel(mods ...qm.QueryMod) channelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	return Channels(queryMods...)
}

// LoadChannel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (quoteL) LoadChannel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQuote interface{}, mods queries.Applicator) error {
	var slice []*Quote
	var object *Quote

	if singular {
		var ok bool
		object, ok = maybeQuote.(*Quote)
		if !ok {
			object = new(Quote)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeQuote)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeQuote))
			}
		}
	} else {
		s, ok := maybeQuote.(*[]*Quote)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeQuote)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeQuote))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &quoteR{}
		}
		args = append(args, object.ChannelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &quoteR{}
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

	query := NewQuery(
		qm.From(`channels`),
		qm.WhereIn(`channels.id in ?`, args...),
	)
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
		foreign.R.Quotes = append(foreign.R.Quotes, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChannelID == foreign.ID {
				local.R.Channel = foreign
				if foreign.R == nil {
					foreign.R = &channelR{}
				}
				foreign.R.Quotes = append(foreign.R.Quotes, local)
				break
			}
		}
	}

	return nil
}

// SetChannel of the quote to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.Quotes.
func (o *Quote) SetChannel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Channel) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"quotes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
		strmangle.WhereClause("\"", "\"", 2, quotePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ChannelID = related.ID
	if o.R == nil {
		o.R = &quoteR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &channelR{
			Quotes: QuoteSlice{o},
		}
	} else {
		related.R.Quotes = append(related.R.Quotes, o)
	}

	return nil
}

// Quotes retrieves all the records using an executor.
func Quotes(mods ...qm.QueryMod) quoteQuery {
	mods = append(mods, qm.From("\"quotes\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"quotes\".*"})
	}

	return quoteQuery{q}
}

// FindQuote retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQuote(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Quote, error) {
	quoteObj := &Quote{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"quotes\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, quoteObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from quotes")
	}

	return quoteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Quote) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no quotes provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(quoteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	quoteInsertCacheMut.RLock()
	cache, cached := quoteInsertCache[key]
	quoteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			quoteAllColumns,
			quoteColumnsWithDefault,
			quoteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(quoteType, quoteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(quoteType, quoteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"quotes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"quotes\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into quotes")
	}

	if !cached {
		quoteInsertCacheMut.Lock()
		quoteInsertCache[key] = cache
		quoteInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Quote.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Quote) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	quoteUpdateCacheMut.RLock()
	cache, cached := quoteUpdateCache[key]
	quoteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			quoteAllColumns,
			quotePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update quotes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"quotes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, quotePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(quoteType, quoteMapping, append(wl, quotePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update quotes row")
	}

	if !cached {
		quoteUpdateCacheMut.Lock()
		quoteUpdateCache[key] = cache
		quoteUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q quoteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for quotes")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o QuoteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), quotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"quotes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, quotePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in quote slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Quote) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no quotes provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(quoteColumnsWithDefault, o)

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

	quoteUpsertCacheMut.RLock()
	cache, cached := quoteUpsertCache[key]
	quoteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			quoteAllColumns,
			quoteColumnsWithDefault,
			quoteColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			quoteAllColumns,
			quotePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert quotes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(quotePrimaryKeyColumns))
			copy(conflict, quotePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"quotes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(quoteType, quoteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(quoteType, quoteMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert quotes")
	}

	if !cached {
		quoteUpsertCacheMut.Lock()
		quoteUpsertCache[key] = cache
		quoteUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Quote record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Quote) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no Quote provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), quotePrimaryKeyMapping)
	sql := "DELETE FROM \"quotes\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from quotes")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q quoteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no quoteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from quotes")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o QuoteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), quotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"quotes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, quotePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from quote slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Quote) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindQuote(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QuoteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QuoteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), quotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"quotes\".* FROM \"quotes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, quotePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in QuoteSlice")
	}

	*o = slice

	return nil
}

// QuoteExists checks if the Quote row exists.
func QuoteExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"quotes\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if quotes exists")
	}

	return exists, nil
}
