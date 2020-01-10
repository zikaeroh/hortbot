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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Variable is an object representing the database table.
type Variable struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ChannelID int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Value     string    `boil:"value" json:"value" toml:"value" yaml:"value"`

	R *variableR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L variableL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VariableColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	ChannelID string
	Name      string
	Value     string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	ChannelID: "channel_id",
	Name:      "name",
	Value:     "value",
}

// Generated where

var VariableWhere = struct {
	ID        whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	ChannelID whereHelperint64
	Name      whereHelperstring
	Value     whereHelperstring
}{
	ID:        whereHelperint64{field: "\"variables\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"variables\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"variables\".\"updated_at\""},
	ChannelID: whereHelperint64{field: "\"variables\".\"channel_id\""},
	Name:      whereHelperstring{field: "\"variables\".\"name\""},
	Value:     whereHelperstring{field: "\"variables\".\"value\""},
}

// VariableRels is where relationship names are stored.
var VariableRels = struct {
	Channel string
}{
	Channel: "Channel",
}

// variableR is where relationships are stored.
type variableR struct {
	Channel *Channel
}

// NewStruct creates a new relationship struct
func (*variableR) NewStruct() *variableR {
	return &variableR{}
}

// variableL is where Load methods for each relationship are stored.
type variableL struct{}

var (
	variableAllColumns            = []string{"id", "created_at", "updated_at", "channel_id", "name", "value"}
	variableColumnsWithoutDefault = []string{"channel_id", "name", "value"}
	variableColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	variablePrimaryKeyColumns     = []string{"id"}
)

type (
	// VariableSlice is an alias for a slice of pointers to Variable.
	// This should generally be used opposed to []Variable.
	VariableSlice []*Variable

	variableQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	variableType                 = reflect.TypeOf(&Variable{})
	variableMapping              = queries.MakeStructMapping(variableType)
	variablePrimaryKeyMapping, _ = queries.BindMapping(variableType, variableMapping, variablePrimaryKeyColumns)
	variableInsertCacheMut       sync.RWMutex
	variableInsertCache          = make(map[string]insertCache)
	variableUpdateCacheMut       sync.RWMutex
	variableUpdateCache          = make(map[string]updateCache)
	variableUpsertCacheMut       sync.RWMutex
	variableUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single variable record from the query.
func (q variableQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Variable, error) {
	o := &Variable{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for variables")
	}

	return o, nil
}

// All returns all Variable records from the query.
func (q variableQuery) All(ctx context.Context, exec boil.ContextExecutor) (VariableSlice, error) {
	var o []*Variable

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Variable slice")
	}

	return o, nil
}

// Count returns the count of all Variable records in the query.
func (q variableQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count variables rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q variableQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if variables exists")
	}

	return count > 0, nil
}

// Channel pointed to by the foreign key.
func (o *Variable) Channel(mods ...qm.QueryMod) channelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := Channels(queryMods...)
	queries.SetFrom(query.Query, "\"channels\"")

	return query
}

// LoadChannel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (variableL) LoadChannel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeVariable interface{}, mods queries.Applicator) error {
	var slice []*Variable
	var object *Variable

	if singular {
		object = maybeVariable.(*Variable)
	} else {
		slice = *maybeVariable.(*[]*Variable)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &variableR{}
		}
		args = append(args, object.ChannelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &variableR{}
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

	query := NewQuery(qm.From(`channels`), qm.WhereIn(`channels.id in ?`, args...))
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
		foreign.R.Variables = append(foreign.R.Variables, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChannelID == foreign.ID {
				local.R.Channel = foreign
				if foreign.R == nil {
					foreign.R = &channelR{}
				}
				foreign.R.Variables = append(foreign.R.Variables, local)
				break
			}
		}
	}

	return nil
}

// SetChannel of the variable to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.Variables.
func (o *Variable) SetChannel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Channel) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"variables\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
		strmangle.WhereClause("\"", "\"", 2, variablePrimaryKeyColumns),
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
		o.R = &variableR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &channelR{
			Variables: VariableSlice{o},
		}
	} else {
		related.R.Variables = append(related.R.Variables, o)
	}

	return nil
}

// Variables retrieves all the records using an executor.
func Variables(mods ...qm.QueryMod) variableQuery {
	mods = append(mods, qm.From("\"variables\""))
	return variableQuery{NewQuery(mods...)}
}

// FindVariable retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVariable(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Variable, error) {
	variableObj := &Variable{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"variables\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, variableObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from variables")
	}

	return variableObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Variable) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no variables provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(variableColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	variableInsertCacheMut.RLock()
	cache, cached := variableInsertCache[key]
	variableInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			variableAllColumns,
			variableColumnsWithDefault,
			variableColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(variableType, variableMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(variableType, variableMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"variables\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"variables\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into variables")
	}

	if !cached {
		variableInsertCacheMut.Lock()
		variableInsertCache[key] = cache
		variableInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Variable.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Variable) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	variableUpdateCacheMut.RLock()
	cache, cached := variableUpdateCache[key]
	variableUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			variableAllColumns,
			variablePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update variables, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"variables\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, variablePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(variableType, variableMapping, append(wl, variablePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update variables row")
	}

	if !cached {
		variableUpdateCacheMut.Lock()
		variableUpdateCache[key] = cache
		variableUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q variableQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for variables")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VariableSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), variablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"variables\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, variablePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in variable slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Variable) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no variables provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(variableColumnsWithDefault, o)

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

	variableUpsertCacheMut.RLock()
	cache, cached := variableUpsertCache[key]
	variableUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			variableAllColumns,
			variableColumnsWithDefault,
			variableColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			variableAllColumns,
			variablePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert variables, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(variablePrimaryKeyColumns))
			copy(conflict, variablePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"variables\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(variableType, variableMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(variableType, variableMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert variables")
	}

	if !cached {
		variableUpsertCacheMut.Lock()
		variableUpsertCache[key] = cache
		variableUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Variable record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Variable) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no Variable provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), variablePrimaryKeyMapping)
	sql := "DELETE FROM \"variables\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from variables")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q variableQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no variableQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from variables")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VariableSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), variablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"variables\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, variablePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from variable slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Variable) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVariable(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VariableSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VariableSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), variablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"variables\".* FROM \"variables\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, variablePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VariableSlice")
	}

	*o = slice

	return nil
}

// VariableExists checks if the Variable row exists.
func VariableExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"variables\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if variables exists")
	}

	return exists, nil
}
