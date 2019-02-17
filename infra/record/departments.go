// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package record

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
)

// Department is an object representing the database table.
type Department struct {
	ID        int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string `boil:"name" json:"name" toml:"name" yaml:"name"`
	ShortName string `boil:"short_name" json:"short_name" toml:"short_name" yaml:"short_name"`

	R *departmentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L departmentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DepartmentColumns = struct {
	ID        string
	Name      string
	ShortName string
}{
	ID:        "id",
	Name:      "name",
	ShortName: "short_name",
}

// DepartmentRels is where relationship names are stored.
var DepartmentRels = struct {
	Profiles string
}{
	Profiles: "Profiles",
}

// departmentR is where relationships are stored.
type departmentR struct {
	Profiles ProfileSlice
}

// NewStruct creates a new relationship struct
func (*departmentR) NewStruct() *departmentR {
	return &departmentR{}
}

// departmentL is where Load methods for each relationship are stored.
type departmentL struct{}

var (
	departmentColumns               = []string{"id", "name", "short_name"}
	departmentColumnsWithoutDefault = []string{"name", "short_name"}
	departmentColumnsWithDefault    = []string{"id"}
	departmentPrimaryKeyColumns     = []string{"id"}
)

type (
	// DepartmentSlice is an alias for a slice of pointers to Department.
	// This should generally be used opposed to []Department.
	DepartmentSlice []*Department
	// DepartmentHook is the signature for custom Department hook methods
	DepartmentHook func(context.Context, boil.ContextExecutor, *Department) error

	departmentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	departmentType                 = reflect.TypeOf(&Department{})
	departmentMapping              = queries.MakeStructMapping(departmentType)
	departmentPrimaryKeyMapping, _ = queries.BindMapping(departmentType, departmentMapping, departmentPrimaryKeyColumns)
	departmentInsertCacheMut       sync.RWMutex
	departmentInsertCache          = make(map[string]insertCache)
	departmentUpdateCacheMut       sync.RWMutex
	departmentUpdateCache          = make(map[string]updateCache)
	departmentUpsertCacheMut       sync.RWMutex
	departmentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var departmentBeforeInsertHooks []DepartmentHook
var departmentBeforeUpdateHooks []DepartmentHook
var departmentBeforeDeleteHooks []DepartmentHook
var departmentBeforeUpsertHooks []DepartmentHook

var departmentAfterInsertHooks []DepartmentHook
var departmentAfterSelectHooks []DepartmentHook
var departmentAfterUpdateHooks []DepartmentHook
var departmentAfterDeleteHooks []DepartmentHook
var departmentAfterUpsertHooks []DepartmentHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Department) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range departmentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Department) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range departmentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Department) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range departmentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Department) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range departmentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Department) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range departmentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Department) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range departmentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Department) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range departmentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Department) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range departmentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Department) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range departmentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDepartmentHook registers your hook function for all future operations.
func AddDepartmentHook(hookPoint boil.HookPoint, departmentHook DepartmentHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		departmentBeforeInsertHooks = append(departmentBeforeInsertHooks, departmentHook)
	case boil.BeforeUpdateHook:
		departmentBeforeUpdateHooks = append(departmentBeforeUpdateHooks, departmentHook)
	case boil.BeforeDeleteHook:
		departmentBeforeDeleteHooks = append(departmentBeforeDeleteHooks, departmentHook)
	case boil.BeforeUpsertHook:
		departmentBeforeUpsertHooks = append(departmentBeforeUpsertHooks, departmentHook)
	case boil.AfterInsertHook:
		departmentAfterInsertHooks = append(departmentAfterInsertHooks, departmentHook)
	case boil.AfterSelectHook:
		departmentAfterSelectHooks = append(departmentAfterSelectHooks, departmentHook)
	case boil.AfterUpdateHook:
		departmentAfterUpdateHooks = append(departmentAfterUpdateHooks, departmentHook)
	case boil.AfterDeleteHook:
		departmentAfterDeleteHooks = append(departmentAfterDeleteHooks, departmentHook)
	case boil.AfterUpsertHook:
		departmentAfterUpsertHooks = append(departmentAfterUpsertHooks, departmentHook)
	}
}

// One returns a single department record from the query.
func (q departmentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Department, error) {
	o := &Department{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: failed to execute a one query for departments")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Department records from the query.
func (q departmentQuery) All(ctx context.Context, exec boil.ContextExecutor) (DepartmentSlice, error) {
	var o []*Department

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "record: failed to assign all query results to Department slice")
	}

	if len(departmentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Department records in the query.
func (q departmentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to count departments rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q departmentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "record: failed to check if departments exists")
	}

	return count > 0, nil
}

// Profiles retrieves all the profile's Profiles with an executor.
func (o *Department) Profiles(mods ...qm.QueryMod) profileQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"profiles\".\"department_id\"=?", o.ID),
	)

	query := Profiles(queryMods...)
	queries.SetFrom(query.Query, "\"profiles\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"profiles\".*"})
	}

	return query
}

