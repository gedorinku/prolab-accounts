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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Entry is an object representing the database table.
type Entry struct {
	ID          int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       string     `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description string     `boil:"description" json:"description" toml:"description" yaml:"description"`
	Content     string     `boil:"content" json:"content" toml:"content" yaml:"content"`
	Link        string     `boil:"link" json:"link" toml:"link" yaml:"link"`
	AuthorID    int64      `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	GUID        string     `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	BlogID      null.Int64 `boil:"blog_id" json:"blog_id,omitempty" toml:"blog_id" yaml:"blog_id,omitempty"`
	CreatedAt   time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time  `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ImageURL    string     `boil:"image_url" json:"image_url" toml:"image_url" yaml:"image_url"`

	R *entryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L entryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EntryColumns = struct {
	ID          string
	Title       string
	Description string
	Content     string
	Link        string
	AuthorID    string
	GUID        string
	BlogID      string
	CreatedAt   string
	UpdatedAt   string
	ImageURL    string
}{
	ID:          "id",
	Title:       "title",
	Description: "description",
	Content:     "content",
	Link:        "link",
	AuthorID:    "author_id",
	GUID:        "guid",
	BlogID:      "blog_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	ImageURL:    "image_url",
}

// Generated where

var EntryWhere = struct {
	ID          whereHelperint64
	Title       whereHelperstring
	Description whereHelperstring
	Content     whereHelperstring
	Link        whereHelperstring
	AuthorID    whereHelperint64
	GUID        whereHelperstring
	BlogID      whereHelpernull_Int64
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
	ImageURL    whereHelperstring
}{
	ID:          whereHelperint64{field: `id`},
	Title:       whereHelperstring{field: `title`},
	Description: whereHelperstring{field: `description`},
	Content:     whereHelperstring{field: `content`},
	Link:        whereHelperstring{field: `link`},
	AuthorID:    whereHelperint64{field: `author_id`},
	GUID:        whereHelperstring{field: `guid`},
	BlogID:      whereHelpernull_Int64{field: `blog_id`},
	CreatedAt:   whereHelpertime_Time{field: `created_at`},
	UpdatedAt:   whereHelpertime_Time{field: `updated_at`},
	ImageURL:    whereHelperstring{field: `image_url`},
}

// EntryRels is where relationship names are stored.
var EntryRels = struct {
	Author string
	Blog   string
}{
	Author: "Author",
	Blog:   "Blog",
}

// entryR is where relationships are stored.
type entryR struct {
	Author *User
	Blog   *Blog
}

// NewStruct creates a new relationship struct
func (*entryR) NewStruct() *entryR {
	return &entryR{}
}

// entryL is where Load methods for each relationship are stored.
type entryL struct{}

var (
	entryColumns               = []string{"id", "title", "description", "content", "link", "author_id", "guid", "blog_id", "created_at", "updated_at", "image_url"}
	entryColumnsWithoutDefault = []string{"title", "description", "content", "link", "author_id", "guid", "blog_id", "created_at", "updated_at", "image_url"}
	entryColumnsWithDefault    = []string{"id"}
	entryPrimaryKeyColumns     = []string{"id"}
)

type (
	// EntrySlice is an alias for a slice of pointers to Entry.
	// This should generally be used opposed to []Entry.
	EntrySlice []*Entry
	// EntryHook is the signature for custom Entry hook methods
	EntryHook func(context.Context, boil.ContextExecutor, *Entry) error

	entryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	entryType                 = reflect.TypeOf(&Entry{})
	entryMapping              = queries.MakeStructMapping(entryType)
	entryPrimaryKeyMapping, _ = queries.BindMapping(entryType, entryMapping, entryPrimaryKeyColumns)
	entryInsertCacheMut       sync.RWMutex
	entryInsertCache          = make(map[string]insertCache)
	entryUpdateCacheMut       sync.RWMutex
	entryUpdateCache          = make(map[string]updateCache)
	entryUpsertCacheMut       sync.RWMutex
	entryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var entryBeforeInsertHooks []EntryHook
var entryBeforeUpdateHooks []EntryHook
var entryBeforeDeleteHooks []EntryHook
var entryBeforeUpsertHooks []EntryHook

var entryAfterInsertHooks []EntryHook
var entryAfterSelectHooks []EntryHook
var entryAfterUpdateHooks []EntryHook
var entryAfterDeleteHooks []EntryHook
var entryAfterUpsertHooks []EntryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Entry) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Entry) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Entry) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Entry) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Entry) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Entry) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Entry) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Entry) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Entry) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEntryHook registers your hook function for all future operations.
func AddEntryHook(hookPoint boil.HookPoint, entryHook EntryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		entryBeforeInsertHooks = append(entryBeforeInsertHooks, entryHook)
	case boil.BeforeUpdateHook:
		entryBeforeUpdateHooks = append(entryBeforeUpdateHooks, entryHook)
	case boil.BeforeDeleteHook:
		entryBeforeDeleteHooks = append(entryBeforeDeleteHooks, entryHook)
	case boil.BeforeUpsertHook:
		entryBeforeUpsertHooks = append(entryBeforeUpsertHooks, entryHook)
	case boil.AfterInsertHook:
		entryAfterInsertHooks = append(entryAfterInsertHooks, entryHook)
	case boil.AfterSelectHook:
		entryAfterSelectHooks = append(entryAfterSelectHooks, entryHook)
	case boil.AfterUpdateHook:
		entryAfterUpdateHooks = append(entryAfterUpdateHooks, entryHook)
	case boil.AfterDeleteHook:
		entryAfterDeleteHooks = append(entryAfterDeleteHooks, entryHook)
	case boil.AfterUpsertHook:
		entryAfterUpsertHooks = append(entryAfterUpsertHooks, entryHook)
	}
}

