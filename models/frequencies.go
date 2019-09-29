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
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// Frequency is an object representing the database table.
type Frequency struct {
	ID        int64         `boil:"id" json:"id" toml:"id" yaml:"id"`
	SoundID   int64         `boil:"sound_id" json:"soundID" toml:"soundID" yaml:"soundID"`
	Frequency int           `boil:"frequency" json:"frequency" toml:"frequency" yaml:"frequency"`
	SPL       types.Decimal `boil:"spl" json:"spl" toml:"spl" yaml:"spl"`
	R         *frequencyR   `boil:"-" json:"-" toml:"-" yaml:"-"`
	L         frequencyL    `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FrequencyColumns = struct {
	ID        string
	SoundID   string
	Frequency string
	SPL       string
}{
	ID:        "id",
	SoundID:   "sound_id",
	Frequency: "frequency",
	SPL:       "spl",
}

// FrequencyRels is where relationship names are stored.
var FrequencyRels = struct {
	Sound string
}{
	Sound: "Sound",
}

// frequencyR is where relationships are stored.
type frequencyR struct {
	Sound *Sound
}

// NewStruct creates a new relationship struct
func (*frequencyR) NewStruct() *frequencyR {
	return &frequencyR{}
}

// frequencyL is where Load methods for each relationship are stored.
type frequencyL struct{}

var (
	frequencyColumns               = []string{"id", "sound_id", "frequency", "spl"}
	frequencyColumnsWithoutDefault = []string{"sound_id", "frequency", "spl"}
	frequencyColumnsWithDefault    = []string{"id"}
	frequencyPrimaryKeyColumns     = []string{"id"}
)

type (
	// FrequencySlice is an alias for a slice of pointers to Frequency.
	// This should generally be used opposed to []Frequency.
	FrequencySlice []*Frequency
	// FrequencyHook is the signature for custom Frequency hook methods
	FrequencyHook func(context.Context, boil.ContextExecutor, *Frequency) error

	frequencyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	frequencyType                 = reflect.TypeOf(&Frequency{})
	frequencyMapping              = queries.MakeStructMapping(frequencyType)
	frequencyPrimaryKeyMapping, _ = queries.BindMapping(frequencyType, frequencyMapping, frequencyPrimaryKeyColumns)
	frequencyInsertCacheMut       sync.RWMutex
	frequencyInsertCache          = make(map[string]insertCache)
	frequencyUpdateCacheMut       sync.RWMutex
	frequencyUpdateCache          = make(map[string]updateCache)
	frequencyUpsertCacheMut       sync.RWMutex
	frequencyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var frequencyBeforeInsertHooks []FrequencyHook
var frequencyBeforeUpdateHooks []FrequencyHook
var frequencyBeforeDeleteHooks []FrequencyHook
var frequencyBeforeUpsertHooks []FrequencyHook

var frequencyAfterInsertHooks []FrequencyHook
var frequencyAfterSelectHooks []FrequencyHook
var frequencyAfterUpdateHooks []FrequencyHook
var frequencyAfterDeleteHooks []FrequencyHook
var frequencyAfterUpsertHooks []FrequencyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Frequency) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range frequencyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Frequency) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range frequencyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Frequency) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range frequencyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Frequency) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range frequencyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Frequency) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range frequencyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Frequency) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range frequencyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Frequency) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range frequencyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Frequency) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range frequencyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Frequency) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range frequencyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFrequencyHook registers your hook function for all future operations.
func AddFrequencyHook(hookPoint boil.HookPoint, frequencyHook FrequencyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		frequencyBeforeInsertHooks = append(frequencyBeforeInsertHooks, frequencyHook)
	case boil.BeforeUpdateHook:
		frequencyBeforeUpdateHooks = append(frequencyBeforeUpdateHooks, frequencyHook)
	case boil.BeforeDeleteHook:
		frequencyBeforeDeleteHooks = append(frequencyBeforeDeleteHooks, frequencyHook)
	case boil.BeforeUpsertHook:
		frequencyBeforeUpsertHooks = append(frequencyBeforeUpsertHooks, frequencyHook)
	case boil.AfterInsertHook:
		frequencyAfterInsertHooks = append(frequencyAfterInsertHooks, frequencyHook)
	case boil.AfterSelectHook:
		frequencyAfterSelectHooks = append(frequencyAfterSelectHooks, frequencyHook)
	case boil.AfterUpdateHook:
		frequencyAfterUpdateHooks = append(frequencyAfterUpdateHooks, frequencyHook)
	case boil.AfterDeleteHook:
		frequencyAfterDeleteHooks = append(frequencyAfterDeleteHooks, frequencyHook)
	case boil.AfterUpsertHook:
		frequencyAfterUpsertHooks = append(frequencyAfterUpsertHooks, frequencyHook)
	}
}

// One returns a single frequency record from the query.
func (q frequencyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Frequency, error) {
	o := &Frequency{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for frequencies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Frequency records from the query.
func (q frequencyQuery) All(ctx context.Context, exec boil.ContextExecutor) (FrequencySlice, error) {
	var o []*Frequency

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Frequency slice")
	}

	if len(frequencyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Frequency records in the query.
func (q frequencyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count frequencies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q frequencyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if frequencies exists")
	}

	return count > 0, nil
}

// Sound pointed to by the foreign key.
func (o *Frequency) Sound(mods ...qm.QueryMod) soundQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.SoundID),
	}

	queryMods = append(queryMods, mods...)

	query := Sounds(queryMods...)
	queries.SetFrom(query.Query, "\"sounds\"")

	return query
}

// LoadSound allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (frequencyL) LoadSound(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFrequency interface{}, mods queries.Applicator) error {
	var slice []*Frequency
	var object *Frequency

	if singular {
		object = maybeFrequency.(*Frequency)
	} else {
		slice = *maybeFrequency.(*[]*Frequency)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &frequencyR{}
		}
		args = append(args, object.SoundID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &frequencyR{}
			}

			for _, a := range args {
				if a == obj.SoundID {
					continue Outer
				}
			}

			args = append(args, obj.SoundID)

		}
	}

	query := NewQuery(qm.From(`sounds`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Sound")
	}

	var resultSlice []*Sound
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Sound")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for sounds")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sounds")
	}

	if len(frequencyAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Sound = foreign
		if foreign.R == nil {
			foreign.R = &soundR{}
		}
		foreign.R.Frequencies = append(foreign.R.Frequencies, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SoundID == foreign.ID {
				local.R.Sound = foreign
				if foreign.R == nil {
					foreign.R = &soundR{}
				}
				foreign.R.Frequencies = append(foreign.R.Frequencies, local)
				break
			}
		}
	}

	return nil
}

// SetSound of the frequency to the related item.
// Sets o.R.Sound to related.
// Adds o to related.R.Frequencies.
func (o *Frequency) SetSound(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Sound) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"frequencies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"sound_id"}),
		strmangle.WhereClause("\"", "\"", 2, frequencyPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SoundID = related.ID
	if o.R == nil {
		o.R = &frequencyR{
			Sound: related,
		}
	} else {
		o.R.Sound = related
	}

	if related.R == nil {
		related.R = &soundR{
			Frequencies: FrequencySlice{o},
		}
	} else {
		related.R.Frequencies = append(related.R.Frequencies, o)
	}

	return nil
}

// Frequencies retrieves all the records using an executor.
func Frequencies(mods ...qm.QueryMod) frequencyQuery {
	mods = append(mods, qm.From("\"frequencies\""))
	return frequencyQuery{NewQuery(mods...)}
}

// FindFrequency retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFrequency(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Frequency, error) {
	frequencyObj := &Frequency{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"frequencies\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, frequencyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from frequencies")
	}

	return frequencyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Frequency) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no frequencies provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(frequencyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	frequencyInsertCacheMut.RLock()
	cache, cached := frequencyInsertCache[key]
	frequencyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			frequencyColumns,
			frequencyColumnsWithDefault,
			frequencyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(frequencyType, frequencyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(frequencyType, frequencyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"frequencies\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"frequencies\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into frequencies")
	}

	if !cached {
		frequencyInsertCacheMut.Lock()
		frequencyInsertCache[key] = cache
		frequencyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Frequency.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Frequency) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	frequencyUpdateCacheMut.RLock()
	cache, cached := frequencyUpdateCache[key]
	frequencyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			frequencyColumns,
			frequencyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update frequencies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"frequencies\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, frequencyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(frequencyType, frequencyMapping, append(wl, frequencyPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update frequencies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for frequencies")
	}

	if !cached {
		frequencyUpdateCacheMut.Lock()
		frequencyUpdateCache[key] = cache
		frequencyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q frequencyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for frequencies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for frequencies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FrequencySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), frequencyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"frequencies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, frequencyPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in frequency slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all frequency")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Frequency) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no frequencies provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(frequencyColumnsWithDefault, o)

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

	frequencyUpsertCacheMut.RLock()
	cache, cached := frequencyUpsertCache[key]
	frequencyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			frequencyColumns,
			frequencyColumnsWithDefault,
			frequencyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			frequencyColumns,
			frequencyPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert frequencies, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(frequencyPrimaryKeyColumns))
			copy(conflict, frequencyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"frequencies\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(frequencyType, frequencyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(frequencyType, frequencyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert frequencies")
	}

	if !cached {
		frequencyUpsertCacheMut.Lock()
		frequencyUpsertCache[key] = cache
		frequencyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Frequency record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Frequency) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Frequency provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), frequencyPrimaryKeyMapping)
	sql := "DELETE FROM \"frequencies\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from frequencies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for frequencies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q frequencyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no frequencyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from frequencies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for frequencies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FrequencySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Frequency slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(frequencyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), frequencyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"frequencies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, frequencyPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from frequency slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for frequencies")
	}

	if len(frequencyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Frequency) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFrequency(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FrequencySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FrequencySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), frequencyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"frequencies\".* FROM \"frequencies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, frequencyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FrequencySlice")
	}

	*o = slice

	return nil
}

// FrequencyExists checks if the Frequency row exists.
func FrequencyExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"frequencies\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if frequencies exists")
	}

	return exists, nil
}
