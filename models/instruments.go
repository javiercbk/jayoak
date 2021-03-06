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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Instrument is an object representing the database table.
type Instrument struct {
	ID             int64        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name           string       `boil:"name" json:"name" toml:"name" yaml:"name"`
	OrganizationID int64        `boil:"organization_id" json:"organizationID" toml:"organizationID" yaml:"organizationID"`
	CreatorID      int64        `boil:"creator_id" json:"creatorID" toml:"creatorID" yaml:"creatorID"`
	CreatedAt      null.Time    `boil:"created_at" json:"createdAt,omitempty" toml:"createdAt" yaml:"createdAt,omitempty"`
	UpdatedAt      null.Time    `boil:"updated_at" json:"updatedAt,omitempty" toml:"updatedAt" yaml:"updatedAt,omitempty"`
	R              *instrumentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L              instrumentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InstrumentColumns = struct {
	ID             string
	Name           string
	OrganizationID string
	CreatorID      string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	Name:           "name",
	OrganizationID: "organization_id",
	CreatorID:      "creator_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// InstrumentRels is where relationship names are stored.
var InstrumentRels = struct {
	Creator      string
	Organization string
	Sounds       string
}{
	Creator:      "Creator",
	Organization: "Organization",
	Sounds:       "Sounds",
}

// instrumentR is where relationships are stored.
type instrumentR struct {
	Creator      *User
	Organization *Organization
	Sounds       SoundSlice
}

// NewStruct creates a new relationship struct
func (*instrumentR) NewStruct() *instrumentR {
	return &instrumentR{}
}

// instrumentL is where Load methods for each relationship are stored.
type instrumentL struct{}

var (
	instrumentColumns               = []string{"id", "name", "organization_id", "creator_id", "created_at", "updated_at"}
	instrumentColumnsWithoutDefault = []string{"name", "organization_id", "creator_id", "created_at", "updated_at"}
	instrumentColumnsWithDefault    = []string{"id"}
	instrumentPrimaryKeyColumns     = []string{"id"}
)

type (
	// InstrumentSlice is an alias for a slice of pointers to Instrument.
	// This should generally be used opposed to []Instrument.
	InstrumentSlice []*Instrument
	// InstrumentHook is the signature for custom Instrument hook methods
	InstrumentHook func(context.Context, boil.ContextExecutor, *Instrument) error

	instrumentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	instrumentType                 = reflect.TypeOf(&Instrument{})
	instrumentMapping              = queries.MakeStructMapping(instrumentType)
	instrumentPrimaryKeyMapping, _ = queries.BindMapping(instrumentType, instrumentMapping, instrumentPrimaryKeyColumns)
	instrumentInsertCacheMut       sync.RWMutex
	instrumentInsertCache          = make(map[string]insertCache)
	instrumentUpdateCacheMut       sync.RWMutex
	instrumentUpdateCache          = make(map[string]updateCache)
	instrumentUpsertCacheMut       sync.RWMutex
	instrumentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var instrumentBeforeInsertHooks []InstrumentHook
var instrumentBeforeUpdateHooks []InstrumentHook
var instrumentBeforeDeleteHooks []InstrumentHook
var instrumentBeforeUpsertHooks []InstrumentHook

var instrumentAfterInsertHooks []InstrumentHook
var instrumentAfterSelectHooks []InstrumentHook
var instrumentAfterUpdateHooks []InstrumentHook
var instrumentAfterDeleteHooks []InstrumentHook
var instrumentAfterUpsertHooks []InstrumentHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Instrument) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range instrumentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Instrument) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range instrumentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Instrument) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range instrumentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Instrument) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range instrumentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Instrument) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range instrumentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Instrument) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range instrumentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Instrument) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range instrumentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Instrument) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range instrumentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Instrument) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range instrumentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInstrumentHook registers your hook function for all future operations.
func AddInstrumentHook(hookPoint boil.HookPoint, instrumentHook InstrumentHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		instrumentBeforeInsertHooks = append(instrumentBeforeInsertHooks, instrumentHook)
	case boil.BeforeUpdateHook:
		instrumentBeforeUpdateHooks = append(instrumentBeforeUpdateHooks, instrumentHook)
	case boil.BeforeDeleteHook:
		instrumentBeforeDeleteHooks = append(instrumentBeforeDeleteHooks, instrumentHook)
	case boil.BeforeUpsertHook:
		instrumentBeforeUpsertHooks = append(instrumentBeforeUpsertHooks, instrumentHook)
	case boil.AfterInsertHook:
		instrumentAfterInsertHooks = append(instrumentAfterInsertHooks, instrumentHook)
	case boil.AfterSelectHook:
		instrumentAfterSelectHooks = append(instrumentAfterSelectHooks, instrumentHook)
	case boil.AfterUpdateHook:
		instrumentAfterUpdateHooks = append(instrumentAfterUpdateHooks, instrumentHook)
	case boil.AfterDeleteHook:
		instrumentAfterDeleteHooks = append(instrumentAfterDeleteHooks, instrumentHook)
	case boil.AfterUpsertHook:
		instrumentAfterUpsertHooks = append(instrumentAfterUpsertHooks, instrumentHook)
	}
}