// One returns a single entry record from the query.
func (q entryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Entry, error) {
	o := &Entry{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: failed to execute a one query for entries")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Entry records from the query.
func (q entryQuery) All(ctx context.Context, exec boil.ContextExecutor) (EntrySlice, error) {
	var o []*Entry

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "record: failed to assign all query results to Entry slice")
	}

	if len(entryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Entry records in the query.
func (q entryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to count entries rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q entryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "record: failed to check if entries exists")
	}

	return count > 0, nil
}

// Author pointed to by the foreign key.
func (o *Entry) Author(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.AuthorID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Blog pointed to by the foreign key.
func (o *Entry) Blog(mods ...qm.QueryMod) blogQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.BlogID),
	}

	queryMods = append(queryMods, mods...)

	query := Blogs(queryMods...)
	queries.SetFrom(query.Query, "\"blogs\"")

	return query
}

// LoadAuthor allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (entryL) LoadAuthor(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEntry interface{}, mods queries.Applicator) error {
	var slice []*Entry
	var object *Entry

	if singular {
		object = maybeEntry.(*Entry)
	} else {
		slice = *maybeEntry.(*[]*Entry)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &entryR{}
		}
		args = append(args, object.AuthorID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &entryR{}
			}

			for _, a := range args {
				if a == obj.AuthorID {
					continue Outer
				}
			}

			args = append(args, obj.AuthorID)

		}
	}

	if len(args) == 0 {
		return nil
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

	if len(entryAfterSelectHooks) != 0 {
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
		object.R.Author = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.AuthorEntries = append(foreign.R.AuthorEntries, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AuthorID == foreign.ID {
				local.R.Author = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.AuthorEntries = append(foreign.R.AuthorEntries, local)
				break
			}
		}
	}

	return nil
}

// LoadBlog allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (entryL) LoadBlog(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEntry interface{}, mods queries.Applicator) error {
	var slice []*Entry
	var object *Entry

	if singular {
		object = maybeEntry.(*Entry)
	} else {
		slice = *maybeEntry.(*[]*Entry)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &entryR{}
		}
		if !queries.IsNil(object.BlogID) {
			args = append(args, object.BlogID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &entryR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.BlogID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.BlogID) {
				args = append(args, obj.BlogID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`blogs`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Blog")
	}

	var resultSlice []*Blog
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Blog")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for blogs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for blogs")
	}

	if len(entryAfterSelectHooks) != 0 {
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
		object.R.Blog = foreign
		if foreign.R == nil {
			foreign.R = &blogR{}
		}
		foreign.R.Entries = append(foreign.R.Entries, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.BlogID, foreign.ID) {
				local.R.Blog = foreign
				if foreign.R == nil {
					foreign.R = &blogR{}
				}
				foreign.R.Entries = append(foreign.R.Entries, local)
				break
			}
		}
	}

	return nil
}