// LoadProfiles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (departmentL) LoadProfiles(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDepartment interface{}, mods queries.Applicator) error {
	var slice []*Department
	var object *Department

	if singular {
		object = maybeDepartment.(*Department)
	} else {
		slice = *maybeDepartment.(*[]*Department)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &departmentR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &departmentR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`profiles`), qm.WhereIn(`department_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load profiles")
	}

	var resultSlice []*Profile
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice profiles")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on profiles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for profiles")
	}

	if len(profileAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Profiles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &profileR{}
			}
			foreign.R.Department = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.DepartmentID) {
				local.R.Profiles = append(local.R.Profiles, foreign)
				if foreign.R == nil {
					foreign.R = &profileR{}
				}
				foreign.R.Department = local
				break
			}
		}
	}

	return nil
}

// AddProfiles adds the given related objects to the existing relationships
// of the department, optionally inserting them as new records.
// Appends related to o.R.Profiles.
// Sets related.R.Department appropriately.
func (o *Department) AddProfiles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Profile) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.DepartmentID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"profiles\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"department_id"}),
				strmangle.WhereClause("\"", "\"", 2, profilePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.DepartmentID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &departmentR{
			Profiles: related,
		}
	} else {
		o.R.Profiles = append(o.R.Profiles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &profileR{
				Department: o,
			}
		} else {
			rel.R.Department = o
		}
	}
	return nil
}

// SetProfiles removes all previously related items of the
// department replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Department's Profiles accordingly.
// Replaces o.R.Profiles with related.
// Sets related.R.Department's Profiles accordingly.
func (o *Department) SetProfiles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Profile) error {
	query := "update \"profiles\" set \"department_id\" = null where \"department_id\" = $1"
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
		for _, rel := range o.R.Profiles {
			queries.SetScanner(&rel.DepartmentID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Department = nil
		}

		o.R.Profiles = nil
	}
	return o.AddProfiles(ctx, exec, insert, related...)
}

// RemoveProfiles relationships from objects passed in.
// Removes related items from R.Profiles (uses pointer comparison, removal does not keep order)
// Sets related.R.Department.
func (o *Department) RemoveProfiles(ctx context.Context, exec boil.ContextExecutor, related ...*Profile) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.DepartmentID, nil)
		if rel.R != nil {
			rel.R.Department = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("department_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Profiles {
			if rel != ri {
				continue
			}

			ln := len(o.R.Profiles)
			if ln > 1 && i < ln-1 {
				o.R.Profiles[i] = o.R.Profiles[ln-1]
			}
			o.R.Profiles = o.R.Profiles[:ln-1]
			break
		}
	}

	return nil
}

// Departments retrieves all the records using an executor.
func Departments(mods ...qm.QueryMod) departmentQuery {
	mods = append(mods, qm.From("\"departments\""))
	return departmentQuery{NewQuery(mods...)}
}

// FindDepartment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDepartment(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Department, error) {
	departmentObj := &Department{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"departments\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, departmentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: unable to select from departments")
	}

	return departmentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Department) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("record: no departments provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(departmentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	departmentInsertCacheMut.RLock()
	cache, cached := departmentInsertCache[key]
	departmentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			departmentColumns,
			departmentColumnsWithDefault,
			departmentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(departmentType, departmentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(departmentType, departmentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"departments\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"departments\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "record: unable to insert into departments")
	}

	if !cached {
		departmentInsertCacheMut.Lock()
		departmentInsertCache[key] = cache
		departmentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Department.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Department) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	departmentUpdateCacheMut.RLock()
	cache, cached := departmentUpdateCache[key]
	departmentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			departmentColumns,
			departmentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("record: unable to update departments, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"departments\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, departmentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(departmentType, departmentMapping, append(wl, departmentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "record: unable to update departments row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by update for departments")
	}

	if !cached {
		departmentUpdateCacheMut.Lock()
		departmentUpdateCache[key] = cache
		departmentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q departmentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all for departments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected for departments")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DepartmentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("record: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), departmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"departments\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, departmentPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all in department slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected all in update all department")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Department) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("record: no departments provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(departmentColumnsWithDefault, o)

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

	departmentUpsertCacheMut.RLock()
	cache, cached := departmentUpsertCache[key]
	departmentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			departmentColumns,
			departmentColumnsWithDefault,
			departmentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			departmentColumns,
			departmentPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("record: unable to upsert departments, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(departmentPrimaryKeyColumns))
			copy(conflict, departmentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"departments\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(departmentType, departmentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(departmentType, departmentMapping, ret)
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
		return errors.Wrap(err, "record: unable to upsert departments")
	}

	if !cached {
		departmentUpsertCacheMut.Lock()
		departmentUpsertCache[key] = cache
		departmentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Department record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Department) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no Department provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), departmentPrimaryKeyMapping)
	sql := "DELETE FROM \"departments\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete from departments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by delete for departments")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q departmentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("record: no departmentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from departments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for departments")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DepartmentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no Department slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(departmentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), departmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"departments\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, departmentPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from department slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for departments")
	}

	if len(departmentAfterDeleteHooks) != 0 {
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
func (o *Department) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDepartment(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DepartmentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DepartmentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), departmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"departments\".* FROM \"departments\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, departmentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "record: unable to reload all in DepartmentSlice")
	}

	*o = slice

	return nil
}

// DepartmentExists checks if the Department row exists.
func DepartmentExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"departments\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "record: unable to check if departments exists")
	}

	return exists, nil
}