// One returns a single instrument record from the query.
func (q instrumentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Instrument, error) {
	o := &Instrument{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for instruments")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Instrument records from the query.
func (q instrumentQuery) All(ctx context.Context, exec boil.ContextExecutor) (InstrumentSlice, error) {
	var o []*Instrument

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Instrument slice")
	}

	if len(instrumentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Instrument records in the query.
func (q instrumentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count instruments rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q instrumentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if instruments exists")
	}

	return count > 0, nil
}

// Creator pointed to by the foreign key.
func (o *Instrument) Creator(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.CreatorID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Organization pointed to by the foreign key.
func (o *Instrument) Organization(mods ...qm.QueryMod) organizationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.OrganizationID),
	}

	queryMods = append(queryMods, mods...)

	query := Organizations(queryMods...)
	queries.SetFrom(query.Query, "\"organizations\"")

	return query
}

// Sounds retrieves all the sound's Sounds with an executor.
func (o *Instrument) Sounds(mods ...qm.QueryMod) soundQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"sounds\".\"instrument_id\"=?", o.ID),
	)

	query := Sounds(queryMods...)
	queries.SetFrom(query.Query, "\"sounds\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"sounds\".*"})
	}

	return query
}

// LoadCreator allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (instrumentL) LoadCreator(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInstrument interface{}, mods queries.Applicator) error {
	var slice []*Instrument
	var object *Instrument

	if singular {
		object = maybeInstrument.(*Instrument)
	} else {
		slice = *maybeInstrument.(*[]*Instrument)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &instrumentR{}
		}
		args = append(args, object.CreatorID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &instrumentR{}
			}

			for _, a := range args {
				if a == obj.CreatorID {
					continue Outer
				}
			}

			args = append(args, obj.CreatorID)

		}
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(instrumentAfterSelectHooks) != 0 {
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
		object.R.Creator = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.CreatorInstruments = append(foreign.R.CreatorInstruments, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CreatorID == foreign.ID {
				local.R.Creator = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.CreatorInstruments = append(foreign.R.CreatorInstruments, local)
				break
			}
		}
	}

	return nil
}

// LoadOrganization allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (instrumentL) LoadOrganization(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInstrument interface{}, mods queries.Applicator) error {
	var slice []*Instrument
	var object *Instrument

	if singular {
		object = maybeInstrument.(*Instrument)
	} else {
		slice = *maybeInstrument.(*[]*Instrument)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &instrumentR{}
		}
		args = append(args, object.OrganizationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &instrumentR{}
			}

			for _, a := range args {
				if a == obj.OrganizationID {
					continue Outer
				}
			}

			args = append(args, obj.OrganizationID)

		}
	}

	query := NewQuery(qm.From(`organizations`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Organization")
	}

	var resultSlice []*Organization
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Organization")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for organizations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for organizations")
	}

	if len(instrumentAfterSelectHooks) != 0 {
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
		object.R.Organization = foreign
		if foreign.R == nil {
			foreign.R = &organizationR{}
		}
		foreign.R.Instruments = append(foreign.R.Instruments, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrganizationID == foreign.ID {
				local.R.Organization = foreign
				if foreign.R == nil {
					foreign.R = &organizationR{}
				}
				foreign.R.Instruments = append(foreign.R.Instruments, local)
				break
			}
		}
	}

	return nil
}