// SetAuthor of the entry to the related item.
// Sets o.R.Author to related.
// Adds o to related.R.AuthorEntries.
func (o *Entry) SetAuthor(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"entries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"author_id"}),
		strmangle.WhereClause("\"", "\"", 2, entryPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AuthorID = related.ID
	if o.R == nil {
		o.R = &entryR{
			Author: related,
		}
	} else {
		o.R.Author = related
	}

	if related.R == nil {
		related.R = &userR{
			AuthorEntries: EntrySlice{o},
		}
	} else {
		related.R.AuthorEntries = append(related.R.AuthorEntries, o)
	}

	return nil
}

// SetBlog of the entry to the related item.
// Sets o.R.Blog to related.
// Adds o to related.R.Entries.
func (o *Entry) SetBlog(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Blog) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"entries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"blog_id"}),
		strmangle.WhereClause("\"", "\"", 2, entryPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.BlogID, related.ID)
	if o.R == nil {
		o.R = &entryR{
			Blog: related,
		}
	} else {
		o.R.Blog = related
	}

	if related.R == nil {
		related.R = &blogR{
			Entries: EntrySlice{o},
		}
	} else {
		related.R.Entries = append(related.R.Entries, o)
	}

	return nil
}

// RemoveBlog relationship.
// Sets o.R.Blog to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Entry) RemoveBlog(ctx context.Context, exec boil.ContextExecutor, related *Blog) error {
	var err error

	queries.SetScanner(&o.BlogID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("blog_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Blog = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Entries {
		if queries.Equal(o.BlogID, ri.BlogID) {
			continue
		}

		ln := len(related.R.Entries)
		if ln > 1 && i < ln-1 {
			related.R.Entries[i] = related.R.Entries[ln-1]
		}
		related.R.Entries = related.R.Entries[:ln-1]
		break
	}
	return nil
}

// Entries retrieves all the records using an executor.
func Entries(mods ...qm.QueryMod) entryQuery {
	mods = append(mods, qm.From("\"entries\""))
	return entryQuery{NewQuery(mods...)}
}

// FindEntry retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEntry(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Entry, error) {
	entryObj := &Entry{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"entries\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, entryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "record: unable to select from entries")
	}

	return entryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Entry) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("record: no entries provided for insertion")
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

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(entryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	entryInsertCacheMut.RLock()
	cache, cached := entryInsertCache[key]
	entryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			entryColumns,
			entryColumnsWithDefault,
			entryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(entryType, entryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(entryType, entryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"entries\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"entries\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "record: unable to insert into entries")
	}

	if !cached {
		entryInsertCacheMut.Lock()
		entryInsertCache[key] = cache
		entryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Entry.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Entry) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	entryUpdateCacheMut.RLock()
	cache, cached := entryUpdateCache[key]
	entryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			entryColumns,
			entryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("record: unable to update entries, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"entries\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, entryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(entryType, entryMapping, append(wl, entryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "record: unable to update entries row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by update for entries")
	}

	if !cached {
		entryUpdateCacheMut.Lock()
		entryUpdateCache[key] = cache
		entryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q entryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all for entries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected for entries")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EntrySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"entries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, entryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to update all in entry slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to retrieve rows affected all in update all entry")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Entry) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("record: no entries provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(entryColumnsWithDefault, o)

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

	entryUpsertCacheMut.RLock()
	cache, cached := entryUpsertCache[key]
	entryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			entryColumns,
			entryColumnsWithDefault,
			entryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			entryColumns,
			entryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("record: unable to upsert entries, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(entryPrimaryKeyColumns))
			copy(conflict, entryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"entries\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(entryType, entryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(entryType, entryMapping, ret)
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
		return errors.Wrap(err, "record: unable to upsert entries")
	}

	if !cached {
		entryUpsertCacheMut.Lock()
		entryUpsertCache[key] = cache
		entryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Entry record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Entry) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no Entry provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), entryPrimaryKeyMapping)
	sql := "DELETE FROM \"entries\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete from entries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by delete for entries")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q entryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("record: no entryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from entries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for entries")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EntrySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("record: no Entry slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(entryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"entries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, entryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "record: unable to delete all from entry slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "record: failed to get rows affected by deleteall for entries")
	}

	if len(entryAfterDeleteHooks) != 0 {
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
func (o *Entry) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEntry(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EntrySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EntrySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"entries\".* FROM \"entries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, entryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "record: unable to reload all in EntrySlice")
	}

	*o = slice

	return nil
}

// EntryExists checks if the Entry row exists.
func EntryExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"entries\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "record: unable to check if entries exists")
	}

	return exists, nil
}