// LoadSounds allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (instrumentL) LoadSounds(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInstrument interface{}, mods queries.Applicator) error {
	var slice []*Instrument
	var object *Instrument

	if singular {
		object = maybeInstrument.(*Instrument)
	} else {
		slice = *maybeInstrument.(*[]*Instrument)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &instrumentR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &instrumentR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`sounds`), qm.WhereIn(`instrument_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load sounds")
	}

	var resultSlice []*Sound
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice sounds")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on sounds")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sounds")
	}

	if len(soundAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Sounds = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &soundR{}
			}
			foreign.R.Instrument = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.InstrumentID) {
				local.R.Sounds = append(local.R.Sounds, foreign)
				if foreign.R == nil {
					foreign.R = &soundR{}
				}
				foreign.R.Instrument = local
				break
			}
		}
	}

	return nil
}

// SetCreator of the instrument to the related item.
// Sets o.R.Creator to related.
// Adds o to related.R.CreatorInstruments.
func (o *Instrument) SetCreator(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"instruments\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"creator_id"}),
		strmangle.WhereClause("\"", "\"", 2, instrumentPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CreatorID = related.ID
	if o.R == nil {
		o.R = &instrumentR{
			Creator: related,
		}
	} else {
		o.R.Creator = related
	}

	if related.R == nil {
		related.R = &userR{
			CreatorInstruments: InstrumentSlice{o},
		}
	} else {
		related.R.CreatorInstruments = append(related.R.CreatorInstruments, o)
	}

	return nil
}

// SetOrganization of the instrument to the related item.
// Sets o.R.Organization to related.
// Adds o to related.R.Instruments.
func (o *Instrument) SetOrganization(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Organization) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"instruments\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"organization_id"}),
		strmangle.WhereClause("\"", "\"", 2, instrumentPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrganizationID = related.ID
	if o.R == nil {
		o.R = &instrumentR{
			Organization: related,
		}
	} else {
		o.R.Organization = related
	}

	if related.R == nil {
		related.R = &organizationR{
			Instruments: InstrumentSlice{o},
		}
	} else {
		related.R.Instruments = append(related.R.Instruments, o)
	}

	return nil
}

// AddSounds adds the given related objects to the existing relationships
// of the instrument, optionally inserting them as new records.
// Appends related to o.R.Sounds.
// Sets related.R.Instrument appropriately.
func (o *Instrument) AddSounds(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Sound) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.InstrumentID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"sounds\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"instrument_id"}),
				strmangle.WhereClause("\"", "\"", 2, soundPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.InstrumentID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &instrumentR{
			Sounds: related,
		}
	} else {
		o.R.Sounds = append(o.R.Sounds, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &soundR{
				Instrument: o,
			}
		} else {
			rel.R.Instrument = o
		}
	}
	return nil
}

// SetSounds removes all previously related items of the
// instrument replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Instrument's Sounds accordingly.
// Replaces o.R.Sounds with related.
// Sets related.R.Instrument's Sounds accordingly.
func (o *Instrument) SetSounds(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Sound) error {
	query := "update \"sounds\" set \"instrument_id\" = null where \"instrument_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Sounds {
			queries.SetScanner(&rel.InstrumentID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Instrument = nil
		}

		o.R.Sounds = nil
	}
	return o.AddSounds(ctx, exec, insert, related...)
}

// RemoveSounds relationships from objects passed in.
// Removes related items from R.Sounds (uses pointer comparison, removal does not keep order)
// Sets related.R.Instrument.
func (o *Instrument) RemoveSounds(ctx context.Context, exec boil.ContextExecutor, related ...*Sound) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.InstrumentID, nil)
		if rel.R != nil {
			rel.R.Instrument = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("instrument_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Sounds {
			if rel != ri {
				continue
			}

			ln := len(o.R.Sounds)
			if ln > 1 && i < ln-1 {
				o.R.Sounds[i] = o.R.Sounds[ln-1]
			}
			o.R.Sounds = o.R.Sounds[:ln-1]
			break
		}
	}

	return nil
}

// Instruments retrieves all the records using an executor.
func Instruments(mods ...qm.QueryMod) instrumentQuery {
	mods = append(mods, qm.From("\"instruments\""))
	return instrumentQuery{NewQuery(mods...)}
}

// FindInstrument retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInstrument(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Instrument, error) {
	instrumentObj := &Instrument{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"instruments\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, instrumentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from instruments")
	}

	return instrumentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Instrument) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no instruments provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(instrumentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	instrumentInsertCacheMut.RLock()
	cache, cached := instrumentInsertCache[key]
	instrumentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			instrumentColumns,
			instrumentColumnsWithDefault,
			instrumentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(instrumentType, instrumentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(instrumentType, instrumentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"instruments\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"instruments\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into instruments")
	}

	if !cached {
		instrumentInsertCacheMut.Lock()
		instrumentInsertCache[key] = cache
		instrumentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Instrument.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Instrument) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	instrumentUpdateCacheMut.RLock()
	cache, cached := instrumentUpdateCache[key]
	instrumentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			instrumentColumns,
			instrumentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update instruments, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"instruments\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, instrumentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(instrumentType, instrumentMapping, append(wl, instrumentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update instruments row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for instruments")
	}

	if !cached {
		instrumentUpdateCacheMut.Lock()
		instrumentUpdateCache[key] = cache
		instrumentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q instrumentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for instruments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for instruments")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InstrumentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), instrumentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"instruments\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, instrumentPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in instrument slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all instrument")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Instrument) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no instruments provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(instrumentColumnsWithDefault, o)

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

	instrumentUpsertCacheMut.RLock()
	cache, cached := instrumentUpsertCache[key]
	instrumentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			instrumentColumns,
			instrumentColumnsWithDefault,
			instrumentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			instrumentColumns,
			instrumentPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert instruments, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(instrumentPrimaryKeyColumns))
			copy(conflict, instrumentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"instruments\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(instrumentType, instrumentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(instrumentType, instrumentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert instruments")
	}

	if !cached {
		instrumentUpsertCacheMut.Lock()
		instrumentUpsertCache[key] = cache
		instrumentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Instrument record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Instrument) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Instrument provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), instrumentPrimaryKeyMapping)
	sql := "DELETE FROM \"instruments\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from instruments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for instruments")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q instrumentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no instrumentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from instruments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for instruments")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InstrumentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Instrument slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(instrumentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), instrumentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"instruments\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, instrumentPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from instrument slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for instruments")
	}

	if len(instrumentAfterDeleteHooks) != 0 {
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
func (o *Instrument) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInstrument(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InstrumentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InstrumentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), instrumentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"instruments\".* FROM \"instruments\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, instrumentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InstrumentSlice")
	}

	*o = slice

	return nil
}

// InstrumentExists checks if the Instrument row exists.
func InstrumentExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"instruments\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if instruments exists")
	}

	return exists, nil
}
